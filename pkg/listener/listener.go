package listener

import (
	"fmt"
	"github.com/ramonhpr/knotc/pkg/generated"
	"github.com/ramonhpr/knotc/pkg/model"
)

// ListenerImpl implements the knotListener interface
type ListenerImpl struct {
	*generated.BaseKnotListener
	Name string
	Sensors []model.Sensor
	currentSensor uint
}

const unitFmt = "%s in %s"

func (k *ListenerImpl) EnterStart(ctx *generated.StartContext) {
	if k.Sensors == nil {
		k.currentSensor = 0
	}
}

func (k *ListenerImpl) EnterDefinition(ctx *generated.DefinitionContext) {
	k.Name = ctx.IDENTIFIER().GetText()
}

func (k *ListenerImpl) EnterThingContent(ctx *generated.ThingContentContext) {
	k.Sensors = append(k.Sensors, model.Sensor{ID: k.currentSensor})
}

func (k *ListenerImpl) ExitThingContent(ctx *generated.ThingContentContext) {
	k.Sensors[k.currentSensor].IsSensor = ctx.GetOp().GetTokenType() == generated.KnotParserSENSOR
	k.currentSensor++
}

func (k *ListenerImpl) ExitBoolOpt(ctx *generated.BoolOptContext) {
	k.Sensors[k.currentSensor].Value = "bool"
	k.Sensors[k.currentSensor].DefaultValue = "false"
	k.Sensors[k.currentSensor].Name = ctx.IDENTIFIER().GetText()
	k.Sensors[k.currentSensor].TypeUnit = ctx.GetOp().GetText()
	if ctx.ConfigChanges() != nil{
		k.Sensors[k.currentSensor].Configs = append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.Sensors[k.currentSensor].Configs = append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}
}

func (k *ListenerImpl) ExitNumberOpt(ctx *generated.NumberOptContext) {
	k.Sensors[k.currentSensor].Value = ctx.GetOp().GetText();
	k.Sensors[k.currentSensor].DefaultValue = "0"
	k.Sensors[k.currentSensor].Name = ctx.IDENTIFIER().GetText()
}


func (k *ListenerImpl) ExitConfig(ctx *generated.ConfigContext) {
	if ctx.ConfigChanges() != nil{
		k.Sensors[k.currentSensor].Configs =  append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.Sensors[k.currentSensor].Configs =  append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigUpper(); tmp != nil {
		k.Sensors[k.currentSensor].Configs =  append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotUpper, Value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigLower(); tmp != nil {
		k.Sensors[k.currentSensor].Configs =  append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotLower, Value: tmp.GetNumber().GetText()})
	}
}


func (k *ListenerImpl) ExitBytesOpt(ctx *generated.BytesOptContext) {
	k.Sensors[k.currentSensor].Value = "char[]" // TODO: verify this
	k.Sensors[k.currentSensor].DefaultValue = "NULL"
	k.Sensors[k.currentSensor].Name = ctx.IDENTIFIER().GetText()
	k.Sensors[k.currentSensor].TypeUnit = "command"
	if ctx.ConfigChanges() != nil{
		k.Sensors[k.currentSensor].Configs =  append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.Sensors[k.currentSensor].Configs =  append(k.Sensors[k.currentSensor].Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}
}

func (k *ListenerImpl) ExitUnitTypeOptions(ctx *generated.UnitTypeOptionsContext) {
	fmt.Sscanf(ctx.GetText(), unitFmt, &k.Sensors[k.currentSensor].TypeUnit, &k.Sensors[k.currentSensor].Unit)
}
