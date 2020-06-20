// Code generated from Knot.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // Knot
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 84, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 6, 2, 22, 10, 2, 13, 2, 14, 2,
	23, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 32, 10, 3, 13, 3, 14, 3,
	33, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 5, 6, 72, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 6, 3, 2,
	16, 17, 3, 2, 26, 29, 3, 2, 8, 10, 3, 2, 12, 14, 2, 99, 2, 21, 3, 2, 2,
	2, 4, 27, 3, 2, 2, 2, 6, 37, 3, 2, 2, 2, 8, 45, 3, 2, 2, 2, 10, 71, 3,
	2, 2, 2, 12, 73, 3, 2, 2, 2, 14, 76, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18,
	81, 3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2,
	2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 26,
	7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 28, 7, 15, 2, 2, 28, 29, 7, 53, 2, 2,
	29, 31, 7, 3, 2, 2, 30, 32, 5, 6, 4, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3,
	2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35,
	36, 7, 4, 2, 2, 36, 5, 3, 2, 2, 2, 37, 38, 9, 2, 2, 2, 38, 39, 5, 8, 5,
	2, 39, 40, 7, 53, 2, 2, 40, 41, 7, 5, 2, 2, 41, 42, 5, 10, 6, 2, 42, 43,
	7, 6, 2, 2, 43, 44, 7, 21, 2, 2, 44, 7, 3, 2, 2, 2, 45, 46, 9, 3, 2, 2,
	46, 9, 3, 2, 2, 2, 47, 72, 5, 12, 7, 2, 48, 72, 7, 30, 2, 2, 49, 72, 7,
	31, 2, 2, 50, 72, 7, 32, 2, 2, 51, 72, 5, 16, 9, 2, 52, 72, 7, 34, 2, 2,
	53, 72, 7, 35, 2, 2, 54, 72, 7, 36, 2, 2, 55, 72, 7, 37, 2, 2, 56, 72,
	7, 38, 2, 2, 57, 72, 7, 39, 2, 2, 58, 72, 7, 40, 2, 2, 59, 72, 7, 41, 2,
	2, 60, 72, 7, 42, 2, 2, 61, 72, 7, 43, 2, 2, 62, 72, 7, 44, 2, 2, 63, 72,
	7, 45, 2, 2, 64, 72, 7, 46, 2, 2, 65, 72, 7, 47, 2, 2, 66, 72, 7, 48, 2,
	2, 67, 72, 7, 49, 2, 2, 68, 72, 7, 50, 2, 2, 69, 72, 7, 51, 2, 2, 70, 72,
	7, 52, 2, 2, 71, 47, 3, 2, 2, 2, 71, 48, 3, 2, 2, 2, 71, 49, 3, 2, 2, 2,
	71, 50, 3, 2, 2, 2, 71, 51, 3, 2, 2, 2, 71, 52, 3, 2, 2, 2, 71, 53, 3,
	2, 2, 2, 71, 54, 3, 2, 2, 2, 71, 55, 3, 2, 2, 2, 71, 56, 3, 2, 2, 2, 71,
	57, 3, 2, 2, 2, 71, 58, 3, 2, 2, 2, 71, 59, 3, 2, 2, 2, 71, 60, 3, 2, 2,
	2, 71, 61, 3, 2, 2, 2, 71, 62, 3, 2, 2, 2, 71, 63, 3, 2, 2, 2, 71, 64,
	3, 2, 2, 2, 71, 65, 3, 2, 2, 2, 71, 66, 3, 2, 2, 2, 71, 67, 3, 2, 2, 2,
	71, 68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 11, 3,
	2, 2, 2, 73, 74, 7, 7, 2, 2, 74, 75, 5, 14, 8, 2, 75, 13, 3, 2, 2, 2, 76,
	77, 9, 4, 2, 2, 77, 15, 3, 2, 2, 2, 78, 79, 7, 11, 2, 2, 79, 80, 5, 18,
	10, 2, 80, 17, 3, 2, 2, 2, 81, 82, 9, 5, 2, 2, 82, 19, 3, 2, 2, 2, 5, 23,
	33, 71,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'('", "')'", "'voltage in '", "'V'", "'mV'", "'kV'",
	"'temperature in '", "'F'", "'C'", "'K'", "'thing'", "'sensor'", "'actuator'",
	"'value'", "'unit'", "'type'", "';'", "':'", "", "", "", "'bool'", "'int'",
	"'float'", "'bytes'", "'current'", "'resistance'", "'power'", "'temperature'",
	"'relativehumidity'", "'luminosity'", "'time'", "'mass'", "'pressure'",
	"'distance'", "'angle'", "'volume'", "'area'", "'rain'", "'density'", "'latitude'",
	"'longitude'", "'speed'", "'volumeflow'", "'energy'", "'switch'", "'presence'",
	"'command'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "THING", "SENSOR",
	"ACTUATOR", "VALUE", "UNIT", "TYPE", "END_CHAR", "SEPARATOR", "WS", "COMMENT",
	"LINE_COMMENT", "BOOL", "INT", "FLOAT", "BYTES", "CURRENT", "RESISTANCE",
	"POWER", "TEMPERATURE", "RELATIVEHUMIDITY", "LUMINOSITY", "TIME", "MASS",
	"PRESSURE", "DISTANCE", "ANGLE", "VOLUME", "AREA", "RAIN", "DENSITY", "LATITUDE",
	"LONGITUDE", "SPEED", "VOLUMEFLOW", "ENERGY", "SWITCH", "PRESENCE", "COMMAND",
	"IDENTIFIER", "ID",
}

var ruleNames = []string{
	"start", "definition", "thingContent", "valueOptions", "unitTypeOptions",
	"voltage", "voltagesUnits", "temperature", "temperatureUnits",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type KnotParser struct {
	*antlr.BaseParser
}

func NewKnotParser(input antlr.TokenStream) *KnotParser {
	this := new(KnotParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Knot.g4"

	return this
}

// KnotParser tokens.
const (
	KnotParserEOF              = antlr.TokenEOF
	KnotParserT__0             = 1
	KnotParserT__1             = 2
	KnotParserT__2             = 3
	KnotParserT__3             = 4
	KnotParserT__4             = 5
	KnotParserT__5             = 6
	KnotParserT__6             = 7
	KnotParserT__7             = 8
	KnotParserT__8             = 9
	KnotParserT__9             = 10
	KnotParserT__10            = 11
	KnotParserT__11            = 12
	KnotParserTHING            = 13
	KnotParserSENSOR           = 14
	KnotParserACTUATOR         = 15
	KnotParserVALUE            = 16
	KnotParserUNIT             = 17
	KnotParserTYPE             = 18
	KnotParserEND_CHAR         = 19
	KnotParserSEPARATOR        = 20
	KnotParserWS               = 21
	KnotParserCOMMENT          = 22
	KnotParserLINE_COMMENT     = 23
	KnotParserBOOL             = 24
	KnotParserINT              = 25
	KnotParserFLOAT            = 26
	KnotParserBYTES            = 27
	KnotParserCURRENT          = 28
	KnotParserRESISTANCE       = 29
	KnotParserPOWER            = 30
	KnotParserTEMPERATURE      = 31
	KnotParserRELATIVEHUMIDITY = 32
	KnotParserLUMINOSITY       = 33
	KnotParserTIME             = 34
	KnotParserMASS             = 35
	KnotParserPRESSURE         = 36
	KnotParserDISTANCE         = 37
	KnotParserANGLE            = 38
	KnotParserVOLUME           = 39
	KnotParserAREA             = 40
	KnotParserRAIN             = 41
	KnotParserDENSITY          = 42
	KnotParserLATITUDE         = 43
	KnotParserLONGITUDE        = 44
	KnotParserSPEED            = 45
	KnotParserVOLUMEFLOW       = 46
	KnotParserENERGY           = 47
	KnotParserSWITCH           = 48
	KnotParserPRESENCE         = 49
	KnotParserCOMMAND          = 50
	KnotParserIDENTIFIER       = 51
	KnotParserID               = 52
)

// KnotParser rules.
const (
	KnotParserRULE_start            = 0
	KnotParserRULE_definition       = 1
	KnotParserRULE_thingContent     = 2
	KnotParserRULE_valueOptions     = 3
	KnotParserRULE_unitTypeOptions  = 4
	KnotParserRULE_voltage          = 5
	KnotParserRULE_voltagesUnits    = 6
	KnotParserRULE_temperature      = 7
	KnotParserRULE_temperatureUnits = 8
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(KnotParserEOF, 0)
}

func (s *StartContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *StartContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KnotParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserTHING {
		{
			p.SetState(18)
			p.Definition()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(KnotParserEOF)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) THING() antlr.TerminalNode {
	return s.GetToken(KnotParserTHING, 0)
}

func (s *DefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KnotParserIDENTIFIER, 0)
}

func (s *DefinitionContext) AllThingContent() []IThingContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThingContentContext)(nil)).Elem())
	var tst = make([]IThingContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThingContentContext)
		}
	}

	return tst
}

func (s *DefinitionContext) ThingContent(i int) IThingContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThingContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThingContentContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KnotParserRULE_definition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(KnotParserTHING)
	}
	{
		p.SetState(26)
		p.Match(KnotParserIDENTIFIER)
	}
	{
		p.SetState(27)
		p.Match(KnotParserT__0)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserSENSOR || _la == KnotParserACTUATOR {
		{
			p.SetState(28)
			p.ThingContent()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(KnotParserT__1)
	}

	return localctx
}

// IThingContentContext is an interface to support dynamic dispatch.
type IThingContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsThingContentContext differentiates from other interfaces.
	IsThingContentContext()
}

type ThingContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyThingContentContext() *ThingContentContext {
	var p = new(ThingContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_thingContent
	return p
}

func (*ThingContentContext) IsThingContentContext() {}

func NewThingContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThingContentContext {
	var p = new(ThingContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_thingContent

	return p
}

func (s *ThingContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ThingContentContext) GetOp() antlr.Token { return s.op }

func (s *ThingContentContext) SetOp(v antlr.Token) { s.op = v }

func (s *ThingContentContext) ValueOptions() IValueOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueOptionsContext)
}

func (s *ThingContentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KnotParserIDENTIFIER, 0)
}

func (s *ThingContentContext) UnitTypeOptions() IUnitTypeOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitTypeOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnitTypeOptionsContext)
}

func (s *ThingContentContext) END_CHAR() antlr.TerminalNode {
	return s.GetToken(KnotParserEND_CHAR, 0)
}

func (s *ThingContentContext) SENSOR() antlr.TerminalNode {
	return s.GetToken(KnotParserSENSOR, 0)
}

func (s *ThingContentContext) ACTUATOR() antlr.TerminalNode {
	return s.GetToken(KnotParserACTUATOR, 0)
}

func (s *ThingContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThingContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThingContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterThingContent(s)
	}
}

func (s *ThingContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitThingContent(s)
	}
}

func (s *ThingContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitThingContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) ThingContent() (localctx IThingContentContext) {
	localctx = NewThingContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KnotParserRULE_thingContent)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ThingContentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == KnotParserSENSOR || _la == KnotParserACTUATOR) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ThingContentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(36)
		p.ValueOptions()
	}
	{
		p.SetState(37)
		p.Match(KnotParserIDENTIFIER)
	}
	{
		p.SetState(38)
		p.Match(KnotParserT__2)
	}
	{
		p.SetState(39)
		p.UnitTypeOptions()
	}
	{
		p.SetState(40)
		p.Match(KnotParserT__3)
	}
	{
		p.SetState(41)
		p.Match(KnotParserEND_CHAR)
	}

	return localctx
}

// IValueOptionsContext is an interface to support dynamic dispatch.
type IValueOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsValueOptionsContext differentiates from other interfaces.
	IsValueOptionsContext()
}

type ValueOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyValueOptionsContext() *ValueOptionsContext {
	var p = new(ValueOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_valueOptions
	return p
}

func (*ValueOptionsContext) IsValueOptionsContext() {}

func NewValueOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueOptionsContext {
	var p = new(ValueOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_valueOptions

	return p
}

func (s *ValueOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueOptionsContext) GetOp() antlr.Token { return s.op }

func (s *ValueOptionsContext) SetOp(v antlr.Token) { s.op = v }

func (s *ValueOptionsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(KnotParserBOOL, 0)
}

func (s *ValueOptionsContext) INT() antlr.TerminalNode {
	return s.GetToken(KnotParserINT, 0)
}

func (s *ValueOptionsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(KnotParserFLOAT, 0)
}

func (s *ValueOptionsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(KnotParserBYTES, 0)
}

func (s *ValueOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterValueOptions(s)
	}
}

func (s *ValueOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitValueOptions(s)
	}
}

func (s *ValueOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitValueOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) ValueOptions() (localctx IValueOptionsContext) {
	localctx = NewValueOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KnotParserRULE_valueOptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ValueOptionsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserBOOL)|(1<<KnotParserINT)|(1<<KnotParserFLOAT)|(1<<KnotParserBYTES))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ValueOptionsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnitTypeOptionsContext is an interface to support dynamic dispatch.
type IUnitTypeOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitTypeOptionsContext differentiates from other interfaces.
	IsUnitTypeOptionsContext()
}

type UnitTypeOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitTypeOptionsContext() *UnitTypeOptionsContext {
	var p = new(UnitTypeOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_unitTypeOptions
	return p
}

func (*UnitTypeOptionsContext) IsUnitTypeOptionsContext() {}

func NewUnitTypeOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitTypeOptionsContext {
	var p = new(UnitTypeOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_unitTypeOptions

	return p
}

func (s *UnitTypeOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitTypeOptionsContext) Voltage() IVoltageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVoltageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVoltageContext)
}

func (s *UnitTypeOptionsContext) CURRENT() antlr.TerminalNode {
	return s.GetToken(KnotParserCURRENT, 0)
}

func (s *UnitTypeOptionsContext) RESISTANCE() antlr.TerminalNode {
	return s.GetToken(KnotParserRESISTANCE, 0)
}

func (s *UnitTypeOptionsContext) POWER() antlr.TerminalNode {
	return s.GetToken(KnotParserPOWER, 0)
}

func (s *UnitTypeOptionsContext) Temperature() ITemperatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemperatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemperatureContext)
}

func (s *UnitTypeOptionsContext) RELATIVEHUMIDITY() antlr.TerminalNode {
	return s.GetToken(KnotParserRELATIVEHUMIDITY, 0)
}

func (s *UnitTypeOptionsContext) LUMINOSITY() antlr.TerminalNode {
	return s.GetToken(KnotParserLUMINOSITY, 0)
}

func (s *UnitTypeOptionsContext) TIME() antlr.TerminalNode {
	return s.GetToken(KnotParserTIME, 0)
}

func (s *UnitTypeOptionsContext) MASS() antlr.TerminalNode {
	return s.GetToken(KnotParserMASS, 0)
}

func (s *UnitTypeOptionsContext) PRESSURE() antlr.TerminalNode {
	return s.GetToken(KnotParserPRESSURE, 0)
}

func (s *UnitTypeOptionsContext) DISTANCE() antlr.TerminalNode {
	return s.GetToken(KnotParserDISTANCE, 0)
}

func (s *UnitTypeOptionsContext) ANGLE() antlr.TerminalNode {
	return s.GetToken(KnotParserANGLE, 0)
}

func (s *UnitTypeOptionsContext) VOLUME() antlr.TerminalNode {
	return s.GetToken(KnotParserVOLUME, 0)
}

func (s *UnitTypeOptionsContext) AREA() antlr.TerminalNode {
	return s.GetToken(KnotParserAREA, 0)
}

func (s *UnitTypeOptionsContext) RAIN() antlr.TerminalNode {
	return s.GetToken(KnotParserRAIN, 0)
}

func (s *UnitTypeOptionsContext) DENSITY() antlr.TerminalNode {
	return s.GetToken(KnotParserDENSITY, 0)
}

func (s *UnitTypeOptionsContext) LATITUDE() antlr.TerminalNode {
	return s.GetToken(KnotParserLATITUDE, 0)
}

func (s *UnitTypeOptionsContext) LONGITUDE() antlr.TerminalNode {
	return s.GetToken(KnotParserLONGITUDE, 0)
}

func (s *UnitTypeOptionsContext) SPEED() antlr.TerminalNode {
	return s.GetToken(KnotParserSPEED, 0)
}

func (s *UnitTypeOptionsContext) VOLUMEFLOW() antlr.TerminalNode {
	return s.GetToken(KnotParserVOLUMEFLOW, 0)
}

func (s *UnitTypeOptionsContext) ENERGY() antlr.TerminalNode {
	return s.GetToken(KnotParserENERGY, 0)
}

func (s *UnitTypeOptionsContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(KnotParserSWITCH, 0)
}

func (s *UnitTypeOptionsContext) PRESENCE() antlr.TerminalNode {
	return s.GetToken(KnotParserPRESENCE, 0)
}

func (s *UnitTypeOptionsContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(KnotParserCOMMAND, 0)
}

func (s *UnitTypeOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitTypeOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitTypeOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterUnitTypeOptions(s)
	}
}

func (s *UnitTypeOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitUnitTypeOptions(s)
	}
}

func (s *UnitTypeOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitUnitTypeOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) UnitTypeOptions() (localctx IUnitTypeOptionsContext) {
	localctx = NewUnitTypeOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KnotParserRULE_unitTypeOptions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KnotParserT__4:
		{
			p.SetState(45)
			p.Voltage()
		}

	case KnotParserCURRENT:
		{
			p.SetState(46)
			p.Match(KnotParserCURRENT)
		}

	case KnotParserRESISTANCE:
		{
			p.SetState(47)
			p.Match(KnotParserRESISTANCE)
		}

	case KnotParserPOWER:
		{
			p.SetState(48)
			p.Match(KnotParserPOWER)
		}

	case KnotParserT__8:
		{
			p.SetState(49)
			p.Temperature()
		}

	case KnotParserRELATIVEHUMIDITY:
		{
			p.SetState(50)
			p.Match(KnotParserRELATIVEHUMIDITY)
		}

	case KnotParserLUMINOSITY:
		{
			p.SetState(51)
			p.Match(KnotParserLUMINOSITY)
		}

	case KnotParserTIME:
		{
			p.SetState(52)
			p.Match(KnotParserTIME)
		}

	case KnotParserMASS:
		{
			p.SetState(53)
			p.Match(KnotParserMASS)
		}

	case KnotParserPRESSURE:
		{
			p.SetState(54)
			p.Match(KnotParserPRESSURE)
		}

	case KnotParserDISTANCE:
		{
			p.SetState(55)
			p.Match(KnotParserDISTANCE)
		}

	case KnotParserANGLE:
		{
			p.SetState(56)
			p.Match(KnotParserANGLE)
		}

	case KnotParserVOLUME:
		{
			p.SetState(57)
			p.Match(KnotParserVOLUME)
		}

	case KnotParserAREA:
		{
			p.SetState(58)
			p.Match(KnotParserAREA)
		}

	case KnotParserRAIN:
		{
			p.SetState(59)
			p.Match(KnotParserRAIN)
		}

	case KnotParserDENSITY:
		{
			p.SetState(60)
			p.Match(KnotParserDENSITY)
		}

	case KnotParserLATITUDE:
		{
			p.SetState(61)
			p.Match(KnotParserLATITUDE)
		}

	case KnotParserLONGITUDE:
		{
			p.SetState(62)
			p.Match(KnotParserLONGITUDE)
		}

	case KnotParserSPEED:
		{
			p.SetState(63)
			p.Match(KnotParserSPEED)
		}

	case KnotParserVOLUMEFLOW:
		{
			p.SetState(64)
			p.Match(KnotParserVOLUMEFLOW)
		}

	case KnotParserENERGY:
		{
			p.SetState(65)
			p.Match(KnotParserENERGY)
		}

	case KnotParserSWITCH:
		{
			p.SetState(66)
			p.Match(KnotParserSWITCH)
		}

	case KnotParserPRESENCE:
		{
			p.SetState(67)
			p.Match(KnotParserPRESENCE)
		}

	case KnotParserCOMMAND:
		{
			p.SetState(68)
			p.Match(KnotParserCOMMAND)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVoltageContext is an interface to support dynamic dispatch.
type IVoltageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVoltageContext differentiates from other interfaces.
	IsVoltageContext()
}

type VoltageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoltageContext() *VoltageContext {
	var p = new(VoltageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_voltage
	return p
}

func (*VoltageContext) IsVoltageContext() {}

func NewVoltageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoltageContext {
	var p = new(VoltageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_voltage

	return p
}

func (s *VoltageContext) GetParser() antlr.Parser { return s.parser }

func (s *VoltageContext) VoltagesUnits() IVoltagesUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVoltagesUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVoltagesUnitsContext)
}

func (s *VoltageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoltageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoltageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterVoltage(s)
	}
}

func (s *VoltageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitVoltage(s)
	}
}

func (s *VoltageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitVoltage(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Voltage() (localctx IVoltageContext) {
	localctx = NewVoltageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KnotParserRULE_voltage)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(KnotParserT__4)
	}
	{
		p.SetState(72)
		p.VoltagesUnits()
	}

	return localctx
}

// IVoltagesUnitsContext is an interface to support dynamic dispatch.
type IVoltagesUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsVoltagesUnitsContext differentiates from other interfaces.
	IsVoltagesUnitsContext()
}

type VoltagesUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyVoltagesUnitsContext() *VoltagesUnitsContext {
	var p = new(VoltagesUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_voltagesUnits
	return p
}

func (*VoltagesUnitsContext) IsVoltagesUnitsContext() {}

func NewVoltagesUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoltagesUnitsContext {
	var p = new(VoltagesUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_voltagesUnits

	return p
}

func (s *VoltagesUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *VoltagesUnitsContext) GetOp() antlr.Token { return s.op }

func (s *VoltagesUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *VoltagesUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoltagesUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoltagesUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterVoltagesUnits(s)
	}
}

func (s *VoltagesUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitVoltagesUnits(s)
	}
}

func (s *VoltagesUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitVoltagesUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) VoltagesUnits() (localctx IVoltagesUnitsContext) {
	localctx = NewVoltagesUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KnotParserRULE_voltagesUnits)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VoltagesUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserT__5)|(1<<KnotParserT__6)|(1<<KnotParserT__7))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VoltagesUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITemperatureContext is an interface to support dynamic dispatch.
type ITemperatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemperatureContext differentiates from other interfaces.
	IsTemperatureContext()
}

type TemperatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemperatureContext() *TemperatureContext {
	var p = new(TemperatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_temperature
	return p
}

func (*TemperatureContext) IsTemperatureContext() {}

func NewTemperatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemperatureContext {
	var p = new(TemperatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_temperature

	return p
}

func (s *TemperatureContext) GetParser() antlr.Parser { return s.parser }

func (s *TemperatureContext) TemperatureUnits() ITemperatureUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemperatureUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemperatureUnitsContext)
}

func (s *TemperatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemperatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemperatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterTemperature(s)
	}
}

func (s *TemperatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitTemperature(s)
	}
}

func (s *TemperatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitTemperature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Temperature() (localctx ITemperatureContext) {
	localctx = NewTemperatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KnotParserRULE_temperature)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(KnotParserT__8)
	}
	{
		p.SetState(77)
		p.TemperatureUnits()
	}

	return localctx
}

// ITemperatureUnitsContext is an interface to support dynamic dispatch.
type ITemperatureUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsTemperatureUnitsContext differentiates from other interfaces.
	IsTemperatureUnitsContext()
}

type TemperatureUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyTemperatureUnitsContext() *TemperatureUnitsContext {
	var p = new(TemperatureUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_temperatureUnits
	return p
}

func (*TemperatureUnitsContext) IsTemperatureUnitsContext() {}

func NewTemperatureUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemperatureUnitsContext {
	var p = new(TemperatureUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_temperatureUnits

	return p
}

func (s *TemperatureUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *TemperatureUnitsContext) GetOp() antlr.Token { return s.op }

func (s *TemperatureUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *TemperatureUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemperatureUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemperatureUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterTemperatureUnits(s)
	}
}

func (s *TemperatureUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitTemperatureUnits(s)
	}
}

func (s *TemperatureUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitTemperatureUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) TemperatureUnits() (localctx ITemperatureUnitsContext) {
	localctx = NewTemperatureUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KnotParserRULE_temperatureUnits)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TemperatureUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserT__9)|(1<<KnotParserT__10)|(1<<KnotParserT__11))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TemperatureUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
