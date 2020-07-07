package listener

import (
	"fmt"
	"log"

	"github.com/ramonhpr/knotc/pkg/generated"
	"github.com/ramonhpr/knotc/pkg/model"
)

// ListenerImpl implements the knotListener interface
type ListenerImpl struct {
	*generated.BaseKnotListener
	Things map[string]*model.Thing
}

const unitFmt = "%s in %s"

// EnterStart is started after after the syntax tree is parsed
func (k *ListenerImpl) EnterStart(ctx *generated.StartContext) {
	k.Things = make(map[string]*model.Thing, len(ctx.GetThings()))
	for _, defCtx := range ctx.GetThings() {
		name := defCtx.GetName().GetText()
		if _, value := k.Things[name]; value {
			log.Fatalf("The thing name \"%s\" is already used", name)
		}

		newThing := model.Thing{Name: name}
		newThing.Sensors = make([]model.DataItem, len(defCtx.GetSensors()))
		for i, thingContentCtx := range defCtx.GetSensors() {
			newDataItem := createNewDataItem(uint(i), thingContentCtx)

			newThing.Sensors[i] = newDataItem
		}
		k.Things[name] = &newThing
	}
}

func createNewDataItem(id uint, ctx generated.IThingContentContext) model.DataItem {
	newDataItem := model.DataItem{}
	newDataItem.IsSensor = ctx.GetSensorType().GetTokenType() == generated.KnotParserSENSOR
	newDataItem.Value = ctx.GetTypeValue().GetText()
	newDataItem.Name = ctx.GetName().GetText()
	newDataItem.ID = id
	s := ctx.(*generated.ThingContentContext)
	setTypeValue(&newDataItem, ctx.GetTypeValue().GetTokenType(), ctx)

	setConfigCommon(s, newDataItem)

	return newDataItem
}

func setConfigCommon(s *generated.ThingContentContext, newDataItem model.DataItem) {
	if s.ConfigChanges() != nil {
		addOnConfigs(&newDataItem, model.KnotChanges, "")
	}

	if configTimeCtx := s.ConfigTime(); configTimeCtx != nil {
		addOnConfigs(&newDataItem, model.KnotTime, configTimeCtx.GetNumber().GetText())
	}
}

func setTypeValue(data *model.DataItem, tokenType int, ctx generated.IThingContentContext) {
	s := ctx.(*generated.ThingContentContext)

	switch tokenType {
	case generated.KnotLexerBOOL:
		data.DefaultValue = "false"
		data.TypeUnit = ctx.GetTypeUnit().GetText()
	case generated.KnotLexerBYTES:
		data.Value = "char[]" // TODO: verify this
		data.DefaultValue = "NULL"
		data.TypeUnit = s.COMMAND().GetText()
	case generated.KnotLexerINT, generated.KnotLexerFLOAT:
		data.DefaultValue = "0"
		if configUpperCtx := s.ConfigUpper(); configUpperCtx != nil {
			addOnConfigs(data, model.KnotUpper, configUpperCtx.GetNumber().GetText())
		}

		if configLowerCtx := s.ConfigLower(); configLowerCtx != nil {
			addOnConfigs(data, model.KnotLower, configLowerCtx.GetNumber().GetText())
		}

		_, err := fmt.Sscanf(s.UnitTypeOptions().GetText(), unitFmt, &data.TypeUnit, &data.Unit)
		if err != nil {
			log.Fatalf("Unable to recognaize unit. Reason: %v", err)
		}
	}
}

func addOnConfigs(data *model.DataItem, configType model.ConfigType, value string) {
	addConfigOnConfigs(data, model.Config{Type: configType, Value: value})
}

func addConfigOnConfigs(data *model.DataItem, config model.Config) {
	data.Configs = append(data.Configs, config)
}