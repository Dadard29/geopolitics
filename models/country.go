package models

import "github.com/arangodb/go-driver"

type Country struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Population int    `json:"population"`
	Coordinate struct {
		Latitude  int `json:"latitude"`
		Longitude int `json:"longitude"`
	} `json:"coordinate"`
	Currencies []string `json:"currencies"`
	Languages  []string `json:"languages"`
	Flag       string   `json:"flag"`
	Rank       int      `json:"rank"`
}

type CountryDto struct {
	Key        string `json:"key"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Population int    `json:"population"`
	Coordinate struct {
		Latitude  int `json:"latitude"`
		Longitude int `json:"longitude"`
	} `json:"coordinate"`
	Currencies []string `json:"currencies"`
	Languages  []string `json:"languages"`
	Flag       string   `json:"flag"`
	Rank       int      `json:"rank"`
}

func (c Country) ToDto(meta driver.DocumentMeta) CountryDto {
	return CountryDto{
		Key:        meta.Key,
		Id:         meta.ID.String(),
		Name:       c.Name,
		Capital:    c.Capital,
		Population: c.Population,
		Coordinate: struct {
			Latitude  int `json:"latitude"`
			Longitude int `json:"longitude"`
		}{
			Latitude:  c.Coordinate.Latitude,
			Longitude: c.Coordinate.Longitude,
		},
		Currencies: c.Currencies,
		Languages:  c.Languages,
		Flag:       c.Flag,
		Rank:       c.Rank,
	}
}
