package models

import (
	"errors"
	"time"
)

var RelationImpactNegativeLow = -1
var RelationImpactNegative = -2
var RelationImpactNegativeHigh = -3

var RelationImpactPositiveLow = 1
var RelationImpactPositive = 2
var RelationImpactPositiveHigh = 3

func checkImpactValue(impact int) bool {
	if impact != RelationImpactNegativeLow &&
		impact != RelationImpactNegative &&
		impact != RelationImpactNegativeHigh &&
		impact != RelationImpactPositiveLow &&
		impact != RelationImpactPositive &&
		impact != RelationImpactPositiveHigh {

		return false
	}

	return true
}

// model given by user
type RelationshipInput struct {
	Subject     string    `json:"subject"`
	ArticleLink string    `json:"article_link"`
	Brief       string    `json:"brief"`
	Sector      string    `json:"sector"`
	Date        time.Time `json:"date"`
	Impact      int       `json:"impact"`
}

// fixme: check sector, check link host
func (rel RelationshipInput) CheckSanity() error {
	if rel.Subject == "" ||
		rel.ArticleLink == "" ||
		rel.Brief == "" ||
		rel.Sector == "" {
		return errors.New("empty parameter")
	}

	if rel.Date.After(time.Now()) {
		return errors.New("wrong date")
	}

	if !checkImpactValue(rel.Impact) {
		return errors.New("wrong impact value, must be between [-3, 3]")
	}

	return nil
}

func (rel RelationshipInput) ToEntity(fromId string, toId string) (RelationshipEntity, error) {
	return RelationshipEntity{
		FromId:      fromId,
		ToId:        toId,
		Subject:     rel.Subject,
		ArticleLink: rel.ArticleLink,
		Brief:       rel.Brief,
		Sector:      rel.Sector,
		Date:        rel.Date,
		Impact:      rel.Impact,
	}, nil
}

// model relationships of db
// event concerning the diplomacy between 2 countries
type RelationshipEntity struct {
	FromId      string    `json:"_from"`
	ToId        string    `json:"_to"`
	Subject     string    `json:"subject"`
	ArticleLink string    `json:"article_link"`
	Brief       string    `json:"brief"`
	Sector      string    `json:"sector"`
	Date        time.Time `json:"date"`
	Impact      int       `json:"impact"`
}
