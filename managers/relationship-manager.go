package managers

import (
	"errors"
	"fmt"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
	"time"
)

// create a relationship between 2 countries, returns a graph with the 2 countries and the created edge
func RelationshipManagerCreate(relInput models.RelationshipInput, fromKey string, toKey string) (models.Graph, error) {
	var f models.Graph

	meta, countryFrom, err := repositories.CountryGet(fromKey)
	if err != nil {
		logger.Warning(fmt.Sprintf("country `fromKey` not found: %s", fromKey))
		return f, err
	}

	orgsFrom, err := repositories.CountryOrganisations(meta.ID.String())
	if err != nil {
		return f, err
	}
	countryFromDto := countryFrom.ToDto(meta, orgsFrom)
	countryFromId := meta.ID.String()

	meta, countryTo, err := repositories.CountryGet(toKey)
	if err != nil {
		logger.Warning(fmt.Sprintf("country `toKey` not found: %s", toKey))
		return f, err
	}

	orgsTo, err := repositories.CountryOrganisations(meta.ID.String())
	if err != nil {
		return f, err
	}
	countryToDto := countryTo.ToDto(meta, orgsTo)
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
		Edges: []models.RelationshipDto{
			rel.ToDto(),
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
	countryAOrgs, err := repositories.CountryOrganisations(meta.ID.String())
	if err != nil {
		return f, err
	}
	countryADto := countryA.ToDto(meta, countryAOrgs)

	meta, countryB, err := repositories.CountryGet(countryKeyB)
	if err != nil {
		return f, nil
	}
	countryBOrgs, err := repositories.CountryOrganisations(meta.ID.String())
	if err != nil {
		return f, err
	}
	countryBDto := countryB.ToDto(meta, countryBOrgs)

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
		SectorRepartition: make([]models.SectorProportion, 0),
		LastUpdate:        time.Time{},
	}
	if len(scores) != 0 {
		scoreValue = scores[0]
	}

	var relListDto []models.RelationshipDto
	for _, r := range relList {
		relListDto = append(relListDto, r.ToDto())
	}

	return models.GraphDetail{
		Nodes:       []models.CountryDto{countryADto, countryBDto},
		EdgeScore:   scoreValue,
		EdgeHistory: relListDto,
	}, nil
}
