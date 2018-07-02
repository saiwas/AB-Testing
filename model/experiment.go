package model

// Experiment table
type Experiment struct {
	ID          uint64
	Name        string
	Slug        string
	Percentage  uint64
	Status      uint64
	Tags        []uint64
	Segments    map[string]string
	Description string
	StartAt     timestamp
	EndAt       timestamp
	TakeOverID  uint64
	Variants    []*Variant
}
