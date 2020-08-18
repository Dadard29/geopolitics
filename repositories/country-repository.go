package repositories

import (
	"github.com/Dadard29/geopolitics/models"
)

// returns ID of country if found
func CountryExists(country string) (string, error) {
	var f string

	var c models.Country
	meta, err := getDocument(countryCollection, country, &c)
	if err != nil {
		return f, err
	}

	return meta.ID.String(), nil
}

