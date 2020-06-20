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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 100, 246,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 6, 2, 94, 10, 2, 13, 2, 14, 2, 95, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 104, 10, 3, 13, 3, 14, 3, 105, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 144, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 2, 2, 47, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
	80, 82, 84, 86, 88, 90, 2, 19, 3, 2, 85, 86, 3, 2, 91, 94, 3, 2, 8, 10,
	3, 2, 12, 13, 3, 2, 17, 19, 3, 2, 21, 23, 3, 2, 25, 27, 3, 2, 29, 33, 3,
	2, 35, 38, 3, 2, 40, 42, 3, 2, 44, 47, 3, 2, 49, 50, 3, 2, 52, 55, 3, 2,
	57, 60, 4, 2, 30, 30, 68, 70, 4, 2, 25, 25, 72, 76, 3, 2, 78, 83, 2, 225,
	2, 93, 3, 2, 2, 2, 4, 99, 3, 2, 2, 2, 6, 109, 3, 2, 2, 2, 8, 117, 3, 2,
	2, 2, 10, 143, 3, 2, 2, 2, 12, 145, 3, 2, 2, 2, 14, 148, 3, 2, 2, 2, 16,
	150, 3, 2, 2, 2, 18, 153, 3, 2, 2, 2, 20, 155, 3, 2, 2, 2, 22, 158, 3,
	2, 2, 2, 24, 160, 3, 2, 2, 2, 26, 163, 3, 2, 2, 2, 28, 165, 3, 2, 2, 2,
	30, 168, 3, 2, 2, 2, 32, 170, 3, 2, 2, 2, 34, 173, 3, 2, 2, 2, 36, 175,
	3, 2, 2, 2, 38, 178, 3, 2, 2, 2, 40, 180, 3, 2, 2, 2, 42, 183, 3, 2, 2,
	2, 44, 185, 3, 2, 2, 2, 46, 188, 3, 2, 2, 2, 48, 190, 3, 2, 2, 2, 50, 193,
	3, 2, 2, 2, 52, 195, 3, 2, 2, 2, 54, 198, 3, 2, 2, 2, 56, 200, 3, 2, 2,
	2, 58, 203, 3, 2, 2, 2, 60, 205, 3, 2, 2, 2, 62, 208, 3, 2, 2, 2, 64, 210,
	3, 2, 2, 2, 66, 213, 3, 2, 2, 2, 68, 215, 3, 2, 2, 2, 70, 218, 3, 2, 2,
	2, 72, 220, 3, 2, 2, 2, 74, 223, 3, 2, 2, 2, 76, 225, 3, 2, 2, 2, 78, 228,
	3, 2, 2, 2, 80, 230, 3, 2, 2, 2, 82, 233, 3, 2, 2, 2, 84, 235, 3, 2, 2,
	2, 86, 238, 3, 2, 2, 2, 88, 240, 3, 2, 2, 2, 90, 243, 3, 2, 2, 2, 92, 94,
	5, 4, 3, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2,
	95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 7, 2, 2, 3, 98, 3, 3, 2,
	2, 2, 99, 100, 7, 84, 2, 2, 100, 101, 7, 99, 2, 2, 101, 103, 7, 3, 2, 2,
	102, 104, 5, 6, 4, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105,
	103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108,
	7, 4, 2, 2, 108, 5, 3, 2, 2, 2, 109, 110, 9, 2, 2, 2, 110, 111, 5, 8, 5,
	2, 111, 112, 7, 99, 2, 2, 112, 113, 7, 5, 2, 2, 113, 114, 5, 10, 6, 2,
	114, 115, 7, 6, 2, 2, 115, 116, 7, 87, 2, 2, 116, 7, 3, 2, 2, 2, 117, 118,
	9, 3, 2, 2, 118, 9, 3, 2, 2, 2, 119, 144, 5, 12, 7, 2, 120, 144, 5, 16,
	9, 2, 121, 144, 5, 20, 11, 2, 122, 144, 5, 24, 13, 2, 123, 144, 5, 28,
	15, 2, 124, 144, 5, 32, 17, 2, 125, 144, 5, 36, 19, 2, 126, 144, 5, 40,
	21, 2, 127, 144, 5, 44, 23, 2, 128, 144, 5, 48, 25, 2, 129, 144, 5, 52,
	27, 2, 130, 144, 5, 56, 29, 2, 131, 144, 5, 60, 31, 2, 132, 144, 5, 64,
	33, 2, 133, 144, 5, 68, 35, 2, 134, 144, 5, 72, 37, 2, 135, 144, 5, 76,
	39, 2, 136, 144, 5, 80, 41, 2, 137, 144, 5, 84, 43, 2, 138, 144, 5, 88,
	45, 2, 139, 144, 7, 95, 2, 2, 140, 144, 7, 96, 2, 2, 141, 144, 7, 97, 2,
	2, 142, 144, 7, 98, 2, 2, 143, 119, 3, 2, 2, 2, 143, 120, 3, 2, 2, 2, 143,
	121, 3, 2, 2, 2, 143, 122, 3, 2, 2, 2, 143, 123, 3, 2, 2, 2, 143, 124,
	3, 2, 2, 2, 143, 125, 3, 2, 2, 2, 143, 126, 3, 2, 2, 2, 143, 127, 3, 2,
	2, 2, 143, 128, 3, 2, 2, 2, 143, 129, 3, 2, 2, 2, 143, 130, 3, 2, 2, 2,
	143, 131, 3, 2, 2, 2, 143, 132, 3, 2, 2, 2, 143, 133, 3, 2, 2, 2, 143,
	134, 3, 2, 2, 2, 143, 135, 3, 2, 2, 2, 143, 136, 3, 2, 2, 2, 143, 137,
	3, 2, 2, 2, 143, 138, 3, 2, 2, 2, 143, 139, 3, 2, 2, 2, 143, 140, 3, 2,
	2, 2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2, 2, 2, 144, 11, 3, 2, 2, 2,
	145, 146, 7, 7, 2, 2, 146, 147, 5, 14, 8, 2, 147, 13, 3, 2, 2, 2, 148,
	149, 9, 4, 2, 2, 149, 15, 3, 2, 2, 2, 150, 151, 7, 11, 2, 2, 151, 152,
	5, 18, 10, 2, 152, 17, 3, 2, 2, 2, 153, 154, 9, 5, 2, 2, 154, 19, 3, 2,
	2, 2, 155, 156, 7, 14, 2, 2, 156, 157, 5, 22, 12, 2, 157, 21, 3, 2, 2,
	2, 158, 159, 7, 15, 2, 2, 159, 23, 3, 2, 2, 2, 160, 161, 7, 16, 2, 2, 161,
	162, 5, 26, 14, 2, 162, 25, 3, 2, 2, 2, 163, 164, 9, 6, 2, 2, 164, 27,
	3, 2, 2, 2, 165, 166, 7, 20, 2, 2, 166, 167, 5, 30, 16, 2, 167, 29, 3,
	2, 2, 2, 168, 169, 9, 7, 2, 2, 169, 31, 3, 2, 2, 2, 170, 171, 7, 24, 2,
	2, 171, 172, 5, 34, 18, 2, 172, 33, 3, 2, 2, 2, 173, 174, 9, 8, 2, 2, 174,
	35, 3, 2, 2, 2, 175, 176, 7, 28, 2, 2, 176, 177, 5, 38, 20, 2, 177, 37,
	3, 2, 2, 2, 178, 179, 9, 9, 2, 2, 179, 39, 3, 2, 2, 2, 180, 181, 7, 34,
	2, 2, 181, 182, 5, 42, 22, 2, 182, 41, 3, 2, 2, 2, 183, 184, 9, 10, 2,
	2, 184, 43, 3, 2, 2, 2, 185, 186, 7, 39, 2, 2, 186, 187, 5, 46, 24, 2,
	187, 45, 3, 2, 2, 2, 188, 189, 9, 11, 2, 2, 189, 47, 3, 2, 2, 2, 190, 191,
	7, 43, 2, 2, 191, 192, 5, 50, 26, 2, 192, 49, 3, 2, 2, 2, 193, 194, 9,
	12, 2, 2, 194, 51, 3, 2, 2, 2, 195, 196, 7, 48, 2, 2, 196, 197, 5, 54,
	28, 2, 197, 53, 3, 2, 2, 2, 198, 199, 9, 13, 2, 2, 199, 55, 3, 2, 2, 2,
	200, 201, 7, 51, 2, 2, 201, 202, 5, 58, 30, 2, 202, 57, 3, 2, 2, 2, 203,
	204, 9, 14, 2, 2, 204, 59, 3, 2, 2, 2, 205, 206, 7, 56, 2, 2, 206, 207,
	5, 62, 32, 2, 207, 61, 3, 2, 2, 2, 208, 209, 9, 15, 2, 2, 209, 63, 3, 2,
	2, 2, 210, 211, 7, 61, 2, 2, 211, 212, 5, 66, 34, 2, 212, 65, 3, 2, 2,
	2, 213, 214, 7, 60, 2, 2, 214, 67, 3, 2, 2, 2, 215, 216, 7, 62, 2, 2, 216,
	217, 5, 70, 36, 2, 217, 69, 3, 2, 2, 2, 218, 219, 7, 63, 2, 2, 219, 71,
	3, 2, 2, 2, 220, 221, 7, 64, 2, 2, 221, 222, 5, 74, 38, 2, 222, 73, 3,
	2, 2, 2, 223, 224, 7, 65, 2, 2, 224, 75, 3, 2, 2, 2, 225, 226, 7, 66, 2,
	2, 226, 227, 5, 78, 40, 2, 227, 77, 3, 2, 2, 2, 228, 229, 7, 65, 2, 2,
	229, 79, 3, 2, 2, 2, 230, 231, 7, 67, 2, 2, 231, 232, 5, 82, 42, 2, 232,
	81, 3, 2, 2, 2, 233, 234, 9, 16, 2, 2, 234, 83, 3, 2, 2, 2, 235, 236, 7,
	71, 2, 2, 236, 237, 5, 86, 44, 2, 237, 85, 3, 2, 2, 2, 238, 239, 9, 17,
	2, 2, 239, 87, 3, 2, 2, 2, 240, 241, 7, 77, 2, 2, 241, 242, 5, 90, 46,
	2, 242, 89, 3, 2, 2, 2, 243, 244, 9, 18, 2, 2, 244, 91, 3, 2, 2, 2, 5,
	95, 105, 143,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'('", "')'", "'voltage in '", "'V'", "'mV'", "'kV'",
	"'current in '", "'A'", "'mA'", "'resistance in '", "'ohm'", "'power in '",
	"'W'", "'kW'", "'MW'", "'temperature in '", "'F'", "'C'", "'K'", "'luminosity in '",
	"'lm'", "'cd'", "'lx'", "'time in '", "'s'", "'ms'", "'us'", "'min'", "'h'",
	"'mass in '", "'g'", "'kg'", "'lb'", "'oz'", "'pressure in '", "'Pa'",
	"'psi'", "'bar'", "'distance in '", "'m'", "'cm'", "'mi'", "'in'", "'angle in '",
	"'deg'", "'rad'", "'volume in '", "'l'", "'gal'", "'ml'", "'floz'", "'area in '",
	"'m2'", "'ha'", "'ac'", "'mm'", "'rain in '", "'density in '", "'kgm3'",
	"'latitude in '", "'degree'", "'longitude in '", "'speed in '", "'cms'",
	"'kms'", "'mih'", "'volumeflow in '", "'m3s'", "'scmm'", "'ls'", "'ft3s'",
	"'galm'", "'energy in '", "'J'", "'nm'", "'Wh'", "'KWh'", "'cal'", "'Kcal'",
	"'thing'", "'sensor'", "'actuator'", "';'", "", "", "", "'bool'", "'int'",
	"'float'", "'bytes'", "'relativehumidity'", "'switch'", "'presence'", "'command'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "THING", "SENSOR", "ACTUATOR",
	"END_CHAR", "WS", "COMMENT", "LINE_COMMENT", "BOOL", "INT", "FLOAT", "BYTES",
	"RELATIVEHUMIDITY", "SWITCH", "PRESENCE", "COMMAND", "IDENTIFIER", "ID",
}

var ruleNames = []string{
	"start", "definition", "thingContent", "valueOptions", "unitTypeOptions",
	"voltage", "voltagesUnits", "current", "currentUnits", "resistance", "resistanceUnits",
	"power", "powerUnits", "temperature", "temperatureUnits", "luminosity",
	"luminosityUnits", "time", "timeUnits", "mass", "massUnits", "pressure",
	"pressureUnits", "distance", "distanceUnits", "angle", "angleUnits", "volume",
	"volumeUnits", "area", "areaUnits", "rain", "rainUnits", "density", "densityUnits",
	"latitude", "latitudeUnits", "longitude", "longitudeUnits", "speed", "speedUnits",
	"volumeflow", "volumeflowUnits", "energy", "energyUnits",
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
	KnotParserT__12            = 13
	KnotParserT__13            = 14
	KnotParserT__14            = 15
	KnotParserT__15            = 16
	KnotParserT__16            = 17
	KnotParserT__17            = 18
	KnotParserT__18            = 19
	KnotParserT__19            = 20
	KnotParserT__20            = 21
	KnotParserT__21            = 22
	KnotParserT__22            = 23
	KnotParserT__23            = 24
	KnotParserT__24            = 25
	KnotParserT__25            = 26
	KnotParserT__26            = 27
	KnotParserT__27            = 28
	KnotParserT__28            = 29
	KnotParserT__29            = 30
	KnotParserT__30            = 31
	KnotParserT__31            = 32
	KnotParserT__32            = 33
	KnotParserT__33            = 34
	KnotParserT__34            = 35
	KnotParserT__35            = 36
	KnotParserT__36            = 37
	KnotParserT__37            = 38
	KnotParserT__38            = 39
	KnotParserT__39            = 40
	KnotParserT__40            = 41
	KnotParserT__41            = 42
	KnotParserT__42            = 43
	KnotParserT__43            = 44
	KnotParserT__44            = 45
	KnotParserT__45            = 46
	KnotParserT__46            = 47
	KnotParserT__47            = 48
	KnotParserT__48            = 49
	KnotParserT__49            = 50
	KnotParserT__50            = 51
	KnotParserT__51            = 52
	KnotParserT__52            = 53
	KnotParserT__53            = 54
	KnotParserT__54            = 55
	KnotParserT__55            = 56
	KnotParserT__56            = 57
	KnotParserT__57            = 58
	KnotParserT__58            = 59
	KnotParserT__59            = 60
	KnotParserT__60            = 61
	KnotParserT__61            = 62
	KnotParserT__62            = 63
	KnotParserT__63            = 64
	KnotParserT__64            = 65
	KnotParserT__65            = 66
	KnotParserT__66            = 67
	KnotParserT__67            = 68
	KnotParserT__68            = 69
	KnotParserT__69            = 70
	KnotParserT__70            = 71
	KnotParserT__71            = 72
	KnotParserT__72            = 73
	KnotParserT__73            = 74
	KnotParserT__74            = 75
	KnotParserT__75            = 76
	KnotParserT__76            = 77
	KnotParserT__77            = 78
	KnotParserT__78            = 79
	KnotParserT__79            = 80
	KnotParserT__80            = 81
	KnotParserTHING            = 82
	KnotParserSENSOR           = 83
	KnotParserACTUATOR         = 84
	KnotParserEND_CHAR         = 85
	KnotParserWS               = 86
	KnotParserCOMMENT          = 87
	KnotParserLINE_COMMENT     = 88
	KnotParserBOOL             = 89
	KnotParserINT              = 90
	KnotParserFLOAT            = 91
	KnotParserBYTES            = 92
	KnotParserRELATIVEHUMIDITY = 93
	KnotParserSWITCH           = 94
	KnotParserPRESENCE         = 95
	KnotParserCOMMAND          = 96
	KnotParserIDENTIFIER       = 97
	KnotParserID               = 98
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
	KnotParserRULE_current          = 7
	KnotParserRULE_currentUnits     = 8
	KnotParserRULE_resistance       = 9
	KnotParserRULE_resistanceUnits  = 10
	KnotParserRULE_power            = 11
	KnotParserRULE_powerUnits       = 12
	KnotParserRULE_temperature      = 13
	KnotParserRULE_temperatureUnits = 14
	KnotParserRULE_luminosity       = 15
	KnotParserRULE_luminosityUnits  = 16
	KnotParserRULE_time             = 17
	KnotParserRULE_timeUnits        = 18
	KnotParserRULE_mass             = 19
	KnotParserRULE_massUnits        = 20
	KnotParserRULE_pressure         = 21
	KnotParserRULE_pressureUnits    = 22
	KnotParserRULE_distance         = 23
	KnotParserRULE_distanceUnits    = 24
	KnotParserRULE_angle            = 25
	KnotParserRULE_angleUnits       = 26
	KnotParserRULE_volume           = 27
	KnotParserRULE_volumeUnits      = 28
	KnotParserRULE_area             = 29
	KnotParserRULE_areaUnits        = 30
	KnotParserRULE_rain             = 31
	KnotParserRULE_rainUnits        = 32
	KnotParserRULE_density          = 33
	KnotParserRULE_densityUnits     = 34
	KnotParserRULE_latitude         = 35
	KnotParserRULE_latitudeUnits    = 36
	KnotParserRULE_longitude        = 37
	KnotParserRULE_longitudeUnits   = 38
	KnotParserRULE_speed            = 39
	KnotParserRULE_speedUnits       = 40
	KnotParserRULE_volumeflow       = 41
	KnotParserRULE_volumeflowUnits  = 42
	KnotParserRULE_energy           = 43
	KnotParserRULE_energyUnits      = 44
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserTHING {
		{
			p.SetState(90)
			p.Definition()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
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
		p.SetState(97)
		p.Match(KnotParserTHING)
	}
	{
		p.SetState(98)
		p.Match(KnotParserIDENTIFIER)
	}
	{
		p.SetState(99)
		p.Match(KnotParserT__0)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == KnotParserSENSOR || _la == KnotParserACTUATOR {
		{
			p.SetState(100)
			p.ThingContent()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
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
		p.SetState(107)

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
		p.SetState(108)
		p.ValueOptions()
	}
	{
		p.SetState(109)
		p.Match(KnotParserIDENTIFIER)
	}
	{
		p.SetState(110)
		p.Match(KnotParserT__2)
	}
	{
		p.SetState(111)
		p.UnitTypeOptions()
	}
	{
		p.SetState(112)
		p.Match(KnotParserT__3)
	}
	{
		p.SetState(113)
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
		p.SetState(115)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ValueOptionsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-89)&-(0x1f+1)) == 0 && ((1<<uint((_la-89)))&((1<<(KnotParserBOOL-89))|(1<<(KnotParserINT-89))|(1<<(KnotParserFLOAT-89))|(1<<(KnotParserBYTES-89)))) != 0) {
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

func (s *UnitTypeOptionsContext) Current() ICurrentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurrentContext)
}

func (s *UnitTypeOptionsContext) Resistance() IResistanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResistanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResistanceContext)
}

func (s *UnitTypeOptionsContext) Power() IPowerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowerContext)
}

func (s *UnitTypeOptionsContext) Temperature() ITemperatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemperatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemperatureContext)
}

func (s *UnitTypeOptionsContext) Luminosity() ILuminosityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILuminosityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILuminosityContext)
}

func (s *UnitTypeOptionsContext) Time() ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *UnitTypeOptionsContext) Mass() IMassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMassContext)
}

func (s *UnitTypeOptionsContext) Pressure() IPressureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPressureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPressureContext)
}

func (s *UnitTypeOptionsContext) Distance() IDistanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDistanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDistanceContext)
}

func (s *UnitTypeOptionsContext) Angle() IAngleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAngleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAngleContext)
}

func (s *UnitTypeOptionsContext) Volume() IVolumeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVolumeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVolumeContext)
}

func (s *UnitTypeOptionsContext) Area() IAreaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAreaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAreaContext)
}

func (s *UnitTypeOptionsContext) Rain() IRainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRainContext)
}

func (s *UnitTypeOptionsContext) Density() IDensityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDensityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDensityContext)
}

func (s *UnitTypeOptionsContext) Latitude() ILatitudeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILatitudeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILatitudeContext)
}

func (s *UnitTypeOptionsContext) Longitude() ILongitudeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILongitudeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILongitudeContext)
}

func (s *UnitTypeOptionsContext) Speed() ISpeedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpeedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpeedContext)
}

func (s *UnitTypeOptionsContext) Volumeflow() IVolumeflowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVolumeflowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVolumeflowContext)
}

func (s *UnitTypeOptionsContext) Energy() IEnergyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnergyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnergyContext)
}

func (s *UnitTypeOptionsContext) RELATIVEHUMIDITY() antlr.TerminalNode {
	return s.GetToken(KnotParserRELATIVEHUMIDITY, 0)
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KnotParserT__4:
		{
			p.SetState(117)
			p.Voltage()
		}

	case KnotParserT__8:
		{
			p.SetState(118)
			p.Current()
		}

	case KnotParserT__11:
		{
			p.SetState(119)
			p.Resistance()
		}

	case KnotParserT__13:
		{
			p.SetState(120)
			p.Power()
		}

	case KnotParserT__17:
		{
			p.SetState(121)
			p.Temperature()
		}

	case KnotParserT__21:
		{
			p.SetState(122)
			p.Luminosity()
		}

	case KnotParserT__25:
		{
			p.SetState(123)
			p.Time()
		}

	case KnotParserT__31:
		{
			p.SetState(124)
			p.Mass()
		}

	case KnotParserT__36:
		{
			p.SetState(125)
			p.Pressure()
		}

	case KnotParserT__40:
		{
			p.SetState(126)
			p.Distance()
		}

	case KnotParserT__45:
		{
			p.SetState(127)
			p.Angle()
		}

	case KnotParserT__48:
		{
			p.SetState(128)
			p.Volume()
		}

	case KnotParserT__53:
		{
			p.SetState(129)
			p.Area()
		}

	case KnotParserT__58:
		{
			p.SetState(130)
			p.Rain()
		}

	case KnotParserT__59:
		{
			p.SetState(131)
			p.Density()
		}

	case KnotParserT__61:
		{
			p.SetState(132)
			p.Latitude()
		}

	case KnotParserT__63:
		{
			p.SetState(133)
			p.Longitude()
		}

	case KnotParserT__64:
		{
			p.SetState(134)
			p.Speed()
		}

	case KnotParserT__68:
		{
			p.SetState(135)
			p.Volumeflow()
		}

	case KnotParserT__74:
		{
			p.SetState(136)
			p.Energy()
		}

	case KnotParserRELATIVEHUMIDITY:
		{
			p.SetState(137)
			p.Match(KnotParserRELATIVEHUMIDITY)
		}

	case KnotParserSWITCH:
		{
			p.SetState(138)
			p.Match(KnotParserSWITCH)
		}

	case KnotParserPRESENCE:
		{
			p.SetState(139)
			p.Match(KnotParserPRESENCE)
		}

	case KnotParserCOMMAND:
		{
			p.SetState(140)
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
		p.SetState(143)
		p.Match(KnotParserT__4)
	}
	{
		p.SetState(144)
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
		p.SetState(146)

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

// ICurrentContext is an interface to support dynamic dispatch.
type ICurrentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrentContext differentiates from other interfaces.
	IsCurrentContext()
}

type CurrentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrentContext() *CurrentContext {
	var p = new(CurrentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_current
	return p
}

func (*CurrentContext) IsCurrentContext() {}

func NewCurrentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrentContext {
	var p = new(CurrentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_current

	return p
}

func (s *CurrentContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrentContext) CurrentUnits() ICurrentUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrentUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurrentUnitsContext)
}

func (s *CurrentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterCurrent(s)
	}
}

func (s *CurrentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitCurrent(s)
	}
}

func (s *CurrentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitCurrent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Current() (localctx ICurrentContext) {
	localctx = NewCurrentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KnotParserRULE_current)

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
		p.SetState(148)
		p.Match(KnotParserT__8)
	}
	{
		p.SetState(149)
		p.CurrentUnits()
	}

	return localctx
}

// ICurrentUnitsContext is an interface to support dynamic dispatch.
type ICurrentUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsCurrentUnitsContext differentiates from other interfaces.
	IsCurrentUnitsContext()
}

type CurrentUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyCurrentUnitsContext() *CurrentUnitsContext {
	var p = new(CurrentUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_currentUnits
	return p
}

func (*CurrentUnitsContext) IsCurrentUnitsContext() {}

func NewCurrentUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrentUnitsContext {
	var p = new(CurrentUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_currentUnits

	return p
}

func (s *CurrentUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrentUnitsContext) GetOp() antlr.Token { return s.op }

func (s *CurrentUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *CurrentUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrentUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterCurrentUnits(s)
	}
}

func (s *CurrentUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitCurrentUnits(s)
	}
}

func (s *CurrentUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitCurrentUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) CurrentUnits() (localctx ICurrentUnitsContext) {
	localctx = NewCurrentUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KnotParserRULE_currentUnits)
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
		p.SetState(151)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CurrentUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == KnotParserT__9 || _la == KnotParserT__10) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CurrentUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IResistanceContext is an interface to support dynamic dispatch.
type IResistanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResistanceContext differentiates from other interfaces.
	IsResistanceContext()
}

type ResistanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResistanceContext() *ResistanceContext {
	var p = new(ResistanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_resistance
	return p
}

func (*ResistanceContext) IsResistanceContext() {}

func NewResistanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResistanceContext {
	var p = new(ResistanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_resistance

	return p
}

func (s *ResistanceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResistanceContext) ResistanceUnits() IResistanceUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResistanceUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResistanceUnitsContext)
}

func (s *ResistanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResistanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResistanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterResistance(s)
	}
}

func (s *ResistanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitResistance(s)
	}
}

func (s *ResistanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitResistance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Resistance() (localctx IResistanceContext) {
	localctx = NewResistanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KnotParserRULE_resistance)

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
		p.SetState(153)
		p.Match(KnotParserT__11)
	}
	{
		p.SetState(154)
		p.ResistanceUnits()
	}

	return localctx
}

// IResistanceUnitsContext is an interface to support dynamic dispatch.
type IResistanceUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsResistanceUnitsContext differentiates from other interfaces.
	IsResistanceUnitsContext()
}

type ResistanceUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyResistanceUnitsContext() *ResistanceUnitsContext {
	var p = new(ResistanceUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_resistanceUnits
	return p
}

func (*ResistanceUnitsContext) IsResistanceUnitsContext() {}

func NewResistanceUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResistanceUnitsContext {
	var p = new(ResistanceUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_resistanceUnits

	return p
}

func (s *ResistanceUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *ResistanceUnitsContext) GetOp() antlr.Token { return s.op }

func (s *ResistanceUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *ResistanceUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResistanceUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResistanceUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterResistanceUnits(s)
	}
}

func (s *ResistanceUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitResistanceUnits(s)
	}
}

func (s *ResistanceUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitResistanceUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) ResistanceUnits() (localctx IResistanceUnitsContext) {
	localctx = NewResistanceUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KnotParserRULE_resistanceUnits)

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
		p.SetState(156)

		var _m = p.Match(KnotParserT__12)

		localctx.(*ResistanceUnitsContext).op = _m
	}

	return localctx
}

// IPowerContext is an interface to support dynamic dispatch.
type IPowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPowerContext differentiates from other interfaces.
	IsPowerContext()
}

type PowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowerContext() *PowerContext {
	var p = new(PowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_power
	return p
}

func (*PowerContext) IsPowerContext() {}

func NewPowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerContext {
	var p = new(PowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_power

	return p
}

func (s *PowerContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerContext) PowerUnits() IPowerUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowerUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowerUnitsContext)
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Power() (localctx IPowerContext) {
	localctx = NewPowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KnotParserRULE_power)

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
		p.SetState(158)
		p.Match(KnotParserT__13)
	}
	{
		p.SetState(159)
		p.PowerUnits()
	}

	return localctx
}

// IPowerUnitsContext is an interface to support dynamic dispatch.
type IPowerUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsPowerUnitsContext differentiates from other interfaces.
	IsPowerUnitsContext()
}

type PowerUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyPowerUnitsContext() *PowerUnitsContext {
	var p = new(PowerUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_powerUnits
	return p
}

func (*PowerUnitsContext) IsPowerUnitsContext() {}

func NewPowerUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowerUnitsContext {
	var p = new(PowerUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_powerUnits

	return p
}

func (s *PowerUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *PowerUnitsContext) GetOp() antlr.Token { return s.op }

func (s *PowerUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *PowerUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowerUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterPowerUnits(s)
	}
}

func (s *PowerUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitPowerUnits(s)
	}
}

func (s *PowerUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitPowerUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) PowerUnits() (localctx IPowerUnitsContext) {
	localctx = NewPowerUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KnotParserRULE_powerUnits)
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
		p.SetState(161)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PowerUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserT__14)|(1<<KnotParserT__15)|(1<<KnotParserT__16))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PowerUnitsContext).op = _ri
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
	p.EnterRule(localctx, 26, KnotParserRULE_temperature)

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
		p.SetState(163)
		p.Match(KnotParserT__17)
	}
	{
		p.SetState(164)
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
	p.EnterRule(localctx, 28, KnotParserRULE_temperatureUnits)
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
		p.SetState(166)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TemperatureUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserT__18)|(1<<KnotParserT__19)|(1<<KnotParserT__20))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TemperatureUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILuminosityContext is an interface to support dynamic dispatch.
type ILuminosityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLuminosityContext differentiates from other interfaces.
	IsLuminosityContext()
}

type LuminosityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLuminosityContext() *LuminosityContext {
	var p = new(LuminosityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_luminosity
	return p
}

func (*LuminosityContext) IsLuminosityContext() {}

func NewLuminosityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LuminosityContext {
	var p = new(LuminosityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_luminosity

	return p
}

func (s *LuminosityContext) GetParser() antlr.Parser { return s.parser }

func (s *LuminosityContext) LuminosityUnits() ILuminosityUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILuminosityUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILuminosityUnitsContext)
}

func (s *LuminosityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LuminosityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LuminosityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterLuminosity(s)
	}
}

func (s *LuminosityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitLuminosity(s)
	}
}

func (s *LuminosityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitLuminosity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Luminosity() (localctx ILuminosityContext) {
	localctx = NewLuminosityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, KnotParserRULE_luminosity)

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
		p.SetState(168)
		p.Match(KnotParserT__21)
	}
	{
		p.SetState(169)
		p.LuminosityUnits()
	}

	return localctx
}

// ILuminosityUnitsContext is an interface to support dynamic dispatch.
type ILuminosityUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsLuminosityUnitsContext differentiates from other interfaces.
	IsLuminosityUnitsContext()
}

type LuminosityUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyLuminosityUnitsContext() *LuminosityUnitsContext {
	var p = new(LuminosityUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_luminosityUnits
	return p
}

func (*LuminosityUnitsContext) IsLuminosityUnitsContext() {}

func NewLuminosityUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LuminosityUnitsContext {
	var p = new(LuminosityUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_luminosityUnits

	return p
}

func (s *LuminosityUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *LuminosityUnitsContext) GetOp() antlr.Token { return s.op }

func (s *LuminosityUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *LuminosityUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LuminosityUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LuminosityUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterLuminosityUnits(s)
	}
}

func (s *LuminosityUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitLuminosityUnits(s)
	}
}

func (s *LuminosityUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitLuminosityUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) LuminosityUnits() (localctx ILuminosityUnitsContext) {
	localctx = NewLuminosityUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, KnotParserRULE_luminosityUnits)
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
		p.SetState(171)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*LuminosityUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserT__22)|(1<<KnotParserT__23)|(1<<KnotParserT__24))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*LuminosityUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITimeContext is an interface to support dynamic dispatch.
type ITimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeContext differentiates from other interfaces.
	IsTimeContext()
}

type TimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeContext() *TimeContext {
	var p = new(TimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeContext) TimeUnits() ITimeUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeUnitsContext)
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitTime(s)
	}
}

func (s *TimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitTime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Time() (localctx ITimeContext) {
	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, KnotParserRULE_time)

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
		p.SetState(173)
		p.Match(KnotParserT__25)
	}
	{
		p.SetState(174)
		p.TimeUnits()
	}

	return localctx
}

// ITimeUnitsContext is an interface to support dynamic dispatch.
type ITimeUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsTimeUnitsContext differentiates from other interfaces.
	IsTimeUnitsContext()
}

type TimeUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyTimeUnitsContext() *TimeUnitsContext {
	var p = new(TimeUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_timeUnits
	return p
}

func (*TimeUnitsContext) IsTimeUnitsContext() {}

func NewTimeUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeUnitsContext {
	var p = new(TimeUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_timeUnits

	return p
}

func (s *TimeUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeUnitsContext) GetOp() antlr.Token { return s.op }

func (s *TimeUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *TimeUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterTimeUnits(s)
	}
}

func (s *TimeUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitTimeUnits(s)
	}
}

func (s *TimeUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitTimeUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) TimeUnits() (localctx ITimeUnitsContext) {
	localctx = NewTimeUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, KnotParserRULE_timeUnits)
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
		p.SetState(176)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TimeUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KnotParserT__26)|(1<<KnotParserT__27)|(1<<KnotParserT__28)|(1<<KnotParserT__29)|(1<<KnotParserT__30))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TimeUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMassContext is an interface to support dynamic dispatch.
type IMassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMassContext differentiates from other interfaces.
	IsMassContext()
}

type MassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMassContext() *MassContext {
	var p = new(MassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_mass
	return p
}

func (*MassContext) IsMassContext() {}

func NewMassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MassContext {
	var p = new(MassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_mass

	return p
}

func (s *MassContext) GetParser() antlr.Parser { return s.parser }

func (s *MassContext) MassUnits() IMassUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMassUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMassUnitsContext)
}

func (s *MassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterMass(s)
	}
}

func (s *MassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitMass(s)
	}
}

func (s *MassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitMass(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Mass() (localctx IMassContext) {
	localctx = NewMassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, KnotParserRULE_mass)

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
		p.SetState(178)
		p.Match(KnotParserT__31)
	}
	{
		p.SetState(179)
		p.MassUnits()
	}

	return localctx
}

// IMassUnitsContext is an interface to support dynamic dispatch.
type IMassUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsMassUnitsContext differentiates from other interfaces.
	IsMassUnitsContext()
}

type MassUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyMassUnitsContext() *MassUnitsContext {
	var p = new(MassUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_massUnits
	return p
}

func (*MassUnitsContext) IsMassUnitsContext() {}

func NewMassUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MassUnitsContext {
	var p = new(MassUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_massUnits

	return p
}

func (s *MassUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *MassUnitsContext) GetOp() antlr.Token { return s.op }

func (s *MassUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *MassUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MassUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MassUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterMassUnits(s)
	}
}

func (s *MassUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitMassUnits(s)
	}
}

func (s *MassUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitMassUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) MassUnits() (localctx IMassUnitsContext) {
	localctx = NewMassUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, KnotParserRULE_massUnits)
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
		p.SetState(181)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*MassUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(KnotParserT__32-33))|(1<<(KnotParserT__33-33))|(1<<(KnotParserT__34-33))|(1<<(KnotParserT__35-33)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*MassUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPressureContext is an interface to support dynamic dispatch.
type IPressureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPressureContext differentiates from other interfaces.
	IsPressureContext()
}

type PressureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPressureContext() *PressureContext {
	var p = new(PressureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_pressure
	return p
}

func (*PressureContext) IsPressureContext() {}

func NewPressureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PressureContext {
	var p = new(PressureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_pressure

	return p
}

func (s *PressureContext) GetParser() antlr.Parser { return s.parser }

func (s *PressureContext) PressureUnits() IPressureUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPressureUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPressureUnitsContext)
}

func (s *PressureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PressureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PressureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterPressure(s)
	}
}

func (s *PressureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitPressure(s)
	}
}

func (s *PressureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitPressure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Pressure() (localctx IPressureContext) {
	localctx = NewPressureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, KnotParserRULE_pressure)

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
		p.SetState(183)
		p.Match(KnotParserT__36)
	}
	{
		p.SetState(184)
		p.PressureUnits()
	}

	return localctx
}

// IPressureUnitsContext is an interface to support dynamic dispatch.
type IPressureUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsPressureUnitsContext differentiates from other interfaces.
	IsPressureUnitsContext()
}

type PressureUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyPressureUnitsContext() *PressureUnitsContext {
	var p = new(PressureUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_pressureUnits
	return p
}

func (*PressureUnitsContext) IsPressureUnitsContext() {}

func NewPressureUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PressureUnitsContext {
	var p = new(PressureUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_pressureUnits

	return p
}

func (s *PressureUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *PressureUnitsContext) GetOp() antlr.Token { return s.op }

func (s *PressureUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *PressureUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PressureUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PressureUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterPressureUnits(s)
	}
}

func (s *PressureUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitPressureUnits(s)
	}
}

func (s *PressureUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitPressureUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) PressureUnits() (localctx IPressureUnitsContext) {
	localctx = NewPressureUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, KnotParserRULE_pressureUnits)
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
		p.SetState(186)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*PressureUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(KnotParserT__37-38))|(1<<(KnotParserT__38-38))|(1<<(KnotParserT__39-38)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*PressureUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDistanceContext is an interface to support dynamic dispatch.
type IDistanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDistanceContext differentiates from other interfaces.
	IsDistanceContext()
}

type DistanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistanceContext() *DistanceContext {
	var p = new(DistanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_distance
	return p
}

func (*DistanceContext) IsDistanceContext() {}

func NewDistanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistanceContext {
	var p = new(DistanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_distance

	return p
}

func (s *DistanceContext) GetParser() antlr.Parser { return s.parser }

func (s *DistanceContext) DistanceUnits() IDistanceUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDistanceUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDistanceUnitsContext)
}

func (s *DistanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterDistance(s)
	}
}

func (s *DistanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitDistance(s)
	}
}

func (s *DistanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitDistance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Distance() (localctx IDistanceContext) {
	localctx = NewDistanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, KnotParserRULE_distance)

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
		p.SetState(188)
		p.Match(KnotParserT__40)
	}
	{
		p.SetState(189)
		p.DistanceUnits()
	}

	return localctx
}

// IDistanceUnitsContext is an interface to support dynamic dispatch.
type IDistanceUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsDistanceUnitsContext differentiates from other interfaces.
	IsDistanceUnitsContext()
}

type DistanceUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyDistanceUnitsContext() *DistanceUnitsContext {
	var p = new(DistanceUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_distanceUnits
	return p
}

func (*DistanceUnitsContext) IsDistanceUnitsContext() {}

func NewDistanceUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistanceUnitsContext {
	var p = new(DistanceUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_distanceUnits

	return p
}

func (s *DistanceUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *DistanceUnitsContext) GetOp() antlr.Token { return s.op }

func (s *DistanceUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *DistanceUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistanceUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterDistanceUnits(s)
	}
}

func (s *DistanceUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitDistanceUnits(s)
	}
}

func (s *DistanceUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitDistanceUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) DistanceUnits() (localctx IDistanceUnitsContext) {
	localctx = NewDistanceUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, KnotParserRULE_distanceUnits)
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
		p.SetState(191)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*DistanceUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(KnotParserT__41-42))|(1<<(KnotParserT__42-42))|(1<<(KnotParserT__43-42))|(1<<(KnotParserT__44-42)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*DistanceUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAngleContext is an interface to support dynamic dispatch.
type IAngleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAngleContext differentiates from other interfaces.
	IsAngleContext()
}

type AngleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAngleContext() *AngleContext {
	var p = new(AngleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_angle
	return p
}

func (*AngleContext) IsAngleContext() {}

func NewAngleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AngleContext {
	var p = new(AngleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_angle

	return p
}

func (s *AngleContext) GetParser() antlr.Parser { return s.parser }

func (s *AngleContext) AngleUnits() IAngleUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAngleUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAngleUnitsContext)
}

func (s *AngleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AngleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AngleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterAngle(s)
	}
}

func (s *AngleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitAngle(s)
	}
}

func (s *AngleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitAngle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Angle() (localctx IAngleContext) {
	localctx = NewAngleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, KnotParserRULE_angle)

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
		p.SetState(193)
		p.Match(KnotParserT__45)
	}
	{
		p.SetState(194)
		p.AngleUnits()
	}

	return localctx
}

// IAngleUnitsContext is an interface to support dynamic dispatch.
type IAngleUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsAngleUnitsContext differentiates from other interfaces.
	IsAngleUnitsContext()
}

type AngleUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAngleUnitsContext() *AngleUnitsContext {
	var p = new(AngleUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_angleUnits
	return p
}

func (*AngleUnitsContext) IsAngleUnitsContext() {}

func NewAngleUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AngleUnitsContext {
	var p = new(AngleUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_angleUnits

	return p
}

func (s *AngleUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *AngleUnitsContext) GetOp() antlr.Token { return s.op }

func (s *AngleUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *AngleUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AngleUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AngleUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterAngleUnits(s)
	}
}

func (s *AngleUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitAngleUnits(s)
	}
}

func (s *AngleUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitAngleUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) AngleUnits() (localctx IAngleUnitsContext) {
	localctx = NewAngleUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, KnotParserRULE_angleUnits)
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
		p.SetState(196)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AngleUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == KnotParserT__46 || _la == KnotParserT__47) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AngleUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVolumeContext is an interface to support dynamic dispatch.
type IVolumeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVolumeContext differentiates from other interfaces.
	IsVolumeContext()
}

type VolumeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVolumeContext() *VolumeContext {
	var p = new(VolumeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_volume
	return p
}

func (*VolumeContext) IsVolumeContext() {}

func NewVolumeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VolumeContext {
	var p = new(VolumeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_volume

	return p
}

func (s *VolumeContext) GetParser() antlr.Parser { return s.parser }

func (s *VolumeContext) VolumeUnits() IVolumeUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVolumeUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVolumeUnitsContext)
}

func (s *VolumeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VolumeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VolumeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterVolume(s)
	}
}

func (s *VolumeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitVolume(s)
	}
}

func (s *VolumeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitVolume(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Volume() (localctx IVolumeContext) {
	localctx = NewVolumeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, KnotParserRULE_volume)

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
		p.SetState(198)
		p.Match(KnotParserT__48)
	}
	{
		p.SetState(199)
		p.VolumeUnits()
	}

	return localctx
}

// IVolumeUnitsContext is an interface to support dynamic dispatch.
type IVolumeUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsVolumeUnitsContext differentiates from other interfaces.
	IsVolumeUnitsContext()
}

type VolumeUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyVolumeUnitsContext() *VolumeUnitsContext {
	var p = new(VolumeUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_volumeUnits
	return p
}

func (*VolumeUnitsContext) IsVolumeUnitsContext() {}

func NewVolumeUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VolumeUnitsContext {
	var p = new(VolumeUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_volumeUnits

	return p
}

func (s *VolumeUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *VolumeUnitsContext) GetOp() antlr.Token { return s.op }

func (s *VolumeUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *VolumeUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VolumeUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VolumeUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterVolumeUnits(s)
	}
}

func (s *VolumeUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitVolumeUnits(s)
	}
}

func (s *VolumeUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitVolumeUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) VolumeUnits() (localctx IVolumeUnitsContext) {
	localctx = NewVolumeUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, KnotParserRULE_volumeUnits)
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
		p.SetState(201)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VolumeUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(KnotParserT__49-50))|(1<<(KnotParserT__50-50))|(1<<(KnotParserT__51-50))|(1<<(KnotParserT__52-50)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VolumeUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAreaContext is an interface to support dynamic dispatch.
type IAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAreaContext differentiates from other interfaces.
	IsAreaContext()
}

type AreaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAreaContext() *AreaContext {
	var p = new(AreaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_area
	return p
}

func (*AreaContext) IsAreaContext() {}

func NewAreaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AreaContext {
	var p = new(AreaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_area

	return p
}

func (s *AreaContext) GetParser() antlr.Parser { return s.parser }

func (s *AreaContext) AreaUnits() IAreaUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAreaUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAreaUnitsContext)
}

func (s *AreaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AreaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AreaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterArea(s)
	}
}

func (s *AreaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitArea(s)
	}
}

func (s *AreaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitArea(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Area() (localctx IAreaContext) {
	localctx = NewAreaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, KnotParserRULE_area)

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
		p.SetState(203)
		p.Match(KnotParserT__53)
	}
	{
		p.SetState(204)
		p.AreaUnits()
	}

	return localctx
}

// IAreaUnitsContext is an interface to support dynamic dispatch.
type IAreaUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsAreaUnitsContext differentiates from other interfaces.
	IsAreaUnitsContext()
}

type AreaUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAreaUnitsContext() *AreaUnitsContext {
	var p = new(AreaUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_areaUnits
	return p
}

func (*AreaUnitsContext) IsAreaUnitsContext() {}

func NewAreaUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AreaUnitsContext {
	var p = new(AreaUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_areaUnits

	return p
}

func (s *AreaUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *AreaUnitsContext) GetOp() antlr.Token { return s.op }

func (s *AreaUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *AreaUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AreaUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AreaUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterAreaUnits(s)
	}
}

func (s *AreaUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitAreaUnits(s)
	}
}

func (s *AreaUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitAreaUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) AreaUnits() (localctx IAreaUnitsContext) {
	localctx = NewAreaUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, KnotParserRULE_areaUnits)
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
		p.SetState(206)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AreaUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(KnotParserT__54-55))|(1<<(KnotParserT__55-55))|(1<<(KnotParserT__56-55))|(1<<(KnotParserT__57-55)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AreaUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRainContext is an interface to support dynamic dispatch.
type IRainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRainContext differentiates from other interfaces.
	IsRainContext()
}

type RainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRainContext() *RainContext {
	var p = new(RainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_rain
	return p
}

func (*RainContext) IsRainContext() {}

func NewRainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RainContext {
	var p = new(RainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_rain

	return p
}

func (s *RainContext) GetParser() antlr.Parser { return s.parser }

func (s *RainContext) RainUnits() IRainUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRainUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRainUnitsContext)
}

func (s *RainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterRain(s)
	}
}

func (s *RainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitRain(s)
	}
}

func (s *RainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitRain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Rain() (localctx IRainContext) {
	localctx = NewRainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, KnotParserRULE_rain)

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
		p.SetState(208)
		p.Match(KnotParserT__58)
	}
	{
		p.SetState(209)
		p.RainUnits()
	}

	return localctx
}

// IRainUnitsContext is an interface to support dynamic dispatch.
type IRainUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsRainUnitsContext differentiates from other interfaces.
	IsRainUnitsContext()
}

type RainUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyRainUnitsContext() *RainUnitsContext {
	var p = new(RainUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_rainUnits
	return p
}

func (*RainUnitsContext) IsRainUnitsContext() {}

func NewRainUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RainUnitsContext {
	var p = new(RainUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_rainUnits

	return p
}

func (s *RainUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *RainUnitsContext) GetOp() antlr.Token { return s.op }

func (s *RainUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *RainUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RainUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RainUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterRainUnits(s)
	}
}

func (s *RainUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitRainUnits(s)
	}
}

func (s *RainUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitRainUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) RainUnits() (localctx IRainUnitsContext) {
	localctx = NewRainUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, KnotParserRULE_rainUnits)

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
		p.SetState(211)

		var _m = p.Match(KnotParserT__57)

		localctx.(*RainUnitsContext).op = _m
	}

	return localctx
}

// IDensityContext is an interface to support dynamic dispatch.
type IDensityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDensityContext differentiates from other interfaces.
	IsDensityContext()
}

type DensityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDensityContext() *DensityContext {
	var p = new(DensityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_density
	return p
}

func (*DensityContext) IsDensityContext() {}

func NewDensityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DensityContext {
	var p = new(DensityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_density

	return p
}

func (s *DensityContext) GetParser() antlr.Parser { return s.parser }

func (s *DensityContext) DensityUnits() IDensityUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDensityUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDensityUnitsContext)
}

func (s *DensityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DensityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DensityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterDensity(s)
	}
}

func (s *DensityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitDensity(s)
	}
}

func (s *DensityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitDensity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Density() (localctx IDensityContext) {
	localctx = NewDensityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, KnotParserRULE_density)

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
		p.SetState(213)
		p.Match(KnotParserT__59)
	}
	{
		p.SetState(214)
		p.DensityUnits()
	}

	return localctx
}

// IDensityUnitsContext is an interface to support dynamic dispatch.
type IDensityUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsDensityUnitsContext differentiates from other interfaces.
	IsDensityUnitsContext()
}

type DensityUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyDensityUnitsContext() *DensityUnitsContext {
	var p = new(DensityUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_densityUnits
	return p
}

func (*DensityUnitsContext) IsDensityUnitsContext() {}

func NewDensityUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DensityUnitsContext {
	var p = new(DensityUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_densityUnits

	return p
}

func (s *DensityUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *DensityUnitsContext) GetOp() antlr.Token { return s.op }

func (s *DensityUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *DensityUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DensityUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DensityUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterDensityUnits(s)
	}
}

func (s *DensityUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitDensityUnits(s)
	}
}

func (s *DensityUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitDensityUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) DensityUnits() (localctx IDensityUnitsContext) {
	localctx = NewDensityUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, KnotParserRULE_densityUnits)

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
		p.SetState(216)

		var _m = p.Match(KnotParserT__60)

		localctx.(*DensityUnitsContext).op = _m
	}

	return localctx
}

// ILatitudeContext is an interface to support dynamic dispatch.
type ILatitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLatitudeContext differentiates from other interfaces.
	IsLatitudeContext()
}

type LatitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLatitudeContext() *LatitudeContext {
	var p = new(LatitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_latitude
	return p
}

func (*LatitudeContext) IsLatitudeContext() {}

func NewLatitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LatitudeContext {
	var p = new(LatitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_latitude

	return p
}

func (s *LatitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *LatitudeContext) LatitudeUnits() ILatitudeUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILatitudeUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILatitudeUnitsContext)
}

func (s *LatitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LatitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LatitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterLatitude(s)
	}
}

func (s *LatitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitLatitude(s)
	}
}

func (s *LatitudeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitLatitude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Latitude() (localctx ILatitudeContext) {
	localctx = NewLatitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, KnotParserRULE_latitude)

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
		p.SetState(218)
		p.Match(KnotParserT__61)
	}
	{
		p.SetState(219)
		p.LatitudeUnits()
	}

	return localctx
}

// ILatitudeUnitsContext is an interface to support dynamic dispatch.
type ILatitudeUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsLatitudeUnitsContext differentiates from other interfaces.
	IsLatitudeUnitsContext()
}

type LatitudeUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyLatitudeUnitsContext() *LatitudeUnitsContext {
	var p = new(LatitudeUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_latitudeUnits
	return p
}

func (*LatitudeUnitsContext) IsLatitudeUnitsContext() {}

func NewLatitudeUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LatitudeUnitsContext {
	var p = new(LatitudeUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_latitudeUnits

	return p
}

func (s *LatitudeUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *LatitudeUnitsContext) GetOp() antlr.Token { return s.op }

func (s *LatitudeUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *LatitudeUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LatitudeUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LatitudeUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterLatitudeUnits(s)
	}
}

func (s *LatitudeUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitLatitudeUnits(s)
	}
}

func (s *LatitudeUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitLatitudeUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) LatitudeUnits() (localctx ILatitudeUnitsContext) {
	localctx = NewLatitudeUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, KnotParserRULE_latitudeUnits)

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
		p.SetState(221)

		var _m = p.Match(KnotParserT__62)

		localctx.(*LatitudeUnitsContext).op = _m
	}

	return localctx
}

// ILongitudeContext is an interface to support dynamic dispatch.
type ILongitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLongitudeContext differentiates from other interfaces.
	IsLongitudeContext()
}

type LongitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLongitudeContext() *LongitudeContext {
	var p = new(LongitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_longitude
	return p
}

func (*LongitudeContext) IsLongitudeContext() {}

func NewLongitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LongitudeContext {
	var p = new(LongitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_longitude

	return p
}

func (s *LongitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *LongitudeContext) LongitudeUnits() ILongitudeUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILongitudeUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILongitudeUnitsContext)
}

func (s *LongitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LongitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterLongitude(s)
	}
}

func (s *LongitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitLongitude(s)
	}
}

func (s *LongitudeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitLongitude(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Longitude() (localctx ILongitudeContext) {
	localctx = NewLongitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, KnotParserRULE_longitude)

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
		p.SetState(223)
		p.Match(KnotParserT__63)
	}
	{
		p.SetState(224)
		p.LongitudeUnits()
	}

	return localctx
}

// ILongitudeUnitsContext is an interface to support dynamic dispatch.
type ILongitudeUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsLongitudeUnitsContext differentiates from other interfaces.
	IsLongitudeUnitsContext()
}

type LongitudeUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyLongitudeUnitsContext() *LongitudeUnitsContext {
	var p = new(LongitudeUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_longitudeUnits
	return p
}

func (*LongitudeUnitsContext) IsLongitudeUnitsContext() {}

func NewLongitudeUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LongitudeUnitsContext {
	var p = new(LongitudeUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_longitudeUnits

	return p
}

func (s *LongitudeUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *LongitudeUnitsContext) GetOp() antlr.Token { return s.op }

func (s *LongitudeUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *LongitudeUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongitudeUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LongitudeUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterLongitudeUnits(s)
	}
}

func (s *LongitudeUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitLongitudeUnits(s)
	}
}

func (s *LongitudeUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitLongitudeUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) LongitudeUnits() (localctx ILongitudeUnitsContext) {
	localctx = NewLongitudeUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, KnotParserRULE_longitudeUnits)

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
		p.SetState(226)

		var _m = p.Match(KnotParserT__62)

		localctx.(*LongitudeUnitsContext).op = _m
	}

	return localctx
}

// ISpeedContext is an interface to support dynamic dispatch.
type ISpeedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpeedContext differentiates from other interfaces.
	IsSpeedContext()
}

type SpeedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpeedContext() *SpeedContext {
	var p = new(SpeedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_speed
	return p
}

func (*SpeedContext) IsSpeedContext() {}

func NewSpeedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpeedContext {
	var p = new(SpeedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_speed

	return p
}

func (s *SpeedContext) GetParser() antlr.Parser { return s.parser }

func (s *SpeedContext) SpeedUnits() ISpeedUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpeedUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpeedUnitsContext)
}

func (s *SpeedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpeedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpeedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterSpeed(s)
	}
}

func (s *SpeedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitSpeed(s)
	}
}

func (s *SpeedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitSpeed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Speed() (localctx ISpeedContext) {
	localctx = NewSpeedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, KnotParserRULE_speed)

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
		p.SetState(228)
		p.Match(KnotParserT__64)
	}
	{
		p.SetState(229)
		p.SpeedUnits()
	}

	return localctx
}

// ISpeedUnitsContext is an interface to support dynamic dispatch.
type ISpeedUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsSpeedUnitsContext differentiates from other interfaces.
	IsSpeedUnitsContext()
}

type SpeedUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptySpeedUnitsContext() *SpeedUnitsContext {
	var p = new(SpeedUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_speedUnits
	return p
}

func (*SpeedUnitsContext) IsSpeedUnitsContext() {}

func NewSpeedUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpeedUnitsContext {
	var p = new(SpeedUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_speedUnits

	return p
}

func (s *SpeedUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *SpeedUnitsContext) GetOp() antlr.Token { return s.op }

func (s *SpeedUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *SpeedUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpeedUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpeedUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterSpeedUnits(s)
	}
}

func (s *SpeedUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitSpeedUnits(s)
	}
}

func (s *SpeedUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitSpeedUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) SpeedUnits() (localctx ISpeedUnitsContext) {
	localctx = NewSpeedUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, KnotParserRULE_speedUnits)
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
		p.SetState(231)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*SpeedUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == KnotParserT__27 || (((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(KnotParserT__65-66))|(1<<(KnotParserT__66-66))|(1<<(KnotParserT__67-66)))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*SpeedUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVolumeflowContext is an interface to support dynamic dispatch.
type IVolumeflowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVolumeflowContext differentiates from other interfaces.
	IsVolumeflowContext()
}

type VolumeflowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVolumeflowContext() *VolumeflowContext {
	var p = new(VolumeflowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_volumeflow
	return p
}

func (*VolumeflowContext) IsVolumeflowContext() {}

func NewVolumeflowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VolumeflowContext {
	var p = new(VolumeflowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_volumeflow

	return p
}

func (s *VolumeflowContext) GetParser() antlr.Parser { return s.parser }

func (s *VolumeflowContext) VolumeflowUnits() IVolumeflowUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVolumeflowUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVolumeflowUnitsContext)
}

func (s *VolumeflowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VolumeflowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VolumeflowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterVolumeflow(s)
	}
}

func (s *VolumeflowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitVolumeflow(s)
	}
}

func (s *VolumeflowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitVolumeflow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Volumeflow() (localctx IVolumeflowContext) {
	localctx = NewVolumeflowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, KnotParserRULE_volumeflow)

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
		p.SetState(233)
		p.Match(KnotParserT__68)
	}
	{
		p.SetState(234)
		p.VolumeflowUnits()
	}

	return localctx
}

// IVolumeflowUnitsContext is an interface to support dynamic dispatch.
type IVolumeflowUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsVolumeflowUnitsContext differentiates from other interfaces.
	IsVolumeflowUnitsContext()
}

type VolumeflowUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyVolumeflowUnitsContext() *VolumeflowUnitsContext {
	var p = new(VolumeflowUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_volumeflowUnits
	return p
}

func (*VolumeflowUnitsContext) IsVolumeflowUnitsContext() {}

func NewVolumeflowUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VolumeflowUnitsContext {
	var p = new(VolumeflowUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_volumeflowUnits

	return p
}

func (s *VolumeflowUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *VolumeflowUnitsContext) GetOp() antlr.Token { return s.op }

func (s *VolumeflowUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *VolumeflowUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VolumeflowUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VolumeflowUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterVolumeflowUnits(s)
	}
}

func (s *VolumeflowUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitVolumeflowUnits(s)
	}
}

func (s *VolumeflowUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitVolumeflowUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) VolumeflowUnits() (localctx IVolumeflowUnitsContext) {
	localctx = NewVolumeflowUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, KnotParserRULE_volumeflowUnits)
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
		p.SetState(236)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*VolumeflowUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == KnotParserT__22 || (((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(KnotParserT__69-70))|(1<<(KnotParserT__70-70))|(1<<(KnotParserT__71-70))|(1<<(KnotParserT__72-70))|(1<<(KnotParserT__73-70)))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*VolumeflowUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEnergyContext is an interface to support dynamic dispatch.
type IEnergyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnergyContext differentiates from other interfaces.
	IsEnergyContext()
}

type EnergyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnergyContext() *EnergyContext {
	var p = new(EnergyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_energy
	return p
}

func (*EnergyContext) IsEnergyContext() {}

func NewEnergyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnergyContext {
	var p = new(EnergyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_energy

	return p
}

func (s *EnergyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnergyContext) EnergyUnits() IEnergyUnitsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnergyUnitsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnergyUnitsContext)
}

func (s *EnergyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnergyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnergyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterEnergy(s)
	}
}

func (s *EnergyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitEnergy(s)
	}
}

func (s *EnergyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitEnergy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) Energy() (localctx IEnergyContext) {
	localctx = NewEnergyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, KnotParserRULE_energy)

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
		p.SetState(238)
		p.Match(KnotParserT__74)
	}
	{
		p.SetState(239)
		p.EnergyUnits()
	}

	return localctx
}

// IEnergyUnitsContext is an interface to support dynamic dispatch.
type IEnergyUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsEnergyUnitsContext differentiates from other interfaces.
	IsEnergyUnitsContext()
}

type EnergyUnitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyEnergyUnitsContext() *EnergyUnitsContext {
	var p = new(EnergyUnitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KnotParserRULE_energyUnits
	return p
}

func (*EnergyUnitsContext) IsEnergyUnitsContext() {}

func NewEnergyUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnergyUnitsContext {
	var p = new(EnergyUnitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KnotParserRULE_energyUnits

	return p
}

func (s *EnergyUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnergyUnitsContext) GetOp() antlr.Token { return s.op }

func (s *EnergyUnitsContext) SetOp(v antlr.Token) { s.op = v }

func (s *EnergyUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnergyUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnergyUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.EnterEnergyUnits(s)
	}
}

func (s *EnergyUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KnotListener); ok {
		listenerT.ExitEnergyUnits(s)
	}
}

func (s *EnergyUnitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KnotVisitor:
		return t.VisitEnergyUnits(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KnotParser) EnergyUnits() (localctx IEnergyUnitsContext) {
	localctx = NewEnergyUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, KnotParserRULE_energyUnits)
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
		p.SetState(241)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*EnergyUnitsContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(KnotParserT__75-76))|(1<<(KnotParserT__76-76))|(1<<(KnotParserT__77-76))|(1<<(KnotParserT__78-76))|(1<<(KnotParserT__79-76))|(1<<(KnotParserT__80-76)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*EnergyUnitsContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
