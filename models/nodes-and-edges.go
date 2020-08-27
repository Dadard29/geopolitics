package models

// contain one edge (score) per country-pair
type CountriesAndRelationships struct {
	Nodes []CountryDto `json:"nodes"`
	Edges []RelationshipScore `json:"edges"`
}

// supposed to contain only 2 countries
type CountriesAndRelationshipsDetails struct {
	Nodes []CountryDto               `json:"nodes"`
	EdgeScore RelationshipScore      `json:"edgeScore"`
	EdgeHistory []RelationshipEntity `json:"edgeHistory"`
}
