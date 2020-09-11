package managers

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

func OrganisationGetManager(organisationKey string) (models.OrganisationNodeDto, error) {

	var f models.OrganisationNodeDto

	meta, entity, err := repositories.OrganisationGet(organisationKey)
	if err != nil {
		return f, err
	}

	return entity.ToDto(meta), nil
}
