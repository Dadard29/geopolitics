package models

// contain arbitrary edges and nodes
type Graph struct {
	Nodes []CountryDto         `json:"nodes"`
	Edges []RelationshipDto `json:"edges"`
}

// contain one edge (score) per country-pair
type GraphScore struct {
	Nodes []CountryDto        `json:"nodes"`
	Edges []RelationshipScore `json:"edges"`
}

// supposed to contain only 2 countries
type GraphDetail struct {
	Nodes       []CountryDto         `json:"nodes"`
	EdgeScore   RelationshipScore    `json:"edgeScore"`
	EdgeHistory []RelationshipDto `json:"edgeHistory"`
}
