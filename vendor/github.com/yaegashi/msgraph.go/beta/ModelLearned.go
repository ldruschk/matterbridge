// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// LearnedLocationModel undocumented
type LearnedLocationModel struct {
	// Object is the base model of LearnedLocationModel
	Object
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// Coordinates undocumented
	Coordinates *OutlookGeoCoordinates `json:"coordinates,omitempty"`
}
