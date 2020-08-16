package controllers

import (
	"github.com/Dadard29/geopolitics/api"
	"github.com/Dadard29/go-api-utils/auth"
	"net/http"
)

// POST
// Authorization: 	token
// Params: 			None
// Body: 			None
func RelationshipPost(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}


}

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None
func RelationshipGet(w http.ResponseWriter, r *http.Request) {
	accessToken := auth.ParseApiKey(r, accessTokenKey, true)
	if !checkToken(accessToken, w) {
		return
	}

	api.Api.BuildJsonResponse(true, "ok", nil, w)
}
