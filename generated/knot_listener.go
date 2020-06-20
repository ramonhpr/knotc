// Code generated from Knot.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // Knot
import "github.com/antlr/antlr4/runtime/Go/antlr"

// KnotListener is a complete listener for a parse tree produced by KnotParser.
type KnotListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterInnerContent is called when entering the innerContent production.
	EnterInnerContent(c *InnerContentContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterValueOptions is called when entering the valueOptions production.
	EnterValueOptions(c *ValueOptionsContext)

	// EnterTypeRule is called when entering the typeRule production.
	EnterTypeRule(c *TypeRuleContext)

	// EnterTypeOptions is called when entering the typeOptions production.
	EnterTypeOptions(c *TypeOptionsContext)

	// EnterVoltage is called when entering the voltage production.
	EnterVoltage(c *VoltageContext)

	// EnterTemperature is called when entering the temperature production.
	EnterTemperature(c *TemperatureContext)

	// EnterVoltagesUnits is called when entering the voltagesUnits production.
	EnterVoltagesUnits(c *VoltagesUnitsContext)

	// EnterTemperatureUnits is called when entering the temperatureUnits production.
	EnterTemperatureUnits(c *TemperatureUnitsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitInnerContent is called when exiting the innerContent production.
	ExitInnerContent(c *InnerContentContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitValueOptions is called when exiting the valueOptions production.
	ExitValueOptions(c *ValueOptionsContext)

	// ExitTypeRule is called when exiting the typeRule production.
	ExitTypeRule(c *TypeRuleContext)

	// ExitTypeOptions is called when exiting the typeOptions production.
	ExitTypeOptions(c *TypeOptionsContext)

	// ExitVoltage is called when exiting the voltage production.
	ExitVoltage(c *VoltageContext)

	// ExitTemperature is called when exiting the temperature production.
	ExitTemperature(c *TemperatureContext)

	// ExitVoltagesUnits is called when exiting the voltagesUnits production.
	ExitVoltagesUnits(c *VoltagesUnitsContext)

	// ExitTemperatureUnits is called when exiting the temperatureUnits production.
	ExitTemperatureUnits(c *TemperatureUnitsContext)
}
