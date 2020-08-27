package managers

import (
	"errors"
	"fmt"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

// returns all edges (scores) connected to a specific country
// todo: useful ?
func RelationshipManagerGet(countryKey string) (models.GraphScore, error) {
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
			countryLinkedKey = repositories.CountryKeyFromId(r.FromId)
		} else if r.ToId != countryId {
			countryLinkedKey = repositories.CountryKeyFromId(r.ToId)
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
	relSetArray := models.NewRelationshipSetArray(rels)
	var scores = make([]models.RelationshipScore, 0)
	for _, rs := range relSetArray.Sets {
		newScore, err := rs.ToScore()
		if err != nil {
			logger.Warning(err.Error())
			continue
		}
		scores = append(scores, newScore)
	}

	return models.GraphScore{
		Nodes: nodes,
		Edges: scores,
	}, nil

}

// create a relationship between 2 countries
// return a graph with the 2 countries and the created edge
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
