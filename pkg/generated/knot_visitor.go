// Code generated from Knot.g4 by ANTLR 4.8. DO NOT EDIT.

package generated // Knot
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by KnotParser.
type KnotVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KnotParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by KnotParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by KnotParser#thingContent.
	VisitThingContent(ctx *ThingContentContext) interface{}

	// Visit a parse tree produced by KnotParser#config.
	VisitConfig(ctx *ConfigContext) interface{}

	// Visit a parse tree produced by KnotParser#configChanges.
	VisitConfigChanges(ctx *ConfigChangesContext) interface{}

	// Visit a parse tree produced by KnotParser#configTime.
	VisitConfigTime(ctx *ConfigTimeContext) interface{}

	// Visit a parse tree produced by KnotParser#configUpper.
	VisitConfigUpper(ctx *ConfigUpperContext) interface{}

	// Visit a parse tree produced by KnotParser#configLower.
	VisitConfigLower(ctx *ConfigLowerContext) interface{}

	// Visit a parse tree produced by KnotParser#valueOptions.
	VisitValueOptions(ctx *ValueOptionsContext) interface{}

	// Visit a parse tree produced by KnotParser#boolOpt.
	VisitBoolOpt(ctx *BoolOptContext) interface{}

	// Visit a parse tree produced by KnotParser#numberOpt.
	VisitNumberOpt(ctx *NumberOptContext) interface{}

	// Visit a parse tree produced by KnotParser#bytesOpt.
	VisitBytesOpt(ctx *BytesOptContext) interface{}

	// Visit a parse tree produced by KnotParser#unitTypeOptions.
	VisitUnitTypeOptions(ctx *UnitTypeOptionsContext) interface{}

	// Visit a parse tree produced by KnotParser#logicUnits.
	VisitLogicUnits(ctx *LogicUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#voltage.
	VisitVoltage(ctx *VoltageContext) interface{}

	// Visit a parse tree produced by KnotParser#voltagesUnits.
	VisitVoltagesUnits(ctx *VoltagesUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#current.
	VisitCurrent(ctx *CurrentContext) interface{}

	// Visit a parse tree produced by KnotParser#currentUnits.
	VisitCurrentUnits(ctx *CurrentUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#resistance.
	VisitResistance(ctx *ResistanceContext) interface{}

	// Visit a parse tree produced by KnotParser#resistanceUnits.
	VisitResistanceUnits(ctx *ResistanceUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by KnotParser#powerUnits.
	VisitPowerUnits(ctx *PowerUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#temperature.
	VisitTemperature(ctx *TemperatureContext) interface{}

	// Visit a parse tree produced by KnotParser#temperatureUnits.
	VisitTemperatureUnits(ctx *TemperatureUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#luminosity.
	VisitLuminosity(ctx *LuminosityContext) interface{}

	// Visit a parse tree produced by KnotParser#luminosityUnits.
	VisitLuminosityUnits(ctx *LuminosityUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#time.
	VisitTime(ctx *TimeContext) interface{}

	// Visit a parse tree produced by KnotParser#timeUnits.
	VisitTimeUnits(ctx *TimeUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#mass.
	VisitMass(ctx *MassContext) interface{}

	// Visit a parse tree produced by KnotParser#massUnits.
	VisitMassUnits(ctx *MassUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#pressure.
	VisitPressure(ctx *PressureContext) interface{}

	// Visit a parse tree produced by KnotParser#pressureUnits.
	VisitPressureUnits(ctx *PressureUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#distance.
	VisitDistance(ctx *DistanceContext) interface{}

	// Visit a parse tree produced by KnotParser#distanceUnits.
	VisitDistanceUnits(ctx *DistanceUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#angle.
	VisitAngle(ctx *AngleContext) interface{}

	// Visit a parse tree produced by KnotParser#angleUnits.
	VisitAngleUnits(ctx *AngleUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#volume.
	VisitVolume(ctx *VolumeContext) interface{}

	// Visit a parse tree produced by KnotParser#volumeUnits.
	VisitVolumeUnits(ctx *VolumeUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#area.
	VisitArea(ctx *AreaContext) interface{}

	// Visit a parse tree produced by KnotParser#areaUnits.
	VisitAreaUnits(ctx *AreaUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#rain.
	VisitRain(ctx *RainContext) interface{}

	// Visit a parse tree produced by KnotParser#rainUnits.
	VisitRainUnits(ctx *RainUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#density.
	VisitDensity(ctx *DensityContext) interface{}

	// Visit a parse tree produced by KnotParser#densityUnits.
	VisitDensityUnits(ctx *DensityUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#latitude.
	VisitLatitude(ctx *LatitudeContext) interface{}

	// Visit a parse tree produced by KnotParser#latitudeUnits.
	VisitLatitudeUnits(ctx *LatitudeUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#longitude.
	VisitLongitude(ctx *LongitudeContext) interface{}

	// Visit a parse tree produced by KnotParser#longitudeUnits.
	VisitLongitudeUnits(ctx *LongitudeUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#speed.
	VisitSpeed(ctx *SpeedContext) interface{}

	// Visit a parse tree produced by KnotParser#speedUnits.
	VisitSpeedUnits(ctx *SpeedUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#volumeflow.
	VisitVolumeflow(ctx *VolumeflowContext) interface{}

	// Visit a parse tree produced by KnotParser#volumeflowUnits.
	VisitVolumeflowUnits(ctx *VolumeflowUnitsContext) interface{}

	// Visit a parse tree produced by KnotParser#energy.
	VisitEnergy(ctx *EnergyContext) interface{}

	// Visit a parse tree produced by KnotParser#energyUnits.
	VisitEnergyUnits(ctx *EnergyUnitsContext) interface{}
}
