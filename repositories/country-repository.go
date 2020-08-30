package repositories

import (
	"errors"
	"fmt"
	"github.com/Dadard29/geopolitics/models"
	"github.com/arangodb/go-driver"
)

// return all countries
func CountryGetAll() ([]models.CountryDto, error) {
	var f []models.CountryDto

	countryColl, err := openCollection(countryCollectionName)
	if err != nil {
		return f, err
	}

	total, err := countryColl.Count(ctx)
	logger.CheckErrLog(err)
	logger.Debug(fmt.Sprintf("fetching %d documents...", total))

	total = 0
	var out []models.CountryDto

	query := `FOR c in country RETURN c`
	cursor, err := executeQuery(query, map[string]interface{}{})
	if err != nil {
		return f, err
	}

	defer cursor.Close()

	for {
		var country models.CountryEntity
		meta, err := cursor.ReadDocument(ctx, &country)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Warning(err.Error())
			continue
		} else {
			out = append(out, country.ToDto(meta))
			total += 1
		}
	}

	logger.Debug(fmt.Sprintf("fetched %d documents", total))

	return out, nil
}

// return all countries that belongs to the region
func CountryGetRegion(region string) ([]models.CountryDto, error) {
	var regionNode models.RegionNode
	regionNodeMeta, err := getDocument(regionNodesCollectionName, region, &regionNode)
	if err != nil {
		return nil, errors.New("region not found")
	}


	query := `FOR doc IN @@collection
				  FILTER doc._to == @region
				  RETURN DOCUMENT(doc._from)`

	bindVars := map[string]interface{}{
		"@collection": regionEdgesCollectionName,
		"region": regionNodeMeta.ID.String(),
	}

	cursor, err := executeQuery(query, bindVars)
	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	var relList = make([]models.CountryDto, 0)
	for {
		var doc models.CountryEntity
		meta, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			logger.Warning(err.Error())
			continue
		}

		relList = append(relList, doc.ToDto(meta))
	}

	return relList, nil
}

// retrieve one specific country from key
func CountryGet(countryKey string) (driver.DocumentMeta, models.CountryEntity, error) {
	var f models.CountryEntity

	var c models.CountryEntity
	meta, err := getDocument(countryCollectionName, countryKey, &c)
	if err != nil {
		return meta, f, err
	}

	return meta, c, nil
}

// returns ID of country if found
func CountryExists(countryKey string) (string, error) {
	var f string

	var c models.CountryEntity
	meta, err := getDocument(countryCollectionName, countryKey, &c)
	if err != nil {
		return f, err
	}

	return meta.ID.String(), nil
}
