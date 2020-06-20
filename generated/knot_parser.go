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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 53, 112,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 6, 2, 29, 10, 2, 13, 2, 14, 2, 30, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 6, 3, 38, 10, 3, 13, 3, 14, 3, 39, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 6, 4, 48, 10, 4, 13, 4, 14, 4, 49, 3, 4, 3, 4, 3, 4, 6, 4,
	55, 10, 4, 13, 4, 14, 4, 56, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 100,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 6,
	3, 2, 15, 16, 3, 2, 25, 28, 3, 2, 9, 11, 3, 2, 12, 14, 2, 126, 2, 26, 3,
	2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 63,
	3, 2, 2, 2, 12, 68, 3, 2, 2, 2, 14, 70, 3, 2, 2, 2, 16, 99, 3, 2, 2, 2,
	18, 101, 3, 2, 2, 2, 20, 104, 3, 2, 2, 2, 22, 107, 3, 2, 2, 2, 24, 109,
	3, 2, 2, 2, 26, 28, 5, 4, 3, 2, 27, 29, 5, 6, 4, 2, 28, 27, 3, 2, 2, 2,
	29, 30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 3,
	2, 2, 2, 32, 33, 7, 2, 2, 3, 33, 3, 3, 2, 2, 2, 34, 35, 7, 3, 2, 2, 35,
	37, 7, 4, 2, 2, 36, 38, 7, 52, 2, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2,
	2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 42,
	7, 4, 2, 2, 42, 43, 7, 20, 2, 2, 43, 5, 3, 2, 2, 2, 44, 45, 9, 2, 2, 2,
	45, 47, 7, 4, 2, 2, 46, 48, 7, 52, 2, 2, 47, 46, 3, 2, 2, 2, 48, 49, 3,
	2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51,
	52, 7, 4, 2, 2, 52, 54, 7, 5, 2, 2, 53, 55, 5, 8, 5, 2, 54, 53, 3, 2, 2,
	2, 55, 56, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 59, 7, 6, 2, 2, 59, 7, 3, 2, 2, 2, 60, 61, 5, 10, 6, 2,
	61, 62, 5, 14, 8, 2, 62, 9, 3, 2, 2, 2, 63, 64, 7, 17, 2, 2, 64, 65, 7,
	21, 2, 2, 65, 66, 5, 12, 7, 2, 66, 67, 7, 20, 2, 2, 67, 11, 3, 2, 2, 2,
	68, 69, 9, 3, 2, 2, 69, 13, 3, 2, 2, 2, 70, 71, 7, 19, 2, 2, 71, 72, 7,
	21, 2, 2, 72, 73, 5, 16, 9, 2, 73, 74, 7, 20, 2, 2, 74, 15, 3, 2, 2, 2,
	75, 100, 5, 18, 10, 2, 76, 100, 7, 29, 2, 2, 77, 100, 7, 30, 2, 2, 78,
	100, 7, 31, 2, 2, 79, 100, 5, 20, 11, 2, 80, 100, 7, 33, 2, 2, 81, 100,
	7, 34, 2, 2, 82, 100, 7, 35, 2, 2, 83, 100, 7, 36, 2, 2, 84, 100, 7, 37,
	2, 2, 85, 100, 7, 38, 2, 2, 86, 100, 7, 39, 2, 2, 87, 100, 7, 40, 2, 2,
	88, 100, 7, 41, 2, 2, 89, 100, 7, 42, 2, 2, 90, 100, 7, 43, 2, 2, 91, 100,
	7, 44, 2, 2, 92, 100, 7, 45, 2, 2, 93, 100, 7, 46, 2, 2, 94, 100, 7, 47,
	2, 2, 95, 100, 7, 48, 2, 2, 96, 100, 7, 49, 2, 2, 97, 100, 7, 50, 2, 2,
	98, 100, 7, 51, 2, 2, 99, 75, 3, 2, 2, 2, 99, 76, 3, 2, 2, 2, 99, 77, 3,
	2, 2, 2, 99, 78, 3, 2, 2, 2, 99, 79, 3, 2, 2, 2, 99, 80, 3, 2, 2, 2, 99,
	81, 3, 2, 2, 2, 99, 82, 3, 2, 2, 2, 99, 83, 3, 2, 2, 2, 99, 84, 3, 2, 2,
	2, 99, 85, 3, 2, 2, 2, 99, 86, 3, 2, 2, 2, 99, 87, 3, 2, 2, 2, 99, 88,
	3, 2, 2, 2, 99, 89, 3, 2, 2, 2, 99, 90, 3, 2, 2, 2, 99, 91, 3, 2, 2, 2,
	99, 92, 3, 2, 2, 2, 99, 93, 3, 2, 2, 2, 99, 94, 3, 2, 2, 2, 99, 95, 3,
	2, 2, 2, 99, 96, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 98, 3, 2, 2, 2, 100,
	17, 3, 2, 2, 2, 101, 102, 7, 7, 2, 2, 102, 103, 5, 22, 12, 2, 103, 19,
	3, 2, 2, 2, 104, 105, 7, 8, 2, 2, 105, 106, 5, 24, 13, 2, 106, 21, 3, 2,
	2, 2, 107, 108, 9, 4, 2, 2, 108, 23, 3, 2, 2, 2, 109, 110, 9, 5, 2, 2,
	110, 25, 3, 2, 2, 2, 7, 30, 39, 49, 56, 99,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'name'", "'\"'", "'{'", "'}'", "'voltage in '", "'temperature in '",
	"'V'", "'mV'", "'kV'", "'F'", "'C'", "'K'", "'sensor'", "'actuator'", "'value'",
	"'unit'", "'type'", "';'", "':'", "", "", "", "'bool'", "'int'", "'float'",
	"'bytes'", "'current'", "'resistance'", "'power'", "'temperature'", "'relativehumidity'",
	"'luminosity'", "'time'", "'mass'", "'pressure'", "'distance'", "'angle'",
	"'volume'", "'area'", "'rain'", "'density'", "'latitude'", "'longitude'",
	"'speed'", "'volumeflow'", "'energy'", "'switch'", "'presence'", "'command'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "SENSOR", "ACTUATOR",
	"VALUE", "UNIT", "TYPE", "END_CHAR", "SEPARATOR", "WS", "COMMENT", "LINE_COMMENT",
	"BOOL", "INT", "FLOAT", "BYTES", "CURRENT", "RESISTANCE", "POWER", "TEMPERATURE",
	"RELATIVEHUMIDITY", "LUMINOSITY", "TIME", "MASS", "PRESSURE", "DISTANCE",
	"ANGLE", "VOLUME", "AREA", "RAIN", "DENSITY", "LATITUDE", "LONGITUDE",
	"SPEED", "VOLUMEFLOW", "ENERGY", "SWITCH", "PRESENCE", "COMMAND", "IDENTIFIER",
	"ID",
}

var ruleNames = []string{
	"start", "name", "definition", "innerContent", "value", "valueOptions",
	"typeRule", "typeOptions", "voltage", "temperature", "voltagesUnits", "temperatureUnits",
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
	KnotParserSENSOR           = 13
	KnotParserACTUATOR         = 14
	KnotParserVALUE            = 15
	KnotParserUNIT             = 16
	KnotParserTYPE             = 17
	KnotParserEND_CHAR         = 18
	KnotParserSEPARATOR        = 19
	KnotParserWS               = 20
	KnotParserCOMMENT          = 21
	KnotParserLINE_COMMENT     = 22
	KnotParserBOOL             = 23
	KnotParserINT              = 24
	KnotParserFLOAT            = 25
	KnotParserBYTES            = 26
	KnotParserCURRENT          = 27
	KnotParserRESISTANCE       = 28
	KnotParserPOWER            = 29
	KnotParserTEMPERATURE      = 30
	KnotParserRELATIVEHUMIDITY = 31
	KnotParserLUMINOSITY       = 32
	KnotParserTIME             = 33
	KnotParserMASS             = 34
	KnotParserPRESSURE         = 35
	KnotParserDISTANCE         = 36
	KnotParserANGLE            = 37
	KnotParserVOLUME           = 38
	KnotParserAREA             = 39
	KnotParserRAIN             = 40
	KnotParserDENSITY          = 41
	KnotParserLATITUDE         = 42
	KnotParserLONGITUDE        = 43
	KnotParserSPEED            = 44
	KnotParserVOLUMEFLOW       = 45
	KnotParserENERGY           = 46
	KnotParserSWITCH           = 47
	KnotParserPRESENCE         = 48
	KnotParserCOMMAND          = 49
	KnotParserIDENTIFIER       = 50
	KnotParserID               = 51
)

// KnotParser rules.
const (
	KnotParserRULE_start            = 0
	KnotParserRULE_name             = 1
	KnotParserRULE_definition       = 2
	KnotParserRULE_innerContent     = 3
	KnotParserRULE_value            = 4
	KnotParserRULE_valueOptions     = 5
	KnotParserRULE_typeRule         = 6
	KnotParserRULE_typeOptions      = 7
	KnotParserRULE_voltage          = 8
	KnotParserRULE_temperature      = 9
	KnotParserRULE_voltagesUnits    = 10
	KnotParserRULE_temperatureUnits = 11
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

func (s *StartContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

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
	{
		p.SetState(24)
		p.Name()
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserSENSOR || _la == KnotParserACTUATOR {
		{
			p.SetState(25)
			p.Definition()
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(KnotParserEOF)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) END_CHAR() antlr.TerminalNode {
	return s.GetToken(KnotParserEND_CHAR, 0)
}

func (s *NameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(KnotParserIDENTIFIER)
}

func (s *NameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(KnotParserIDENTIFIER, i)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KnotParserRULE_name)
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
		p.SetState(32)
		p.Match(KnotParserT__0)
	}
	{
		p.SetState(33)
		p.Match(KnotParserT__1)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserIDENTIFIER {
		{
			p.SetState(34)
			p.Match(KnotParserIDENTIFIER)
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(KnotParserT__1)
	}
	{
		p.SetState(40)
		p.Match(KnotParserEND_CHAR)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
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

func (s *DefinitionContext) GetOp() antlr.Token { return s.op }

func (s *DefinitionContext) SetOp(v antlr.Token) { s.op = v }

func (s *DefinitionContext) SENSOR() antlr.TerminalNode {
	return s.GetToken(KnotParserSENSOR, 0)
}

func (s *DefinitionContext) ACTUATOR() antlr.TerminalNode {
	return s.GetToken(KnotParserACTUATOR, 0)
}

func (s *DefinitionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(KnotParserIDENTIFIER)
}

func (s *DefinitionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(KnotParserIDENTIFIER, i)
}

func (s *DefinitionContext) AllInnerContent() []IInnerContentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInnerContentContext)(nil)).Elem())
	var tst = make([]IInnerContentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInnerContentContext)
		}
	}

	return tst
}

func (s *DefinitionContext) InnerContent(i int) IInnerContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInnerContentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInnerContentContext)
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
	p.EnterRule(localctx, 4, KnotParserRULE_definition)
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
		p.SetState(42)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*DefinitionContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == KnotParserSENSOR || _la == KnotParserACTUATOR) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*DefinitionContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(43)
		p.Match(KnotParserT__1)
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserIDENTIFIER {
		{
			p.SetState(44)
			p.Match(KnotParserIDENTIFIER)
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(KnotParserT__1)
	}
	{
		p.SetState(50)
		p.Match(KnotParserT__2)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserVALUE {
		{
			p.SetState(51)
			p.InnerContent()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(KnotParserT__3)
	}

	return localctx
}

// IInnerContentContext is an interface to support dynamic dispatch.
type IInnerContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInnerContentContext differentiates from other interfaces.
	IsInnerContentContext()
}

type InnerContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerContentContext() *InnerContentContext {
	var p = new(InnerContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_innerContent
	return p
}

func (*InnerContentContext) IsInnerContentContext() {}

func NewInnerContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerContentContext {
	var p = new(InnerContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_innerContent

	return p
}

func (s *InnerContentContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerContentContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *InnerContentContext) TypeRule() ITypeRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRuleContext)
}

func (s *InnerContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterInnerContent(s)
	}
}

func (s *InnerContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitInnerContent(s)
	}
}

func (s *InnerContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitInnerContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) InnerContent() (localctx IInnerContentContext) {
	localctx = NewInnerContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KnotParserRULE_innerContent)

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
		p.SetState(58)
		p.Value()
	}
	{
		p.SetState(59)
		p.TypeRule()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(KnotParserVALUE, 0)
}

func (s *ValueContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(KnotParserSEPARATOR, 0)
}

func (s *ValueContext) ValueOptions() IValueOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueOptionsContext)
}

func (s *ValueContext) END_CHAR() antlr.TerminalNode {
	return s.GetToken(KnotParserEND_CHAR, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KnotParserRULE_value)

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
		p.SetState(61)
		p.Match(KnotParserVALUE)
	}
	{
		p.SetState(62)
		p.Match(KnotParserSEPARATOR)
	}
	{
		p.SetState(63)
		p.ValueOptions()
	}
	{
		p.SetState(64)
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
	p.EnterRule(localctx, 10, KnotParserRULE_valueOptions)
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
		p.SetState(66)

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

// ITypeRuleContext is an interface to support dynamic dispatch.
type ITypeRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op rule contexts.
	GetOp() ITypeOptionsContext

	// SetOp sets the op rule contexts.
	SetOp(ITypeOptionsContext)

	// IsTypeRuleContext differentiates from other interfaces.
	IsTypeRuleContext()
}

type TypeRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     ITypeOptionsContext
}

func NewEmptyTypeRuleContext() *TypeRuleContext {
	var p = new(TypeRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_typeRule
	return p
}

func (*TypeRuleContext) IsTypeRuleContext() {}

func NewTypeRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRuleContext {
	var p = new(TypeRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_typeRule

	return p
}

func (s *TypeRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRuleContext) GetOp() ITypeOptionsContext { return s.op }

func (s *TypeRuleContext) SetOp(v ITypeOptionsContext) { s.op = v }

func (s *TypeRuleContext) TYPE() antlr.TerminalNode {
	return s.GetToken(KnotParserTYPE, 0)
}

func (s *TypeRuleContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(KnotParserSEPARATOR, 0)
}

func (s *TypeRuleContext) END_CHAR() antlr.TerminalNode {
	return s.GetToken(KnotParserEND_CHAR, 0)
}

func (s *TypeRuleContext) TypeOptions() ITypeOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeOptionsContext)
}

func (s *TypeRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterTypeRule(s)
	}
}

func (s *TypeRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitTypeRule(s)
	}
}

func (s *TypeRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitTypeRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) TypeRule() (localctx ITypeRuleContext) {
	localctx = NewTypeRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KnotParserRULE_typeRule)

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
		p.SetState(68)
		p.Match(KnotParserTYPE)
	}
	{
		p.SetState(69)
		p.Match(KnotParserSEPARATOR)
	}
	{
		p.SetState(70)

		var _x = p.TypeOptions()

		localctx.(*TypeRuleContext).op = _x
	}
	{
		p.SetState(71)
		p.Match(KnotParserEND_CHAR)
	}

	return localctx
}

// ITypeOptionsContext is an interface to support dynamic dispatch.
type ITypeOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeOptionsContext differentiates from other interfaces.
	IsTypeOptionsContext()
}

type TypeOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeOptionsContext() *TypeOptionsContext {
	var p = new(TypeOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_typeOptions
	return p
}

func (*TypeOptionsContext) IsTypeOptionsContext() {}

func NewTypeOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeOptionsContext {
	var p = new(TypeOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_typeOptions

	return p
}

func (s *TypeOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeOptionsContext) Voltage() IVoltageContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVoltageContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVoltageContext)
}

func (s *TypeOptionsContext) CURRENT() antlr.TerminalNode {
	return s.GetToken(KnotParserCURRENT, 0)
}

func (s *TypeOptionsContext) RESISTANCE() antlr.TerminalNode {
	return s.GetToken(KnotParserRESISTANCE, 0)
}

func (s *TypeOptionsContext) POWER() antlr.TerminalNode {
	return s.GetToken(KnotParserPOWER, 0)
}

func (s *TypeOptionsContext) Temperature() ITemperatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemperatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemperatureContext)
}

func (s *TypeOptionsContext) RELATIVEHUMIDITY() antlr.TerminalNode {
	return s.GetToken(KnotParserRELATIVEHUMIDITY, 0)
}

func (s *TypeOptionsContext) LUMINOSITY() antlr.TerminalNode {
	return s.GetToken(KnotParserLUMINOSITY, 0)
}

func (s *TypeOptionsContext) TIME() antlr.TerminalNode {
	return s.GetToken(KnotParserTIME, 0)
}

func (s *TypeOptionsContext) MASS() antlr.TerminalNode {
	return s.GetToken(KnotParserMASS, 0)
}

func (s *TypeOptionsContext) PRESSURE() antlr.TerminalNode {
	return s.GetToken(KnotParserPRESSURE, 0)
}

func (s *TypeOptionsContext) DISTANCE() antlr.TerminalNode {
	return s.GetToken(KnotParserDISTANCE, 0)
}

func (s *TypeOptionsContext) ANGLE() antlr.TerminalNode {
	return s.GetToken(KnotParserANGLE, 0)
}

func (s *TypeOptionsContext) VOLUME() antlr.TerminalNode {
	return s.GetToken(KnotParserVOLUME, 0)
}

func (s *TypeOptionsContext) AREA() antlr.TerminalNode {
	return s.GetToken(KnotParserAREA, 0)
}

func (s *TypeOptionsContext) RAIN() antlr.TerminalNode {
	return s.GetToken(KnotParserRAIN, 0)
}

func (s *TypeOptionsContext) DENSITY() antlr.TerminalNode {
	return s.GetToken(KnotParserDENSITY, 0)
}

func (s *TypeOptionsContext) LATITUDE() antlr.TerminalNode {
	return s.GetToken(KnotParserLATITUDE, 0)
}

func (s *TypeOptionsContext) LONGITUDE() antlr.TerminalNode {
	return s.GetToken(KnotParserLONGITUDE, 0)
}

func (s *TypeOptionsContext) SPEED() antlr.TerminalNode {
	return s.GetToken(KnotParserSPEED, 0)
}

func (s *TypeOptionsContext) VOLUMEFLOW() antlr.TerminalNode {
	return s.GetToken(KnotParserVOLUMEFLOW, 0)
}

func (s *TypeOptionsContext) ENERGY() antlr.TerminalNode {
	return s.GetToken(KnotParserENERGY, 0)
}

func (s *TypeOptionsContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(KnotParserSWITCH, 0)
}

func (s *TypeOptionsContext) PRESENCE() antlr.TerminalNode {
	return s.GetToken(KnotParserPRESENCE, 0)
}

func (s *TypeOptionsContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(KnotParserCOMMAND, 0)
}

func (s *TypeOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterTypeOptions(s)
	}
}

func (s *TypeOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitTypeOptions(s)
	}
}

func (s *TypeOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitTypeOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) TypeOptions() (localctx ITypeOptionsContext) {
	localctx = NewTypeOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KnotParserRULE_typeOptions)

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
	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KnotParserT__4:
		{
			p.SetState(73)
			p.Voltage()
		}

	case KnotParserCURRENT:
		{
			p.SetState(74)
			p.Match(KnotParserCURRENT)
		}

	case KnotParserRESISTANCE:
		{
			p.SetState(75)
			p.Match(KnotParserRESISTANCE)
		}

	case KnotParserPOWER:
		{
			p.SetState(76)
			p.Match(KnotParserPOWER)
		}

	case KnotParserT__5:
		{
			p.SetState(77)
			p.Temperature()
		}

	case KnotParserRELATIVEHUMIDITY:
		{
			p.SetState(78)
			p.Match(KnotParserRELATIVEHUMIDITY)
		}

	case KnotParserLUMINOSITY:
		{
			p.SetState(79)
			p.Match(KnotParserLUMINOSITY)
		}

	case KnotParserTIME:
		{
			p.SetState(80)
			p.Match(KnotParserTIME)
		}

	case KnotParserMASS:
		{
			p.SetState(81)
			p.Match(KnotParserMASS)
		}

	case KnotParserPRESSURE:
		{
			p.SetState(82)
			p.Match(KnotParserPRESSURE)
		}

	case KnotParserDISTANCE:
		{
			p.SetState(83)
			p.Match(KnotParserDISTANCE)
		}

	case KnotParserANGLE:
		{
			p.SetState(84)
			p.Match(KnotParserANGLE)
		}

	case KnotParserVOLUME:
		{
			p.SetState(85)
			p.Match(KnotParserVOLUME)
		}

	case KnotParserAREA:
		{
			p.SetState(86)
			p.Match(KnotParserAREA)
		}

	case KnotParserRAIN:
		{
			p.SetState(87)
			p.Match(KnotParserRAIN)
		}

	case KnotParserDENSITY:
		{
			p.SetState(88)
			p.Match(KnotParserDENSITY)
		}

	case KnotParserLATITUDE:
		{
			p.SetState(89)
			p.Match(KnotParserLATITUDE)
		}

	case KnotParserLONGITUDE:
		{
			p.SetState(90)
			p.Match(KnotParserLONGITUDE)
		}

	case KnotParserSPEED:
		{
			p.SetState(91)
			p.Match(KnotParserSPEED)
		}

	case KnotParserVOLUMEFLOW:
		{
			p.SetState(92)
			p.Match(KnotParserVOLUMEFLOW)
		}

	case KnotParserENERGY:
		{
			p.SetState(93)
			p.Match(KnotParserENERGY)
		}

	case KnotParserSWITCH:
		{
			p.SetState(94)
			p.Match(KnotParserSWITCH)
		}

	case KnotParserPRESENCE:
		{
			p.SetState(95)
			p.Match(KnotParserPRESENCE)
		}

	case KnotParserCOMMAND:
		{
			p.SetState(96)
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
	p.EnterRule(localctx, 16, KnotParserRULE_voltage)

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
		p.SetState(99)
		p.Match(KnotParserT__4)
	}
	{
		p.SetState(100)
		p.VoltagesUnits()
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
	p.EnterRule(localctx, 18, KnotParserRULE_temperature)

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
		p.SetState(102)
		p.Match(KnotParserT__5)
	}
	{
		p.SetState(103)
		p.TemperatureUnits()
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
	p.EnterRule(localctx, 20, KnotParserRULE_voltagesUnits)
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
		p.SetState(105)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VoltagesUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserT__6)|(1<<KnotParserT__7)|(1<<KnotParserT__8))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VoltagesUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 22, KnotParserRULE_temperatureUnits)
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
		p.SetState(107)

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
