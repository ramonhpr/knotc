package model

// Thing represents a thing in knot
type Thing struct {
	Name  string
	Items map[string]*DataItem
}