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
// Params: 			organisationKey
// Body: 			None
func OrganisationGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	var out models.OrganisationNodeDto
	var err error

	orgKey := r.URL.Query().Get("organisationKey")
	if orgKey == "" {
		api.Api.BuildMissingParameter(w)
		return
	}

	out, err = managers.OrganisationGetManager(orgKey)
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "error getting organisation", w)
		return
	}

	api.Api.BuildJsonResponse(true, "organisation retrieved", out, w)
}