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

type RelationshipInput struct {
	Subject     string    `json:"subject"`
	ArticleLink string    `json:"article_link"`
	Brief       string    `json:"brief"`
	Sector      string    `json:"sector"`
	Date        time.Time `json:"date"`
	Impact      int       `json:"impact"`
}

func (r RelationshipInput) CheckSanity() error {
	if r.Subject == "" ||
		r.ArticleLink == "" ||
		r.Brief == "" ||
		r.Sector == "" {
		return errors.New("empty parameter")
	}

	if r.Date.After(time.Now()) {
		return errors.New("wrong date")
	}

	if !checkImpactValue(r.Impact) {
		return errors.New("wrong impact value, must be between [-3, 3]")
	}

	return nil

}

type Relationship struct {
	FromId      string    `json:"_from"`
	ToId        string    `json:"_to"`
	Subject     string    `json:"subject"`
	ArticleLink string    `json:"article_link"`
	Brief       string    `json:"brief"`
	Sector      string    `json:"sector"`
	Date        time.Time `json:"date"`
	Impact      int       `json:"impact"`
}

func NewRelationshipFromInput(rel RelationshipInput, fromId string, toId string) (Relationship, error) {
	return Relationship{
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
