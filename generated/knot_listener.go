// Code generated from Knot.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // Knot
import "github.com/antlr/antlr4/runtime/Go/antlr"

// KnotListener is a complete listener for a parse tree produced by KnotParser.
type KnotListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterThingContent is called when entering the thingContent production.
	EnterThingContent(c *ThingContentContext)

	// EnterValueOptions is called when entering the valueOptions production.
	EnterValueOptions(c *ValueOptionsContext)

	// EnterUnitTypeOptions is called when entering the unitTypeOptions production.
	EnterUnitTypeOptions(c *UnitTypeOptionsContext)

	// EnterVoltage is called when entering the voltage production.
	EnterVoltage(c *VoltageContext)

	// EnterVoltagesUnits is called when entering the voltagesUnits production.
	EnterVoltagesUnits(c *VoltagesUnitsContext)

	// EnterTemperature is called when entering the temperature production.
	EnterTemperature(c *TemperatureContext)

	// EnterTemperatureUnits is called when entering the temperatureUnits production.
	EnterTemperatureUnits(c *TemperatureUnitsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitThingContent is called when exiting the thingContent production.
	ExitThingContent(c *ThingContentContext)

	// ExitValueOptions is called when exiting the valueOptions production.
	ExitValueOptions(c *ValueOptionsContext)

	// ExitUnitTypeOptions is called when exiting the unitTypeOptions production.
	ExitUnitTypeOptions(c *UnitTypeOptionsContext)

	// ExitVoltage is called when exiting the voltage production.
	ExitVoltage(c *VoltageContext)

	// ExitVoltagesUnits is called when exiting the voltagesUnits production.
	ExitVoltagesUnits(c *VoltagesUnitsContext)

	// ExitTemperature is called when exiting the temperature production.
	ExitTemperature(c *TemperatureContext)

	// ExitTemperatureUnits is called when exiting the temperatureUnits production.
	ExitTemperatureUnits(c *TemperatureUnitsContext)
}
