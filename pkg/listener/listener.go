package listener

import (
	"fmt"
	"github.com/ramonhpr/knotc/pkg/generated"
	"github.com/ramonhpr/knotc/pkg/model"
)

// ListenerImpl implements the knotListener interface
type ListenerImpl struct {
	*generated.BaseKnotListener
	Things []model.Thing
	currentThing  uint
	currentSensor uint
}

const unitFmt = "%s in %s"

func (k *ListenerImpl) EnterStart(ctx *generated.StartContext) {
	if k.Things == nil {
		k.currentSensor = 0
	} else if k.Things[k.currentThing].Sensors == nil {
		k.currentThing = 0
	}
}

func (k *ListenerImpl) EnterDefinition(ctx *generated.DefinitionContext) {
	k.Things =  append(k.Things, model.Thing{})
	k.Things[k.currentThing].Name = ctx.IDENTIFIER().GetText()
}

func (k *ListenerImpl) ExitDefinition(ctx *generated.DefinitionContext) {
	k.currentThing++
	k.currentSensor = 0
}

func (k *ListenerImpl) EnterThingContent(ctx *generated.ThingContentContext) {
	k.Things[k.currentThing].Sensors = append(k.Things[k.currentThing].Sensors, model.DataItem{ID: k.currentSensor})
}

func (k *ListenerImpl) ExitThingContent(ctx *generated.ThingContentContext) {
	k.Things[k.currentThing].Sensors[k.currentSensor].IsSensor = ctx.GetOp().GetTokenType() == generated.KnotParserSENSOR
	k.currentSensor++
}

func (k *ListenerImpl) ExitBoolOpt(ctx *generated.BoolOptContext) {
	k.Things[k.currentThing].Sensors[k.currentSensor].Value = "bool"
	k.Things[k.currentThing].Sensors[k.currentSensor].DefaultValue = "false"
	k.Things[k.currentThing].Sensors[k.currentSensor].Name = ctx.IDENTIFIER().GetText()
	k.Things[k.currentThing].Sensors[k.currentSensor].TypeUnit = ctx.GetOp().GetText()
	if ctx.ConfigChanges() != nil{
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs = append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs = append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}
}

func (k *ListenerImpl) ExitNumberOpt(ctx *generated.NumberOptContext) {
	k.Things[k.currentThing].Sensors[k.currentSensor].Value = ctx.GetOp().GetText();
	k.Things[k.currentThing].Sensors[k.currentSensor].DefaultValue = "0"
	k.Things[k.currentThing].Sensors[k.currentSensor].Name = ctx.IDENTIFIER().GetText()
}


func (k *ListenerImpl) ExitConfig(ctx *generated.ConfigContext) {
	if ctx.ConfigChanges() != nil{
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs =  append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs =  append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigUpper(); tmp != nil {
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs =  append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotUpper, Value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigLower(); tmp != nil {
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs =  append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotLower, Value: tmp.GetNumber().GetText()})
	}
}


func (k *ListenerImpl) ExitBytesOpt(ctx *generated.BytesOptContext) {
	k.Things[k.currentThing].Sensors[k.currentSensor].Value = "char[]" // TODO: verify this
	k.Things[k.currentThing].Sensors[k.currentSensor].DefaultValue = "NULL"
	k.Things[k.currentThing].Sensors[k.currentSensor].Name = ctx.IDENTIFIER().GetText()
	k.Things[k.currentThing].Sensors[k.currentSensor].TypeUnit = "command"
	if ctx.ConfigChanges() != nil{
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs =  append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.Things[k.currentThing].Sensors[k.currentSensor].Configs =  append(k.Things[k.currentThing].Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}
}

func (k *ListenerImpl) ExitUnitTypeOptions(ctx *generated.UnitTypeOptionsContext) {
	fmt.Sscanf(ctx.GetText(), unitFmt, &k.Things[k.currentThing].Sensors[k.currentSensor].TypeUnit, &k.Things[k.currentThing].Sensors[k.currentSensor].Unit)
}
