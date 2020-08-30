package managers

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

// return all countries with computed scored
func CountryManagerGetAll() (models.GraphScore, error) {
	var f models.GraphScore

	countries, err := repositories.CountryGetAll()
	if err != nil {
		return f, err
	}

	rels, err := repositories.RelationshipGetAll()
	scores, err := models.NewRelationshipSetArray(rels).ToScoreArray()
	if err != nil {
		return f, nil
	}

	return models.GraphScore{
		Nodes: countries,
		Edges: scores,
	}, nil
}

// return countries from a region with computed scores
func CountryManagerGetRegion(region string) (models.GraphScore, error) {
	var f models.GraphScore

	countriesRegion, err := repositories.CountryGetRegion(region)

	if err != nil {
		return models.GraphScore{}, err
	}

	rels, err := repositories.RelationshipGetAll()

	relsRegion := make([]models.RelationshipEntity, 0)
	for _, r := range rels {
		// checkin from
		fromFound := false
		for _, c := range countriesRegion {
			if r.FromId == c.Id {
				fromFound = true
			}
		}

		toFound := false
		for _, c := range countriesRegion {
			if r.ToId == c.Id {
				toFound = true
			}
		}

		if fromFound && toFound {
			relsRegion = append(relsRegion, r)
		}
	}

	scores, err := models.NewRelationshipSetArray(relsRegion).ToScoreArray()
	if err != nil {
		return f, nil
	}



	return models.GraphScore{
		Nodes: countriesRegion,
		Edges: scores,
	}, nil
}
