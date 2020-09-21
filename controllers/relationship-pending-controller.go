package controllers

import (
	"encoding/json"
	"github.com/Dadard29/geopolitics/api"
	"github.com/Dadard29/geopolitics/managers"
	"github.com/Dadard29/geopolitics/models"
	"github.com/Dadard29/go-api-utils/auth"
	"io/ioutil"
	"net/http"
	"strconv"
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

	if !rel.CheckSanity() {
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

// GET
// Authorization: 	sub
// Params: 			rel key
// Body: 			None
func RelationshipPendingDelete(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	k := r.URL.Query().Get("key")
	if k == "" {
		api.Api.BuildMissingParameter(w)
		return
	}

	out, err := managers.RelationshipPendingManagerDelete(k)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to delete rel", w)
		return
	}

	api.Api.BuildJsonResponse(true, "rel deleted", out, w)
}

// GET
// Authorization: 	sub
// Params: 			rel key
// Body: 			None
func RelationshipPendingPut(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	k := r.URL.Query().Get("key")
	fromKey := r.URL.Query().Get("fromKey")
	toKey := r.URL.Query().Get("toKey")
	subject := r.URL.Query().Get("subject")
	sector := r.URL.Query().Get("sector")
	impactStr := r.URL.Query().Get("impact")

	if k == "" || fromKey == "" || toKey == "" || subject == "" ||
		sector == "" || impactStr == "" {

		api.Api.BuildMissingParameter(w)
		return
	}

	impact, err := strconv.Atoi(impactStr)
	if err != nil {
		api.Api.BuildMissingParameter(w)
		return
	}

	out, err := managers.RelationshipPendingManagerConfirm(k, fromKey, toKey, subject, sector, impact)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, "failed to confirm the rel", w)
		return
	}

	api.Api.BuildJsonResponse(true, "rel added", out, w)
}
