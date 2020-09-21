package repositories

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/arangodb/go-driver"
)

// create pending rel in db
func RelationshipPendingCreate(r models.RelationshipPendingEntity) (
	models.RelationshipPendingDto, error) {

	var f models.RelationshipPendingDto

	meta, err := createDocument(relationshipPendingCollectionName, r)
	if err != nil {
		return f, err
	}

	var out models.RelationshipPendingEntity
	_, err = getDocument(relationshipPendingCollectionName, meta.Key, &out)
	if err != nil {
		return f, err
	}

	return out.ToDto(meta), nil
}

// get specific rel
func RelationshipPendingGet(key string) (models.RelationshipPendingEntity, error) {
	var f models.RelationshipPendingEntity

	col, err := openCollection(relationshipPendingCollectionName)
	if err != nil {
		return f, err
	}

	var out models.RelationshipPendingEntity
	_, err = col.ReadDocument(ctx, key, &out)
	if err != nil {
		return f, err
	}

	return out, nil
}

// delete specific rel
func RelationshipPendingDelete(key string) (models.RelationshipPendingDto, error) {
	var f models.RelationshipPendingDto

	col, err := openCollection(relationshipPendingCollectionName)
	if err != nil {
		return f, err
	}

	var out models.RelationshipPendingEntity
	_, err = col.ReadDocument(ctx, key, &out)
	if err != nil {
		return f, err
	}

	meta, err := col.RemoveDocument(ctx, key)
	if err != nil {
		return f, err
	}

	return out.ToDto(meta), nil
}

// list all rels stored
func RelationshipPendingGetAll() ([]models.RelationshipPendingDto, error) {
	var f []models.RelationshipPendingDto

	query := `FOR r IN @@collection RETURN r`
	bindVars := map[string]interface{}{
		"@collection": "relationship_pending",
	}
	cursor, err := executeQuery(query, bindVars)
	if err != nil {
		return f, err
	}

	defer cursor.Close()

	var out = make([]models.RelationshipPendingDto, 0)

	for {
		var r models.RelationshipPendingEntity
		meta, err := cursor.ReadDocument(ctx, &r)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Warning(err.Error())
			continue
		}

		out = append(out, r.ToDto(meta))
	}

	return out, nil
}
