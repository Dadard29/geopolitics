package controllers

import (
	"encoding/json"
	"github.com/Dadard29/geopolitics/api"
	"github.com/Dadard29/geopolitics/managers"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/go-api-utils/auth"
	"io/ioutil"
	"net/http"
)

// GET
// Authorization: 	sub
// Params: 			None
// Body: 			None
func RelationshipPendingPost(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid body", w)
		return
	}

	var rel models.RelationshipPendingInput
	err = json.Unmarshal(body, &rel)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusBadRequest, "invalid json body", w)
		return
	}

	if !rel.CheckSanity(){
		api.Api.BuildMissingParameter(w)
		return
	}


	out, err := managers.RelationshipPendingManagerCreate(rel)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "error creating pending request", w)
		return
	}

	api.Api.BuildJsonResponse(true, "pending relationship created", out, w)
}


// GET
// Authorization: 	sub
// Params: 			None
// Body: 			None
func RelationshipPendingGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	out, err := managers.RelationshipPendingManagerGetAll()
	if err != nil {
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "error getting all pending relationships", w)
		return
	}

	api.Api.BuildJsonResponse(true, "pending relationsips retrieved", out, w)
}
