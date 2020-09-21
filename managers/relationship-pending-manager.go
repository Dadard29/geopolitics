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

func RelationshipPendingManagerDelete(key string) (models.RelationshipPendingDto, error) {
	return repositories.RelationshipPendingDelete(key)
}

func RelationshipPendingManagerConfirm(key string, fromKey string, toKey string, subject string,
	sector string, impact int) (models.RelationshipEntity, error) {

	var f models.RelationshipEntity

	pe, err := repositories.RelationshipPendingGet(key)
	if err != nil {
		return f, err
	}

	metaFrom, _, err := repositories.CountryGet(fromKey)
	if err != nil {
		return f, err
	}

	metaTo, _, err := repositories.CountryGet(toKey)
	if err != nil {
		return f, err
	}

	relationshipEntity := models.RelationshipEntity{
		FromId:      metaFrom.ID.String(),
		ToId:        metaTo.ID.String(),
		Subject:     subject,
		ArticleLink: pe.ArticleLink,
		Brief:       pe.TweetText,
		Sector:      sector,
		Date:        pe.Date,
		Impact:      impact,
	}

	relEntity, err := repositories.RelationshipCreate(relationshipEntity)
	if err != nil {
		return f, err
	}

	_, err = repositories.RelationshipPendingDelete(key)
	if err != nil {
		logger.Warning("failed to remove pending_relationship after confirming it")
		return f, err
	}

	return relEntity, nil
}
