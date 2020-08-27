package managers

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

// return all countries with computed scored
func CountryManagerGetAll() (models.CountriesAndRelationships, error) {
	out, err := repositories.CountryGetAll()
	if err != nil {
		return models.CountriesAndRelationships{}, err
	}

	return models.CountriesAndRelationships{
		Nodes: out,
		Edges: make([]models.RelationshipScore, 0),
	}, nil
}

// return countries from a region with computed scores
func CountryManagerGetRegion(region string) (models.CountriesAndRelationships, error) {
	out, err := repositories.CountryGetRegion(region)

	if err != nil {
		return models.CountriesAndRelationships{}, err
	}

	return models.CountriesAndRelationships{
		Nodes: out,
		Edges: make([]models.RelationshipScore, 0),
	}, nil
}
