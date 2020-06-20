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

// EnterCurrent is called when production current is entered.
func (s *BaseKnotListener) EnterCurrent(ctx *CurrentContext) {}

// ExitCurrent is called when production current is exited.
func (s *BaseKnotListener) ExitCurrent(ctx *CurrentContext) {}

// EnterCurrentUnits is called when production currentUnits is entered.
func (s *BaseKnotListener) EnterCurrentUnits(ctx *CurrentUnitsContext) {}

// ExitCurrentUnits is called when production currentUnits is exited.
func (s *BaseKnotListener) ExitCurrentUnits(ctx *CurrentUnitsContext) {}

// EnterResistance is called when production resistance is entered.
func (s *BaseKnotListener) EnterResistance(ctx *ResistanceContext) {}

// ExitResistance is called when production resistance is exited.
func (s *BaseKnotListener) ExitResistance(ctx *ResistanceContext) {}

// EnterResistanceUnits is called when production resistanceUnits is entered.
func (s *BaseKnotListener) EnterResistanceUnits(ctx *ResistanceUnitsContext) {}

// ExitResistanceUnits is called when production resistanceUnits is exited.
func (s *BaseKnotListener) ExitResistanceUnits(ctx *ResistanceUnitsContext) {}

// EnterPower is called when production power is entered.
func (s *BaseKnotListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BaseKnotListener) ExitPower(ctx *PowerContext) {}

// EnterPowerUnits is called when production powerUnits is entered.
func (s *BaseKnotListener) EnterPowerUnits(ctx *PowerUnitsContext) {}

// ExitPowerUnits is called when production powerUnits is exited.
func (s *BaseKnotListener) ExitPowerUnits(ctx *PowerUnitsContext) {}

// EnterTemperature is called when production temperature is entered.
func (s *BaseKnotListener) EnterTemperature(ctx *TemperatureContext) {}

// ExitTemperature is called when production temperature is exited.
func (s *BaseKnotListener) ExitTemperature(ctx *TemperatureContext) {}

// EnterTemperatureUnits is called when production temperatureUnits is entered.
func (s *BaseKnotListener) EnterTemperatureUnits(ctx *TemperatureUnitsContext) {}

// ExitTemperatureUnits is called when production temperatureUnits is exited.
func (s *BaseKnotListener) ExitTemperatureUnits(ctx *TemperatureUnitsContext) {}

// EnterLuminosity is called when production luminosity is entered.
func (s *BaseKnotListener) EnterLuminosity(ctx *LuminosityContext) {}

// ExitLuminosity is called when production luminosity is exited.
func (s *BaseKnotListener) ExitLuminosity(ctx *LuminosityContext) {}

// EnterLuminosityUnits is called when production luminosityUnits is entered.
func (s *BaseKnotListener) EnterLuminosityUnits(ctx *LuminosityUnitsContext) {}

// ExitLuminosityUnits is called when production luminosityUnits is exited.
func (s *BaseKnotListener) ExitLuminosityUnits(ctx *LuminosityUnitsContext) {}

// EnterTime is called when production time is entered.
func (s *BaseKnotListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseKnotListener) ExitTime(ctx *TimeContext) {}

// EnterTimeUnits is called when production timeUnits is entered.
func (s *BaseKnotListener) EnterTimeUnits(ctx *TimeUnitsContext) {}

// ExitTimeUnits is called when production timeUnits is exited.
func (s *BaseKnotListener) ExitTimeUnits(ctx *TimeUnitsContext) {}

// EnterMass is called when production mass is entered.
func (s *BaseKnotListener) EnterMass(ctx *MassContext) {}

// ExitMass is called when production mass is exited.
func (s *BaseKnotListener) ExitMass(ctx *MassContext) {}

// EnterMassUnits is called when production massUnits is entered.
func (s *BaseKnotListener) EnterMassUnits(ctx *MassUnitsContext) {}

// ExitMassUnits is called when production massUnits is exited.
func (s *BaseKnotListener) ExitMassUnits(ctx *MassUnitsContext) {}

// EnterPressure is called when production pressure is entered.
func (s *BaseKnotListener) EnterPressure(ctx *PressureContext) {}

// ExitPressure is called when production pressure is exited.
func (s *BaseKnotListener) ExitPressure(ctx *PressureContext) {}

// EnterPressureUnits is called when production pressureUnits is entered.
func (s *BaseKnotListener) EnterPressureUnits(ctx *PressureUnitsContext) {}

// ExitPressureUnits is called when production pressureUnits is exited.
func (s *BaseKnotListener) ExitPressureUnits(ctx *PressureUnitsContext) {}

// EnterDistance is called when production distance is entered.
func (s *BaseKnotListener) EnterDistance(ctx *DistanceContext) {}

// ExitDistance is called when production distance is exited.
func (s *BaseKnotListener) ExitDistance(ctx *DistanceContext) {}

// EnterDistanceUnits is called when production distanceUnits is entered.
func (s *BaseKnotListener) EnterDistanceUnits(ctx *DistanceUnitsContext) {}

// ExitDistanceUnits is called when production distanceUnits is exited.
func (s *BaseKnotListener) ExitDistanceUnits(ctx *DistanceUnitsContext) {}

// EnterAngle is called when production angle is entered.
func (s *BaseKnotListener) EnterAngle(ctx *AngleContext) {}

// ExitAngle is called when production angle is exited.
func (s *BaseKnotListener) ExitAngle(ctx *AngleContext) {}

// EnterAngleUnits is called when production angleUnits is entered.
func (s *BaseKnotListener) EnterAngleUnits(ctx *AngleUnitsContext) {}

// ExitAngleUnits is called when production angleUnits is exited.
func (s *BaseKnotListener) ExitAngleUnits(ctx *AngleUnitsContext) {}

// EnterVolume is called when production volume is entered.
func (s *BaseKnotListener) EnterVolume(ctx *VolumeContext) {}

// ExitVolume is called when production volume is exited.
func (s *BaseKnotListener) ExitVolume(ctx *VolumeContext) {}

// EnterVolumeUnits is called when production volumeUnits is entered.
func (s *BaseKnotListener) EnterVolumeUnits(ctx *VolumeUnitsContext) {}

// ExitVolumeUnits is called when production volumeUnits is exited.
func (s *BaseKnotListener) ExitVolumeUnits(ctx *VolumeUnitsContext) {}

// EnterArea is called when production area is entered.
func (s *BaseKnotListener) EnterArea(ctx *AreaContext) {}

// ExitArea is called when production area is exited.
func (s *BaseKnotListener) ExitArea(ctx *AreaContext) {}

// EnterAreaUnits is called when production areaUnits is entered.
func (s *BaseKnotListener) EnterAreaUnits(ctx *AreaUnitsContext) {}

// ExitAreaUnits is called when production areaUnits is exited.
func (s *BaseKnotListener) ExitAreaUnits(ctx *AreaUnitsContext) {}

// EnterRain is called when production rain is entered.
func (s *BaseKnotListener) EnterRain(ctx *RainContext) {}

// ExitRain is called when production rain is exited.
func (s *BaseKnotListener) ExitRain(ctx *RainContext) {}

// EnterRainUnits is called when production rainUnits is entered.
func (s *BaseKnotListener) EnterRainUnits(ctx *RainUnitsContext) {}

// ExitRainUnits is called when production rainUnits is exited.
func (s *BaseKnotListener) ExitRainUnits(ctx *RainUnitsContext) {}

// EnterDensity is called when production density is entered.
func (s *BaseKnotListener) EnterDensity(ctx *DensityContext) {}

// ExitDensity is called when production density is exited.
func (s *BaseKnotListener) ExitDensity(ctx *DensityContext) {}

// EnterDensityUnits is called when production densityUnits is entered.
func (s *BaseKnotListener) EnterDensityUnits(ctx *DensityUnitsContext) {}

// ExitDensityUnits is called when production densityUnits is exited.
func (s *BaseKnotListener) ExitDensityUnits(ctx *DensityUnitsContext) {}

// EnterLatitude is called when production latitude is entered.
func (s *BaseKnotListener) EnterLatitude(ctx *LatitudeContext) {}

// ExitLatitude is called when production latitude is exited.
func (s *BaseKnotListener) ExitLatitude(ctx *LatitudeContext) {}

// EnterLatitudeUnits is called when production latitudeUnits is entered.
func (s *BaseKnotListener) EnterLatitudeUnits(ctx *LatitudeUnitsContext) {}

// ExitLatitudeUnits is called when production latitudeUnits is exited.
func (s *BaseKnotListener) ExitLatitudeUnits(ctx *LatitudeUnitsContext) {}

// EnterLongitude is called when production longitude is entered.
func (s *BaseKnotListener) EnterLongitude(ctx *LongitudeContext) {}

// ExitLongitude is called when production longitude is exited.
func (s *BaseKnotListener) ExitLongitude(ctx *LongitudeContext) {}

// EnterLongitudeUnits is called when production longitudeUnits is entered.
func (s *BaseKnotListener) EnterLongitudeUnits(ctx *LongitudeUnitsContext) {}

// ExitLongitudeUnits is called when production longitudeUnits is exited.
func (s *BaseKnotListener) ExitLongitudeUnits(ctx *LongitudeUnitsContext) {}

// EnterSpeed is called when production speed is entered.
func (s *BaseKnotListener) EnterSpeed(ctx *SpeedContext) {}

// ExitSpeed is called when production speed is exited.
func (s *BaseKnotListener) ExitSpeed(ctx *SpeedContext) {}

// EnterSpeedUnits is called when production speedUnits is entered.
func (s *BaseKnotListener) EnterSpeedUnits(ctx *SpeedUnitsContext) {}

// ExitSpeedUnits is called when production speedUnits is exited.
func (s *BaseKnotListener) ExitSpeedUnits(ctx *SpeedUnitsContext) {}

// EnterVolumeflow is called when production volumeflow is entered.
func (s *BaseKnotListener) EnterVolumeflow(ctx *VolumeflowContext) {}

// ExitVolumeflow is called when production volumeflow is exited.
func (s *BaseKnotListener) ExitVolumeflow(ctx *VolumeflowContext) {}

// EnterVolumeflowUnits is called when production volumeflowUnits is entered.
func (s *BaseKnotListener) EnterVolumeflowUnits(ctx *VolumeflowUnitsContext) {}

// ExitVolumeflowUnits is called when production volumeflowUnits is exited.
func (s *BaseKnotListener) ExitVolumeflowUnits(ctx *VolumeflowUnitsContext) {}

// EnterEnergy is called when production energy is entered.
func (s *BaseKnotListener) EnterEnergy(ctx *EnergyContext) {}

// ExitEnergy is called when production energy is exited.
func (s *BaseKnotListener) ExitEnergy(ctx *EnergyContext) {}

// EnterEnergyUnits is called when production energyUnits is entered.
func (s *BaseKnotListener) EnterEnergyUnits(ctx *EnergyUnitsContext) {}

// ExitEnergyUnits is called when production energyUnits is exited.
func (s *BaseKnotListener) ExitEnergyUnits(ctx *EnergyUnitsContext) {}
