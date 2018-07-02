package model

// Variant table
type Variant struct {
	ID                  uint64
	Name                string
	Slug                string
	PercentageOfTest    uint64
	PercentageOfTraffic uint64
	Body                string
}
