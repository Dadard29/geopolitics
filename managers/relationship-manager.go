package managers

import (
	"errors"
	"fmt"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
	"time"
)

// returns all edges (scores) connected to a specific country
func RelationshipManagerGetFromCountry(countryKey string) (models.GraphScore, error) {
	var f models.GraphScore

	meta, countryEntity, err := repositories.CountryGet(countryKey)
	if err != nil {
		return f, err
	}

	countryId := meta.ID.String()

	// get the rels array
	rels, err := repositories.RelationshipGetFromCountry(countryId)
	if err != nil {
		return f, err
	}

	// init node array
	nodes := []models.CountryDto{
		countryEntity.ToDto(meta),
	}

	// get all connected nodes
	for _, r := range rels {
		var countryLinkedKey string

		if r.FromId != countryId {
			countryLinkedKey = repositories.KeyFromId(r.FromId)
		} else if r.ToId != countryId {
			countryLinkedKey = repositories.KeyFromId(r.ToId)
		} else {
			return f, errors.New(fmt.Sprintf("relationship loop detected on country %s", countryId))
		}

		// check if duplicate
		found := false
		for _, c := range nodes {
			if c.Key == countryLinkedKey {
				found = true
			}
		}

		if found {
			continue
		}

		// update node array
		meta, countryLinked, err := repositories.CountryGet(countryLinkedKey)
		if err != nil {
			return f, err
		}
		nodes = append(nodes, countryLinked.ToDto(meta))
	}

	// convert relationships to score edges
	scores, err := models.NewRelationshipSetArray(rels).ToScoreArray()
	if err != nil {
		return f, err
	}

	return models.GraphScore{
		Nodes: nodes,
		Edges: scores,
	}, nil

}

// create a relationship between 2 countries, returns a graph with the 2 countries and the created edge
func RelationshipManagerCreate(relInput models.RelationshipInput, fromKey string, toKey string) (models.Graph, error) {
	var f models.Graph

	meta, countryFrom, err := repositories.CountryGet(fromKey)
	if err != nil {
		logger.Warning(fmt.Sprintf("country `fromKey` not found: %s", fromKey))
		return f, err
	}

	countryFromDto := countryFrom.ToDto(meta)
	countryFromId := meta.ID.String()

	meta, countryTo, err := repositories.CountryGet(toKey)
	if err != nil {
		logger.Warning(fmt.Sprintf("country `toKey` not found: %s", toKey))
		return f, err
	}

	countryToDto := countryTo.ToDto(meta)
	countryToId := meta.ID.String()

	entity, err := relInput.ToEntity(countryFromId, countryToId)
	if err != nil {
		return f, err
	}

	rel, err := repositories.RelationshipCreate(entity)
	if err != nil {
		return f, err
	}

	return models.Graph{
		Nodes: []models.CountryDto{
			countryFromDto,
			countryToDto,
		},
		Edges: []models.RelationshipEntity{
			rel,
		},
	}, nil
}


// return all edges connecting 2 countries
func RelationshipManagerDetails(countryKeyA string, countryKeyB string) (models.GraphDetail, error) {
	var f models.GraphDetail

	meta, countryA, err := repositories.CountryGet(countryKeyA)
	if err != nil {
		return f, nil
	}
	countryADto := countryA.ToDto(meta)

	meta, countryB, err := repositories.CountryGet(countryKeyB)
	if err != nil {
		return f, nil
	}
	countryBDto := countryB.ToDto(meta)

	relList, err := repositories.RelationshipGetDetails(countryADto.Id, countryBDto.Id)
	sets := models.NewRelationshipSetArray(relList)
	scores, err := sets.ToScoreArray()
	if err != nil {
		return f, err
	}

	if len(scores) > 1 {
		return f, errors.New("the relationship query fucked up")
	}

	// fallback value (no scored value found in db)
	scoreValue := models.RelationshipScore{
		Country_A_Id:      countryADto.Id,
		Country_B_Id:      countryBDto.Id,
		Score:             0,
		SectorRepartition: nil,
		LastUpdate:        time.Time{},
	}
	if len(scores) != 0 {
		scoreValue = scores[0]
	}

	return models.GraphDetail{
		Nodes:       []models.CountryDto{countryADto, countryBDto},
		EdgeScore:   scoreValue,
		EdgeHistory: relList,
	}, nil
}
