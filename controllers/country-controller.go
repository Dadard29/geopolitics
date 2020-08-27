package controllers

import (
	"github.com/Dadard29/geopolitics/api"
	"github.com/Dadard29/geopolitics/managers"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/go-api-utils/auth"
	"net/http"
)

// GET
// Authorization: 	sub
// Params: 			region
// Body: 			None
func CountryAllGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	var out models.GraphScore
	var err error

	region := r.URL.Query().Get("region")
	if region == "" {
		out, err = managers.CountryManagerGetAll()
	} else {
		out, err = managers.CountryManagerGetRegion(region)
	}


	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError,
			"error getting all countries", w)
		return
	}

	api.Api.BuildJsonResponse(true, "countries retrieved", out, w)
}
