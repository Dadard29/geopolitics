package models

import "github.com/arangodb/go-driver"

type OrganisationNode struct {
	Name          string `json:"name"`
	Formation     string `json:"formation"`
	Description   string `json:"description"`
	Documentation string `json:"documentation"`
	Logo          string `json:"logo"`
}

type OrganisationNodeDto struct {
	Id            string `json:"id"`
	Key           string `json:"key"`
	Name          string `json:"name"`
	Formation     string `json:"formation"`
	Description   string `json:"description"`
	Documentation string `json:"documentation"`
	Logo          string `json:"logo"`
}

func (o OrganisationNode) ToDto(meta driver.DocumentMeta) OrganisationNodeDto {
	return OrganisationNodeDto{
		Id:            meta.ID.String(),
		Key:           meta.Key,
		Name:          o.Name,
		Formation:     o.Formation,
		Description:   o.Description,
		Documentation: o.Documentation,
		Logo:          o.Logo,
	}
}

// only metadata
type OrganisationEdge struct {
}
