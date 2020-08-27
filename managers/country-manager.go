package managers

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

func CountryManagerGetAll() (models.CountriesAndRelationships, error) {
	out, err := repositories.CountryGetAll()
	if err != nil {
		return models.CountriesAndRelationships{}, err
	}

	return models.CountriesAndRelationships{
		Nodes: out,
		Edges: make([]models.Relationship, 0),
	}, nil
}

func CountryManagerGetRegion(region string) (models.CountriesAndRelationships, error) {
	out, err := repositories.CountryGetRegion(region)

	if err != nil {
		return models.CountriesAndRelationships{}, err
	}

	return models.CountriesAndRelationships{
		Nodes: out,
		Edges: make([]models.Relationship, 0),
	}, nil
}
