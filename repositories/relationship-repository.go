package repositories

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/arangodb/go-driver"
)

// get all edges connected to a country
func RelationshipGetFromCountry(countryId string) ([]models.RelationshipEntity, error) {
	query := `FOR r in relationship
		FILTER r._to == @countryId OR r._from == @countryId
		return r`

	bindVars := map[string]interface{}{
		"countryId": countryId,
	}

	cursor, err := executeQuery(query, bindVars)
	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	var relList = make([]models.RelationshipEntity, 0)
	for {
		var doc models.RelationshipEntity
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Warning(err.Error())
			continue
		}

		relList = append(relList, doc)
	}

	return relList, nil
}


// create rel in db
func RelationshipCreate(rel models.RelationshipEntity) (models.RelationshipEntity, error) {
	var f models.RelationshipEntity

	meta, err := createDocument(relationshipCollectionName, rel)
	if err != nil {
		return f, err
	}

	var out models.RelationshipEntity
	_, err = getDocument(relationshipCollectionName, meta.Key, &out)
	if err != nil {
		return f, err
	}

	return out, nil
}



func RelationshipGetAll() ([]models.RelationshipEntity, error) {
	// todo
	return nil, nil
}
