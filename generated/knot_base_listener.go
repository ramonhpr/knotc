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

// EnterDefinition is called when production definition is entered.
func (s *BaseKnotListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseKnotListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterThingContent is called when production thingContent is entered.
func (s *BaseKnotListener) EnterThingContent(ctx *ThingContentContext) {}

// ExitThingContent is called when production thingContent is exited.
func (s *BaseKnotListener) ExitThingContent(ctx *ThingContentContext) {}

// EnterValueOptions is called when production valueOptions is entered.
func (s *BaseKnotListener) EnterValueOptions(ctx *ValueOptionsContext) {}

// ExitValueOptions is called when production valueOptions is exited.
func (s *BaseKnotListener) ExitValueOptions(ctx *ValueOptionsContext) {}

// EnterUnitTypeOptions is called when production unitTypeOptions is entered.
func (s *BaseKnotListener) EnterUnitTypeOptions(ctx *UnitTypeOptionsContext) {}

// ExitUnitTypeOptions is called when production unitTypeOptions is exited.
func (s *BaseKnotListener) ExitUnitTypeOptions(ctx *UnitTypeOptionsContext) {}

// EnterVoltage is called when production voltage is entered.
func (s *BaseKnotListener) EnterVoltage(ctx *VoltageContext) {}

// ExitVoltage is called when production voltage is exited.
func (s *BaseKnotListener) ExitVoltage(ctx *VoltageContext) {}

// EnterVoltagesUnits is called when production voltagesUnits is entered.
func (s *BaseKnotListener) EnterVoltagesUnits(ctx *VoltagesUnitsContext) {}

// ExitVoltagesUnits is called when production voltagesUnits is exited.
func (s *BaseKnotListener) ExitVoltagesUnits(ctx *VoltagesUnitsContext) {}

// EnterTemperature is called when production temperature is entered.
func (s *BaseKnotListener) EnterTemperature(ctx *TemperatureContext) {}

// ExitTemperature is called when production temperature is exited.
func (s *BaseKnotListener) ExitTemperature(ctx *TemperatureContext) {}

// EnterTemperatureUnits is called when production temperatureUnits is entered.
func (s *BaseKnotListener) EnterTemperatureUnits(ctx *TemperatureUnitsContext) {}

// ExitTemperatureUnits is called when production temperatureUnits is exited.
func (s *BaseKnotListener) ExitTemperatureUnits(ctx *TemperatureUnitsContext) {}
