package managers

import (
	"errors"
	"fmt"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

// returns all edges connected to a specific country
// todo
// /!\ NOT USED BY GUI AS OF 27/08 /!\
func RelationshipManagerGet(countryKey string) (models.CountriesAndRelationships, error) {
	var f models.CountriesAndRelationships

	meta, countryEntity, err := repositories.CountryGet(countryKey)
	if err != nil {
		return f, err
	}

	countryId := meta.ID.String()

	// init node array
	nodes := []models.CountryDto{
		countryEntity.ToDto(meta),
	}

	// get the rels array
	rels, err := repositories.RelationshipGetFromCountry(countryId)
	if err != nil {
		return f, err
	}

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

	return models.CountriesAndRelationships{
		Nodes: nodes,
		Edges: rels,
	}, nil

}

// create a relationship between 2 countries
func RelationshipManagerCreate(relInput models.RelationshipInput, fromKey string, toKey string) (models.Relationship, error) {
	var f models.CountriesAndRelationships

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

	return models.CountriesAndRelationships{
		Nodes: []models.CountryDto{
			countryFromDto,
			countryToDto,
		},
		Edges: []models.RelationshipScore{
			rel,
		},
	}, nil
}
