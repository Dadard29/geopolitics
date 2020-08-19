package repositories

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/arangodb/go-driver"
	"strings"
)

func CountryKeyFromId(id string) string {
	return strings.Split(id, "/")[1]
}

func CountryGet(countryKey string) (driver.DocumentMeta, models.Country, error) {
	var f models.Country

	var c models.Country
	meta, err := getDocument(countryCollection, countryKey, &c)
	if err != nil {
		return meta, f, err
	}

	return meta, c, nil
}

// returns ID of country if found
func CountryExists(countryKey string) (string, error) {
	var f string

	var c models.Country
	meta, err := getDocument(countryCollection, countryKey, &c)
	if err != nil {
		return f, err
	}

	return meta.ID.String(), nil
}
