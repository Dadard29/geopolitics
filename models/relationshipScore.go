package models

import (
	"errors"
	"time"
)

type RelationshipSet struct {
	Country_A_Id  string               `json:"_from"`
	Country_B_Id  string               `json:"_to"`
	Relationships []RelationshipEntity `json:"relationships"`
}

type RelationshipSetArray struct {
	Sets []RelationshipSet
}

func (setArray RelationshipSetArray) ToScoreArray() ([]RelationshipScore, error) {
	var rs = make([]RelationshipScore, 0)
	for _, s := range setArray.Sets {
		score, err := s.ToScore()
		if err != nil {
			return nil, err
		}
		rs = append(rs, score)
	}

	return rs, nil
}

func NewRelationshipSetArray(rels []RelationshipEntity) RelationshipSetArray {
	var array = make([]RelationshipSet, 0)

	for _, r := range rels {
		// check if a set exists already
		exists := false
		for i, s := range array {
			if 	(s.Country_A_Id == r.FromId && s.Country_B_Id == r.ToId) ||
				(s.Country_A_Id == r.ToId && s.Country_B_Id == r.FromId) {
				array[i].Relationships = append(s.Relationships, r)
				exists = true
			}
		}

		if exists {
			continue
		}

		// else add a new set and init it
		array = append(array, RelationshipSet{
			Country_A_Id:  r.FromId,
			Country_B_Id:  r.ToId,
			Relationships: []RelationshipEntity{r},
		})
	}

	return RelationshipSetArray{Sets:array}
}

func (set RelationshipSet) ToScore() (RelationshipScore, error) {
	// assuming that only 2 countries are concerned
	var f RelationshipScore

	var out RelationshipScore
	out.Country_A_Id = set.Country_A_Id
	out.Country_B_Id = set.Country_B_Id

	var sectors []SectorQuantity

	for _, r := range set.Relationships {
		// checking `from` is concerned
		if r.FromId != out.Country_A_Id && r.FromId != out.Country_B_Id {
			return f, errors.New("a third country interfered in score computation")
		}

		// checking `to` is concerned
		if r.ToId != out.Country_A_Id && r.ToId != out.Country_B_Id {
			return f, errors.New("a third country interfered in score computation")
		}

		// add the sector
		found := false
		for i, q := range sectors {
			if q.Name == r.Sector {
				sectors[i].Quantity += 1
				found = true
			}
		}

		// sector not found in array, adding it
		if found == false {
			sectors = append(sectors, SectorQuantity{
				Name:     r.Sector,
				Quantity: 1,
			})
		}

		// update time
		if r.Date.After(out.LastUpdate) {
			out.LastUpdate = r.Date
		}

		// update score
		// todo: test if indicator relevant
		out.Score += r.Impact
	}

	// compute sector repartition
	sectorsArray := SectorQuantityArray{
		Total: len(set.Relationships),
		Array: sectors,
	}
	out.SectorRepartition = sectorsArray.toProportion()

	return out, nil
}

// recap about the global relations between 2 countries
// computed from relationships analysis
type RelationshipScore struct {
	Country_A_Id string `json:"_from"`
	Country_B_Id string `json:"_to"`
	Score        int    `json:"score"`
	SectorRepartition []SectorProportion `json:"sectorRepartition"`
	LastUpdate   time.Time `json:"lastUpdate"`
}