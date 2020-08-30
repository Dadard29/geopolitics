package models

import "math"

type SectorProportion struct {
	Name       string `json:"name"`
	Proportion float64 `json:"proportion"`
}

type SectorQuantityArray struct {
	Total int
	Array []SectorQuantity
}

type SectorQuantity struct {
	Name string
	Quantity int
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

	return sp
}
