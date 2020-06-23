package model

// ConfigType is the possibles configurations to a config 
type ConfigType int

const (
	// KnotChanges is a configuration when a sensor sends a data in every changes
	KnotChanges = iota
	// KnotTime is a configuration when a sensor sends a data in a specific time interval
	KnotTime
	// KnotUpper is a configuration when a sensor sends a data in a upper threshold 
	KnotUpper
	// KnotLower is a configuration when a sensor sends a data in a lower threshold
	KnotLower
)

func (k ConfigType) String() string {
	return [...]string{"CHANGE", "TIME", "UPPER_THRESHOLD", "LOWER_THRESHOLD"}[k]
}

// Config is the config associated to knot
type Config struct {
	Type ConfigType
	Value string
}