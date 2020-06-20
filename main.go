package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"./generated"
)

type sensorType struct {
	name, value, typeUnit, unit string
	isSensor bool
}

type knotListener struct {
	*generated.BaseKnotListener
	sensors []sensorType
	currentSensor int
	name string

}


func (k *knotListener) EnterStart(ctx *generated.StartContext) {
	if k.sensors == nil {
		k.currentSensor = 0
	}
}


func (k *knotListener) EnterDefinition(ctx *generated.DefinitionContext) {
	k.name = ctx.IDENTIFIER().GetText()
}

func (k *knotListener) EnterThingContent(ctx *generated.ThingContentContext) {
	k.sensors = append(k.sensors, sensorType{})
}

func (k *knotListener) ExitThingContent(ctx *generated.ThingContentContext) {
	k.sensors[k.currentSensor].name = ctx.IDENTIFIER().GetText()
	k.sensors[k.currentSensor].isSensor = ctx.GetOp().GetTokenType() == generated.KnotParserSENSOR
	k.currentSensor++
}

func (k *knotListener) ExitTemperatureUnits(ctx *generated.TemperatureUnitsContext) {
	k.sensors[k.currentSensor].unit = ctx.GetOp().GetText();
}

func (k *knotListener) ExitValueOptions(ctx *generated.ValueOptionsContext) {
	k.sensors[k.currentSensor].value = ctx.GetOp().GetText();
}

func compileToZephyr(l knotListener, path string) {
	data, err := ioutil.ReadFile(path + "headers_template.c")
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(data), l.name)
	fmt.Println()

	define, err := ioutil.ReadFile(path + "defines_template.c")
	if err != nil {
		panic(err)
	}

	for _, sensor := range l.sensors {
		fmt.Printf(string(define), strings.ToUpper(sensor.name))
		fmt.Println()
	}

	fmt.Println()

	variable, err := ioutil.ReadFile(path + "variables_template.c")
	if err != nil {
		panic(err)
	}


	for _, sensor := range l.sensors {
		if !sensor.isSensor {
			fmt.Printf(string(variable), strings.ToLower(sensor.name), sensor.value)
			fmt.Println()
		}
	}

	cb, err := ioutil.ReadFile(path + "callback_write_template.c")
	if err != nil {
		panic(err)
	}

	for _, sensor := range l.sensors {
		if !sensor.isSensor {
			fmt.Printf(string(cb), strings.ToLower(sensor.name))
			fmt.Println()
		}
	}

	last, err := ioutil.ReadFile(path + "setup_loop_template.c")
	if err != nil {
		panic(err)
	}

	for id, sensor := range l.sensors {
		if !sensor.isSensor {
			unit := sensor.unit
			if sensor.unit == "" {
				unit = "NOT_APPLICABLE"
			}
			fmt.Printf(string(last), strings.ToLower(sensor.name), strings.ToUpper(sensor.name), id, strings.ToUpper(sensor.typeUnit), strings.ToUpper(sensor.value), unit)
			fmt.Println()
		}
	}

	fmt.Println("------prj.conf-----")
	data, err = ioutil.ReadFile(path + "prj.conf")
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(data), l.name, len(l.sensors))
}

func main() {
	// Setup the input
	data, err := ioutil.ReadFile("./blink.knot")
	if err != nil {
		panic(err)
	}

	is := antlr.NewInputStream(string(data))

	// Create the Lexer
	lexer := generated.NewKnotLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := generated.NewKnotParser(stream)

	var listener knotListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	fmt.Println(listener.sensors)

	compileToZephyr(listener, "zephyrProject/src/")
}