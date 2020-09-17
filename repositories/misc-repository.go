package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"strings"
)

var logger = log.NewLogger("REPOSITORY", logLevel.DEBUG)

var connector driver.Database
var ctx = context.Background()

var countryCollectionName = "country"
var relationshipCollectionName = "relationship"
var regionNodesCollectionName = "region_nodes"
var regionEdgesCollectionName = "region_edges"
var organisationNodesCollectionName = "organisation_nodes"
var organisationEdgesCollectionName = "organisation_edges"
var relationshipPendingCollectionName = "relationship_pending"

func executeQuery(query string, bindVars map[string]interface{}) (driver.Cursor, error) {
	return connector.Query(ctx, query, bindVars)
}

func openCollection(collection string) (driver.Collection, error) {
	var f driver.Collection

	found, err := connector.CollectionExists(ctx, collection)
	if err != nil {
		logger.Warning(fmt.Sprintf("failed checkin collection %s", collection))
		return f, err
	}
	if !found {
		return f, errors.New(fmt.Sprintf("collection %s not found", collection))
	}

	col, err := connector.Collection(ctx, collection)
	if err != nil {
		logger.Warning(fmt.Sprintf("failed openin collection %s", collection))
		return f, err
	}

	return col, nil
}

func documentExist(collection string, key string) bool {
	col, err := openCollection(collection)
	if err != nil {
		logger.Warning(fmt.Sprintf("failed openin collection %s", collection))
		return false
	}

	b, err := col.DocumentExists(ctx, key)
	if err != nil {
		logger.Warning(fmt.Sprintf("failed checkin document with key %s", key))
		return false
	}

	return b
}

// expect doc to be a pointer
func getDocument(collection string, docKey string, doc interface{}) (driver.DocumentMeta, error) {
	var f driver.DocumentMeta

	col, err := openCollection(collection)
	if err != nil {
		return f, err
	}

	return col.ReadDocument(ctx, docKey, doc)
}

// expect doc to be an object
func createDocument(collection string, doc interface{}) (driver.DocumentMeta, error) {
	var f driver.DocumentMeta

	col, err := openCollection(collection)
	if err != nil {
		return f, err
	}

	meta, err := col.CreateDocument(ctx, doc)
	if err != nil {
		logger.Warning("failed to create doc")
		return f, err
	}
	logger.Debug(fmt.Sprintf("created document with Key %s", meta.Key))
	return meta, nil
}

// return key from id
// 'country/FRA' gives 'FRA'
func KeyFromId(id string) string {
	return strings.Split(id, "/")[1]
}

// init driver
func SetArangoDBConnector(url string, user string, password string, database string) error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{url},
	})
	if err != nil {
		return err
	}

	auth := driver.BasicAuthentication(user, password)

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: auth,
	})

	if err != nil {
		return err
	}

	connector, err = client.Database(ctx, database)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("connected to database %s", database))

	return nil
}
