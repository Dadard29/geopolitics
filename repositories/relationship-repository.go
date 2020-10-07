package repositories

import (
	"fmt"
	"github.com/Dadard29/geopolitics/models"
	"github.com/arangodb/go-driver"
)

// get all edges connected to a country
func RelationshipGetFromCountry(countryId string) ([]models.RelationshipEntity, error) {
	query := `FOR r in relationship
		FILTER r._to == @countryId OR r._from == @countryId
		SORT r.date DESC
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

// returns all rels from db
func RelationshipGetAll() ([]models.RelationshipEntity, error) {
	var f []models.RelationshipEntity

	relationshipColl, err := openCollection(relationshipCollectionName)
	if err != nil {
		return f, err
	}

	total, err := relationshipColl.Count(ctx)
	logger.CheckErrLog(err)
	logger.Debug(fmt.Sprintf("fetching %d edges...", total))

	total = 0
	var out []models.RelationshipEntity

	query := `FOR r in relationship RETURN r`
	cursor, err := executeQuery(query, map[string]interface{}{})
	if err != nil {
		return f, err
	}

	defer cursor.Close()

	for {
		var r models.RelationshipEntity
		_, err := cursor.ReadDocument(ctx, &r)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Warning(err.Error())
			continue
		} else {
			out = append(out, r)
			total += 1
		}
	}

	logger.Debug(fmt.Sprintf("fetched %d edges", total))

	return out, nil
}

// return all rels connecting 2 countries
func RelationshipGetDetails(countryAId string, countryBId string) ([]models.RelationshipEntity, error) {
	query := `
		FOR e IN relationship
			FILTER e._from == @countryA || e._from == @countryB
			FILTER e._to == @countryA || e._to == @countryB
			return e`

	bindVars := map[string]interface{}{
		"countryA": countryAId,
		"countryB": countryBId,
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
