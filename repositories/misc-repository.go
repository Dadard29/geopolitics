package repositories

import (
	"context"
	"github.com/Dadard29/go-api-utils/log"
	"github.com/Dadard29/go-api-utils/log/logLevel"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var logger = log.NewLogger("CONTROLLER", logLevel.DEBUG)

var connector driver.Database

func SetArangoDBConnector(url string, user string, password string, database string) error {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints:          []string{url},
	})
	if err != nil {
		return err
	}


	auth := driver.BasicAuthentication(user, password)

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:                   conn,
		Authentication:               auth,
	})

	if err != nil {
		return err
	}

	ctx := context.Background()
	connector, err = client.Database(ctx, database)
	if err != nil {
		return err
	}


	return nil
}