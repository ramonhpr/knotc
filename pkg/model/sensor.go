package model

// DataItem is the data item that the thing is associated in knot
type DataItem struct {
	Name, Value, TypeUnit, Unit, DefaultValue string
	ID uint
	IsSensor bool
	Configs []Config
}