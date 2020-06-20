// Code generated from Knot.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // Knot
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by KnotParser.
type KnotVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KnotParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by KnotParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by KnotParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by KnotParser#innerContent.
	VisitInnerContent(ctx *InnerContentContext) interface{}

	// Visit a parse tree produced by KnotParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by KnotParser#valueOptions.
	VisitValueOptions(ctx *ValueOptionsContext) interface{}

	// Visit a parse tree produced by KnotParser#typeRule.
	VisitTypeRule(ctx *TypeRuleContext) interface{}

	// Visit a parse tree produced by KnotParser#typeOptions.
	VisitTypeOptions(ctx *TypeOptionsContext) interface{}

	// Visit a parse tree produced by KnotParser#voltage.
	VisitVoltage(ctx *VoltageContext) interface{}

	// Visit a parse tree produced by KnotParser#temperature.
	VisitTemperature(ctx *TemperatureContext) interface{}

	// Visit a parse tree produced by KnotParser#voltagesUnits.
	VisitVoltagesUnits(ctx *VoltagesUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#temperatureUnits.
	VisitTemperatureUnits(ctx *TemperatureUnitsContext) interface{}
}
