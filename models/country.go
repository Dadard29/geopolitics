package models

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
}


