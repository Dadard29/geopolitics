package models

import (
	"errors"
	"math"
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

// model rels of db
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

// recap about the global relations between 2 countries
// computed from relationships analysis
type RelationshipScore struct {
	Country_A_Id string `json:"_from"`
	Country_B_Id string `json:"_to"`
	Score        int    `json:"score"`
	SectorRepartition []SectorProportion `json:"sectorRepartition"`
	LastUpdate   time.Time
}


type SectorProportion struct {
	Name       string
	Proportion float64
}

type SectorQuantityArray struct {
	Total int
	Array []SectorQuantity
}

type SectorQuantity struct {
	Name string
	Quantity int
}

func (a SectorQuantityArray) increment(name string) {
	for _, q := range a.Array {
		if q.Name == name {
			q.Quantity += 1
			a.Total += 1
			return
		}
	}

	// sector not found in array, adding it
	a.Array = append(a.Array, SectorQuantity{
		Name:     name,
		Quantity: 1,
	})
	a.Total += 1
}

func (a SectorQuantityArray) toProportion() []SectorProportion {
	var sp = make([]SectorProportion, 0)
	for _, s := range a.Array {
		p := float64(s.Quantity * 100 / a.Total)

		pRounded := math.Round(p)

		sp = append(sp, SectorProportion{
			Name:       s.Name,
			Proportion: pRounded,
		})
	}
}

func ScoreFromRelationships(rels []RelationshipEntity) (RelationshipScore, error) {
	// assuming that only 2 countries are concerned
	var f RelationshipScore

	var out RelationshipScore
	out.Country_A_Id = rels[0].FromId
	out.Country_B_Id = rels[0].ToId

	var sectors SectorQuantityArray

	for _, r := range rels {
		// checking `from` is concerned
		if r.FromId != out.Country_A_Id && r.FromId != out.Country_B_Id {
			return f, errors.New("a third country interfered in score computation")
		}

		// checking `to` is concerned
		if r.ToId != out.Country_A_Id && r.ToId != out.Country_B_Id {
			return f, errors.New("a third country interfered in score computation")
		}

		// add the sector
		sectors.increment(r.Sector)

		// update time
		if r.Date.After(out.LastUpdate) {
			out.LastUpdate = r.Date
		}

		// update score
		// todo
		out.Score += r.Impact
	}

	// compute sector repartition
	out.SectorRepartition = sectors.toProportion()

	return out, nil


}
