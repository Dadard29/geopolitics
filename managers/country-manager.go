package managers

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

// return all countries with computed scored
func CountryManagerGetAll() (models.GraphScore, error) {
	countries, err := repositories.CountryGetAll()
	if err != nil {
		return models.GraphScore{}, err
	}

	//rels, err := repositories.RelationshipGetAll()


	return models.GraphScore{
		Nodes: countries,
		Edges: make([]models.RelationshipScore, 0),
	}, nil
}

// return countries from a region with computed scores
func CountryManagerGetRegion(region string) (models.GraphScore, error) {
	countries, err := repositories.CountryGetRegion(region)

	if err != nil {
		return models.GraphScore{}, err
	}

	return models.GraphScore{
		Nodes: countries,
		Edges: make([]models.RelationshipScore, 0),
	}, nil
}
