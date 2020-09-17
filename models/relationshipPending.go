package models

import "github.com/arangodb/go-driver"

// model given by user
type RelationshipPendingInput struct {
	ArticleLink string `json:"article_link"`
	TweetText string `json:"brief"`
	Hashtags []string `json:"hashtags"`
}

func (r RelationshipPendingInput) CheckSanity() bool {
	if r.ArticleLink == "" || r.TweetText == "" {
		return false
	}
	return true
}

func (r RelationshipPendingInput) ToEntity() RelationshipPendingEntity {
	return RelationshipPendingEntity{
		ArticleLink: r.ArticleLink,
		TweetText:   r.TweetText,
		Hashtags:    r.Hashtags,
	}
}

// model in db
type RelationshipPendingEntity struct {
	ArticleLink string `json:"article_link"`
	TweetText string `json:"brief"`
	Hashtags []string `json:"hashtags"`
}

func (r RelationshipPendingEntity) ToDto(meta driver.DocumentMeta) RelationshipPendingDto {

	return RelationshipPendingDto{
		Key:         meta.Key,
		ArticleLink: r.ArticleLink,
		TweetText:   r.TweetText,
		Hashtags:    r.Hashtags,
	}
}

// model output
type RelationshipPendingDto struct {
	Key string `json:"key"`
	ArticleLink string `json:"article_link"`
	TweetText string `json:"brief"`
	Hashtags []string `json:"hashtags"`
}
