package models

type CountriesAndRelationships struct {
	Nodes []CountryDto `json:"nodes"`
	Edges []Relationship `json:"edges"`
}
