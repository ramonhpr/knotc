package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"./generated"
)

type KnotConfigType int

const (
	KnotChanges = iota
	KnotTime
	KnotUpper
	KnotLower
)

func (k KnotConfigType) String() string {
	return [...]string{"CHANGE", "TIME", "UPPER_THRESHOLD", "LOWER_THRESHOLD"}[k]
}

type knotConfig struct {
	typeConfig KnotConfigType
	value string
}

type sensorType struct {
	name, value, typeUnit, unit string
	isSensor bool
	configs []knotConfig
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
	k.sensors[k.currentSensor].isSensor = ctx.GetOp().GetTokenType() == generated.KnotParserSENSOR
	k.currentSensor++
}

func (k *knotListener) ExitBoolOpt(ctx *generated.BoolOptContext) {
	k.sensors[k.currentSensor].value = "bool"
	k.sensors[k.currentSensor].name = ctx.IDENTIFIER().GetText()
	k.sensors[k.currentSensor].typeUnit = ctx.GetOp().GetText()
	if ctx.ConfigChanges() != nil{
		k.sensors[k.currentSensor].configs = append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.sensors[k.currentSensor].configs = append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotTime, value: tmp.GetNumber().GetText()})
	}
}

func (k *knotListener) ExitNumberOpt(ctx *generated.NumberOptContext) {
	k.sensors[k.currentSensor].value = ctx.GetOp().GetText();
	k.sensors[k.currentSensor].name = ctx.IDENTIFIER().GetText()
}


func (k *knotListener) ExitConfig(ctx *generated.ConfigContext) {
	if ctx.ConfigChanges() != nil{
		k.sensors[k.currentSensor].configs =  append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.sensors[k.currentSensor].configs =  append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotTime, value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigUpper(); tmp != nil {
		k.sensors[k.currentSensor].configs =  append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotUpper, value: tmp.GetNumber().GetText()})
	}

	if tmp := ctx.ConfigLower(); tmp != nil {
		k.sensors[k.currentSensor].configs =  append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotLower, value: tmp.GetNumber().GetText()})
	}
}


func (k *knotListener) ExitBytesOpt(ctx *generated.BytesOptContext) {
	k.sensors[k.currentSensor].value = "bytes"
	k.sensors[k.currentSensor].name = ctx.IDENTIFIER().GetText()
	k.sensors[k.currentSensor].typeUnit = "command"
	if ctx.ConfigChanges() != nil{
		k.sensors[k.currentSensor].configs =  append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotChanges})
	}

	if tmp := ctx.ConfigTime(); tmp != nil {
		k.sensors[k.currentSensor].configs =  append(k.sensors[k.currentSensor].configs, knotConfig{typeConfig: KnotTime, value: tmp.GetNumber().GetText()})
	}
}

const unitFmt = "%s in %s"

func (k *knotListener) ExitUnitTypeOptions(ctx *generated.UnitTypeOptionsContext) {
	fmt.Sscanf(ctx.GetText(), unitFmt, &k.sensors[k.currentSensor].typeUnit, &k.sensors[k.currentSensor].unit)
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
		unit := strings.ToUpper(sensor.typeUnit + "_" +sensor.unit)
		if sensor.unit == "" {
			unit = "NOT_APPLICABLE"
		}
		configParameter := ""
		for i, config := range sensor.configs {
			configParameter += fmt.Sprintf("KNOT_EVT_FLAG_%s", config.typeConfig)
			if config.typeConfig != KnotChanges {
				configParameter += fmt.Sprintf(", %s", config.value)
			}

			if i + 1 != len(sensor.configs) {
				configParameter += fmt.Sprintf(", ")
			} else {
				configParameter += fmt.Sprintf(", NULL")
			}
		}
		setupOut += fmt.Sprintf(setupContentFmt + "\n\n", strings.ToLower(sensor.name), strings.ToUpper(sensor.name), id, strings.ToUpper(sensor.typeUnit), strings.ToUpper(sensor.value), unit, configParameter)
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