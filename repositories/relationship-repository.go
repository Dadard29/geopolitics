package repositories

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/arangodb/go-driver"
)

func RelationshipGetFromCountry(countryId string) ([]models.Relationship, error) {
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

	var relList = make([]models.Relationship, 0)
	for {
		var doc models.Relationship
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Warning(err.Error())
		}

		relList = append(relList, doc)
	}

	return relList, nil
}

func RelationshipCreate(rel models.Relationship) (models.Relationship, error) {
	var f models.Relationship

	meta, err := createDocument(relationshipCollection, rel)
	if err != nil {
		return f, err
	}

	var out models.Relationship
	_, err = getDocument(relationshipCollection, meta.Key, &out)
	if err != nil {
		return f, err
	}

	return out, nil
}
