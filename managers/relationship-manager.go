package managers

import (
	"fmt"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

func RelationshipManagerGet(country string) ([]models.Relationship, error) {
	countryId, err := repositories.CountryExists(country)
	if err != nil {
		return nil, err
	}

	return repositories.RelationshipGetFromCountry(countryId)
}

func RelationshipManagerCreate(rel models.RelationshipDto, from string, to string) (models.Relationship, error) {
	var f models.Relationship

	fromId, err := repositories.CountryExists(from)
	if err != nil {
		logger.Warning(fmt.Sprintf("country `from` not found: %s", from))
		return f, err
	}

	toId, err := repositories.CountryExists(to)
	if err != nil {
		logger.Warning(fmt.Sprintf("country `to` not found: %s", to))
		return f, err
	}

	entity, err := models.NewRelationshipFromDto(rel, fromId, toId)
	if err != nil {
		return f, err
	}

	out, err := repositories.RelationshipCreate(entity)
	if err != nil {
		return f, err
	}

	return out, nil
}
