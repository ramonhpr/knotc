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

	// EnterConfigChanges is called when entering the configChanges production.
	EnterConfigChanges(c *ConfigChangesContext)

	// EnterConfigTime is called when entering the configTime production.
	EnterConfigTime(c *ConfigTimeContext)

	// EnterConfigUpper is called when entering the configUpper production.
	EnterConfigUpper(c *ConfigUpperContext)

	// EnterConfigLower is called when entering the configLower production.
	EnterConfigLower(c *ConfigLowerContext)

	// EnterUnitTypeOptions is called when entering the unitTypeOptions production.
	EnterUnitTypeOptions(c *UnitTypeOptionsContext)

	// EnterLogicUnits is called when entering the logicUnits production.
	EnterLogicUnits(c *LogicUnitsContext)

	// EnterVoltage is called when entering the voltage production.
	EnterVoltage(c *VoltageContext)

	// EnterVoltagesUnits is called when entering the voltagesUnits production.
	EnterVoltagesUnits(c *VoltagesUnitsContext)

	// EnterCurrent is called when entering the current production.
	EnterCurrent(c *CurrentContext)

	// EnterCurrentUnits is called when entering the currentUnits production.
	EnterCurrentUnits(c *CurrentUnitsContext)

	// EnterResistance is called when entering the resistance production.
	EnterResistance(c *ResistanceContext)

	// EnterResistanceUnits is called when entering the resistanceUnits production.
	EnterResistanceUnits(c *ResistanceUnitsContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterPowerUnits is called when entering the powerUnits production.
	EnterPowerUnits(c *PowerUnitsContext)

	// EnterTemperature is called when entering the temperature production.
	EnterTemperature(c *TemperatureContext)

	// EnterTemperatureUnits is called when entering the temperatureUnits production.
	EnterTemperatureUnits(c *TemperatureUnitsContext)

	// EnterLuminosity is called when entering the luminosity production.
	EnterLuminosity(c *LuminosityContext)

	// EnterLuminosityUnits is called when entering the luminosityUnits production.
	EnterLuminosityUnits(c *LuminosityUnitsContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// EnterTimeUnits is called when entering the timeUnits production.
	EnterTimeUnits(c *TimeUnitsContext)

	// EnterMass is called when entering the mass production.
	EnterMass(c *MassContext)

	// EnterMassUnits is called when entering the massUnits production.
	EnterMassUnits(c *MassUnitsContext)

	// EnterPressure is called when entering the pressure production.
	EnterPressure(c *PressureContext)

	// EnterPressureUnits is called when entering the pressureUnits production.
	EnterPressureUnits(c *PressureUnitsContext)

	// EnterDistance is called when entering the distance production.
	EnterDistance(c *DistanceContext)

	// EnterDistanceUnits is called when entering the distanceUnits production.
	EnterDistanceUnits(c *DistanceUnitsContext)

	// EnterAngle is called when entering the angle production.
	EnterAngle(c *AngleContext)

	// EnterAngleUnits is called when entering the angleUnits production.
	EnterAngleUnits(c *AngleUnitsContext)

	// EnterVolume is called when entering the volume production.
	EnterVolume(c *VolumeContext)

	// EnterVolumeUnits is called when entering the volumeUnits production.
	EnterVolumeUnits(c *VolumeUnitsContext)

	// EnterArea is called when entering the area production.
	EnterArea(c *AreaContext)

	// EnterAreaUnits is called when entering the areaUnits production.
	EnterAreaUnits(c *AreaUnitsContext)

	// EnterRain is called when entering the rain production.
	EnterRain(c *RainContext)

	// EnterRainUnits is called when entering the rainUnits production.
	EnterRainUnits(c *RainUnitsContext)

	// EnterDensity is called when entering the density production.
	EnterDensity(c *DensityContext)

	// EnterDensityUnits is called when entering the densityUnits production.
	EnterDensityUnits(c *DensityUnitsContext)

	// EnterLatitude is called when entering the latitude production.
	EnterLatitude(c *LatitudeContext)

	// EnterLatitudeUnits is called when entering the latitudeUnits production.
	EnterLatitudeUnits(c *LatitudeUnitsContext)

	// EnterLongitude is called when entering the longitude production.
	EnterLongitude(c *LongitudeContext)

	// EnterLongitudeUnits is called when entering the longitudeUnits production.
	EnterLongitudeUnits(c *LongitudeUnitsContext)

	// EnterSpeed is called when entering the speed production.
	EnterSpeed(c *SpeedContext)

	// EnterSpeedUnits is called when entering the speedUnits production.
	EnterSpeedUnits(c *SpeedUnitsContext)

	// EnterVolumeflow is called when entering the volumeflow production.
	EnterVolumeflow(c *VolumeflowContext)

	// EnterVolumeflowUnits is called when entering the volumeflowUnits production.
	EnterVolumeflowUnits(c *VolumeflowUnitsContext)

	// EnterEnergy is called when entering the energy production.
	EnterEnergy(c *EnergyContext)

	// EnterEnergyUnits is called when entering the energyUnits production.
	EnterEnergyUnits(c *EnergyUnitsContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitThingContent is called when exiting the thingContent production.
	ExitThingContent(c *ThingContentContext)

	// ExitConfigChanges is called when exiting the configChanges production.
	ExitConfigChanges(c *ConfigChangesContext)

	// ExitConfigTime is called when exiting the configTime production.
	ExitConfigTime(c *ConfigTimeContext)

	// ExitConfigUpper is called when exiting the configUpper production.
	ExitConfigUpper(c *ConfigUpperContext)

	// ExitConfigLower is called when exiting the configLower production.
	ExitConfigLower(c *ConfigLowerContext)

	// ExitUnitTypeOptions is called when exiting the unitTypeOptions production.
	ExitUnitTypeOptions(c *UnitTypeOptionsContext)

	// ExitLogicUnits is called when exiting the logicUnits production.
	ExitLogicUnits(c *LogicUnitsContext)

	// ExitVoltage is called when exiting the voltage production.
	ExitVoltage(c *VoltageContext)

	// ExitVoltagesUnits is called when exiting the voltagesUnits production.
	ExitVoltagesUnits(c *VoltagesUnitsContext)

	// ExitCurrent is called when exiting the current production.
	ExitCurrent(c *CurrentContext)

	// ExitCurrentUnits is called when exiting the currentUnits production.
	ExitCurrentUnits(c *CurrentUnitsContext)

	// ExitResistance is called when exiting the resistance production.
	ExitResistance(c *ResistanceContext)

	// ExitResistanceUnits is called when exiting the resistanceUnits production.
	ExitResistanceUnits(c *ResistanceUnitsContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitPowerUnits is called when exiting the powerUnits production.
	ExitPowerUnits(c *PowerUnitsContext)

	// ExitTemperature is called when exiting the temperature production.
	ExitTemperature(c *TemperatureContext)

	// ExitTemperatureUnits is called when exiting the temperatureUnits production.
	ExitTemperatureUnits(c *TemperatureUnitsContext)

	// ExitLuminosity is called when exiting the luminosity production.
	ExitLuminosity(c *LuminosityContext)

	// ExitLuminosityUnits is called when exiting the luminosityUnits production.
	ExitLuminosityUnits(c *LuminosityUnitsContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)

	// ExitTimeUnits is called when exiting the timeUnits production.
	ExitTimeUnits(c *TimeUnitsContext)

	// ExitMass is called when exiting the mass production.
	ExitMass(c *MassContext)

	// ExitMassUnits is called when exiting the massUnits production.
	ExitMassUnits(c *MassUnitsContext)

	// ExitPressure is called when exiting the pressure production.
	ExitPressure(c *PressureContext)

	// ExitPressureUnits is called when exiting the pressureUnits production.
	ExitPressureUnits(c *PressureUnitsContext)

	// ExitDistance is called when exiting the distance production.
	ExitDistance(c *DistanceContext)

	// ExitDistanceUnits is called when exiting the distanceUnits production.
	ExitDistanceUnits(c *DistanceUnitsContext)

	// ExitAngle is called when exiting the angle production.
	ExitAngle(c *AngleContext)

	// ExitAngleUnits is called when exiting the angleUnits production.
	ExitAngleUnits(c *AngleUnitsContext)

	// ExitVolume is called when exiting the volume production.
	ExitVolume(c *VolumeContext)

	// ExitVolumeUnits is called when exiting the volumeUnits production.
	ExitVolumeUnits(c *VolumeUnitsContext)

	// ExitArea is called when exiting the area production.
	ExitArea(c *AreaContext)

	// ExitAreaUnits is called when exiting the areaUnits production.
	ExitAreaUnits(c *AreaUnitsContext)

	// ExitRain is called when exiting the rain production.
	ExitRain(c *RainContext)

	// ExitRainUnits is called when exiting the rainUnits production.
	ExitRainUnits(c *RainUnitsContext)

	// ExitDensity is called when exiting the density production.
	ExitDensity(c *DensityContext)

	// ExitDensityUnits is called when exiting the densityUnits production.
	ExitDensityUnits(c *DensityUnitsContext)

	// ExitLatitude is called when exiting the latitude production.
	ExitLatitude(c *LatitudeContext)

	// ExitLatitudeUnits is called when exiting the latitudeUnits production.
	ExitLatitudeUnits(c *LatitudeUnitsContext)

	// ExitLongitude is called when exiting the longitude production.
	ExitLongitude(c *LongitudeContext)

	// ExitLongitudeUnits is called when exiting the longitudeUnits production.
	ExitLongitudeUnits(c *LongitudeUnitsContext)

	// ExitSpeed is called when exiting the speed production.
	ExitSpeed(c *SpeedContext)

	// ExitSpeedUnits is called when exiting the speedUnits production.
	ExitSpeedUnits(c *SpeedUnitsContext)

	// ExitVolumeflow is called when exiting the volumeflow production.
	ExitVolumeflow(c *VolumeflowContext)

	// ExitVolumeflowUnits is called when exiting the volumeflowUnits production.
	ExitVolumeflowUnits(c *VolumeflowUnitsContext)

	// ExitEnergy is called when exiting the energy production.
	ExitEnergy(c *EnergyContext)

	// ExitEnergyUnits is called when exiting the energyUnits production.
	ExitEnergyUnits(c *EnergyUnitsContext)
}
