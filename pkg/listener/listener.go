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
	sensorIter chan *model.DataItem
	currentSensor *model.DataItem
}

const unitFmt = "%s in %s"

func (k *ListenerImpl) EnterStart(ctx *generated.StartContext) {
	k.sensorIter = make(chan *model.DataItem, len(ctx.GetThings())*255)
	k.Things = make(map[string]*model.Thing, len(ctx.GetThings()))
	for _, thing := range ctx.GetThings() {
		name := thing.GetName().GetText()
		if _, value := k.Things[name]; value {
			log.Fatalf("The thing name \"%s\" is already used", name)
		}

		newThing := model.Thing{Name: name}
		newThing.Sensors = make([]model.DataItem, len(thing.GetSensors()))
		for i, sensor := range thing.GetSensors() {
			newSensor := model.DataItem{}
			newSensor.IsSensor = sensor.GetSensorType().GetTokenType() == generated.KnotParserSENSOR
			newSensor.ID = uint(i)
			newThing.Sensors[i] = newSensor
			k.sensorIter <- &newThing.Sensors[i] // Use channel as iterator
		}
		k.Things[name] = &newThing
	}
}

func (k *ListenerImpl) EnterBoolOpt(ctx *generated.BoolOptContext) {
	sensor := <- k.sensorIter
	sensor.Value = "bool"
	sensor.DefaultValue = "false"
	sensor.Name = ctx.GetName().GetText()
	sensor.TypeUnit = ctx.GetTypeUnit().GetText()
	if ctx.ConfigChanges() != nil{
		sensor.Configs = append(sensor.Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		sensor.Configs = append(sensor.Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}

	k.currentSensor = sensor
}

func (k *ListenerImpl) EnterNumberOpt(ctx *generated.NumberOptContext) {
	sensor := <- k.sensorIter
	sensor.Value = ctx.GetTypeValue().GetText();
	sensor.DefaultValue = "0"
	sensor.Name = ctx.GetName().GetText()
	k.currentSensor = sensor
}

func (k *ListenerImpl) EnterBytesOpt(ctx *generated.BytesOptContext) {
	sensor := <- k.sensorIter
	sensor.Value = "char[]" // TODO: verify this
	sensor.DefaultValue = "NULL"
	sensor.Name = ctx.GetName().GetText()
	sensor.TypeUnit = "command"
	if ctx.ConfigChanges() != nil{
		sensor.Configs =  append(sensor.Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		sensor.Configs =  append(sensor.Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}
	k.currentSensor = sensor
}

func (k *ListenerImpl) ExitConfig(ctx *generated.ConfigContext) {
	if ctx.ConfigChanges() != nil {
		k.currentSensor.Configs =  append(k.currentSensor.Configs, model.Config{Type: model.KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.currentSensor.Configs =  append(k.currentSensor.Configs, model.Config{Type: model.KnotTime, Value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigUpper(); tmp != nil {
		k.currentSensor.Configs =  append(k.currentSensor.Configs, model.Config{Type: model.KnotUpper, Value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigLower(); tmp != nil {
		k.currentSensor.Configs =  append(k.currentSensor.Configs, model.Config{Type: model.KnotLower, Value: tmp.GetNumber().GetText()})
	}
}

func (k *ListenerImpl) ExitUnitTypeOptions(ctx *generated.UnitTypeOptionsContext) {
	fmt.Sscanf(ctx.GetText(), unitFmt, &k.currentSensor.TypeUnit, &k.currentSensor.Unit)
}
