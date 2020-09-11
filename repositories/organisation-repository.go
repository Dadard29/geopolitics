package repositories

import (
	"github.com/Dadard29/geopolitics/models"
	"github.com/arangodb/go-driver"
)

func OrganisationGet(orgKey string) (driver.DocumentMeta, models.OrganisationNode, error) {
	var f models.OrganisationNode

	var o models.OrganisationNode
	meta, err := getDocument(organisationNodesCollectionName, orgKey, &o)
	if err != nil {
		return meta, f, err
	}

	return meta, o, nil
}