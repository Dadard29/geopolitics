package main

import (
	"github.com/Dadard29/geopolitics/api"
	"github.com/Dadard29/geopolitics/controllers"
	"github.com/Dadard29/geopolitics/repositories"
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/service"
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"net/http"
)

var routes = service.RouteMapping{
	"/relationships": service.Route{
		Description:   "manage edges between country nodes",
		MethodMapping: service.MethodMapping{
			http.MethodPost: controllers.RelationshipPost,
			http.MethodGet: controllers.RelationshipGet,
		},
	},
	"/relationships/all": service.Route{
		Description:   "manage all edges at once",
		MethodMapping: service.MethodMapping{
			http.MethodGet: controllers.RelationshipAllGet,
		},
	},
}

func main() {
	var err error
	api.Api = API.NewAPI("Geopolitics",
		"config/config.json", routes, true)

	// init the connectors
	controllers.Sc = subChecker.NewSubChecker(api.Api.Config.GetEnv("HOST_SUB"))

	dbConfig, err := api.Api.Config.GetSubcategoryFromFile("api", "db")
	api.Api.Logger.CheckErrFatal(err)
	err = repositories.SetArangoDBConnector(dbConfig["url"], dbConfig["user"],
		api.Api.Config.GetEnv(dbConfig["passwordKey"]), dbConfig["database"])
	api.Api.Logger.CheckErrFatal(err)

	api.Api.Service.Start()
	api.Api.Service.Stop()
}
