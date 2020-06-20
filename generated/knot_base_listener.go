// Code generated from Knot.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // Knot
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKnotListener is a complete listener for a parse tree produced by KnotParser.
type BaseKnotListener struct{}

var _ KnotListener = &BaseKnotListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKnotListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKnotListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKnotListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKnotListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseKnotListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseKnotListener) ExitStart(ctx *StartContext) {}

// EnterName is called when production name is entered.
func (s *BaseKnotListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseKnotListener) ExitName(ctx *NameContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseKnotListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseKnotListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterInnerContent is called when production innerContent is entered.
func (s *BaseKnotListener) EnterInnerContent(ctx *InnerContentContext) {}

// ExitInnerContent is called when production innerContent is exited.
func (s *BaseKnotListener) ExitInnerContent(ctx *InnerContentContext) {}

// EnterValue is called when production value is entered.
func (s *BaseKnotListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseKnotListener) ExitValue(ctx *ValueContext) {}

// EnterValueOptions is called when production valueOptions is entered.
func (s *BaseKnotListener) EnterValueOptions(ctx *ValueOptionsContext) {}

// ExitValueOptions is called when production valueOptions is exited.
func (s *BaseKnotListener) ExitValueOptions(ctx *ValueOptionsContext) {}

// EnterTypeRule is called when production typeRule is entered.
func (s *BaseKnotListener) EnterTypeRule(ctx *TypeRuleContext) {}

// ExitTypeRule is called when production typeRule is exited.
func (s *BaseKnotListener) ExitTypeRule(ctx *TypeRuleContext) {}

// EnterTypeOptions is called when production typeOptions is entered.
func (s *BaseKnotListener) EnterTypeOptions(ctx *TypeOptionsContext) {}

// ExitTypeOptions is called when production typeOptions is exited.
func (s *BaseKnotListener) ExitTypeOptions(ctx *TypeOptionsContext) {}

// EnterVoltage is called when production voltage is entered.
func (s *BaseKnotListener) EnterVoltage(ctx *VoltageContext) {}

// ExitVoltage is called when production voltage is exited.
func (s *BaseKnotListener) ExitVoltage(ctx *VoltageContext) {}

// EnterTemperature is called when production temperature is entered.
func (s *BaseKnotListener) EnterTemperature(ctx *TemperatureContext) {}

// ExitTemperature is called when production temperature is exited.
func (s *BaseKnotListener) ExitTemperature(ctx *TemperatureContext) {}

// EnterVoltagesUnits is called when production voltagesUnits is entered.
func (s *BaseKnotListener) EnterVoltagesUnits(ctx *VoltagesUnitsContext) {}

// ExitVoltagesUnits is called when production voltagesUnits is exited.
func (s *BaseKnotListener) ExitVoltagesUnits(ctx *VoltagesUnitsContext) {}

// EnterTemperatureUnits is called when production temperatureUnits is entered.
func (s *BaseKnotListener) EnterTemperatureUnits(ctx *TemperatureUnitsContext) {}

// ExitTemperatureUnits is called when production temperatureUnits is exited.
func (s *BaseKnotListener) ExitTemperatureUnits(ctx *TemperatureUnitsContext) {}
