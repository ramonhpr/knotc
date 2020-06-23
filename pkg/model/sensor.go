package model

// Sensor is the sensor associated to knot
type Sensor struct {
	Name, Value, TypeUnit, Unit, DefaultValue string
	ID uint
	IsSensor bool
	Configs []Config
}