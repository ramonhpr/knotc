// Code generated from Knot.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // Knot
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseKnotVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseKnotVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitInnerContent(ctx *InnerContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitValueOptions(ctx *ValueOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitTypeRule(ctx *TypeRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitTypeOptions(ctx *TypeOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitVoltage(ctx *VoltageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitTemperature(ctx *TemperatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitVoltagesUnits(ctx *VoltagesUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseKnotVisitor) VisitTemperatureUnits(ctx *TemperatureUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}
