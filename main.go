package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

const unitFmt = "%s in %s"

func (k *knotListener) ExitUnitTypeOptions(ctx *generated.UnitTypeOptionsContext) {
	fmt.Sscanf(ctx.GetText(), unitFmt, &k.sensors[k.currentSensor].typeUnit, &k.sensors[k.currentSensor].unit)
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
		defaultValue := ""
		value := sensor.value
		switch sensor.value {
		case "int":
			defaultValue = "0"
		case "bool":
			defaultValue = "true"
		case "float":
			defaultValue = "0.0"
		case "bytes":
			defaultValue = "NULL"
			value = "char[]"
		}
		fmt.Printf(string(variable), strings.ToLower(sensor.name), value, defaultValue)
		fmt.Println()
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

	setupContent, err := ioutil.ReadFile(path + "setup_content.c")
	if err != nil {
		panic(err)
	}

	setupContentFmt := string(setupContent)
	setupOut := ""

	for id, sensor := range l.sensors {
		unit := sensor.unit
		if sensor.unit == "" {
			unit = "NOT_APPLICABLE"
		}
		setupOut += fmt.Sprintf(setupContentFmt + "\n\n", strings.ToLower(sensor.name), strings.ToUpper(sensor.name), id, strings.ToUpper(sensor.typeUnit), strings.ToUpper(sensor.value), unit)
	}

	last, err := ioutil.ReadFile(path + "setup_loop_template.c")
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(last), setupOut)
	fmt.Println()

	fmt.Println("------prj.conf-----")
	data, err = ioutil.ReadFile(path + "prj.conf")
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(data), l.name, len(l.sensors))
}

type exitErrorListener struct {
	*antlr.ConsoleErrorListener
}

func (c *exitErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	os.Exit(1)
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
	p.AddErrorListener(new(exitErrorListener))

	var listener knotListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
	fmt.Println(listener.sensors)

	compileToZephyr(listener, "zephyrProject/src/")
}