package managers

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/geopolitics/repositories"
)

func RelationshipPendingManagerCreate(r models.RelationshipPendingInput) (
	models.RelationshipPendingDto, error) {

	entity := r.ToEntity()

	return repositories.RelationshipPendingCreate(entity)
}


func RelationshipPendingManagerGetAll() ([]models.RelationshipPendingDto, error) {
	return repositories.RelationshipPendingGetAll()
}