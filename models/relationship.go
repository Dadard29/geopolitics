package models

type Relationship struct {
	Subject     string `json:"subject"`
	ArticleLink string `json:"article_link"`
	Brief       string `json:"brief"`
	Sector      string `json:"sector"`
	Date        struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"date"`
	Impact string `json:"impact"`
}

func NewRelationship(subject string, link string, brief string, ) Relationship {

}