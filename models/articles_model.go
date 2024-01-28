package models

import (
	"github.com/google/uuid"
)

type Article struct {
	Id      string   `bson:"_id" json:"_id"`
	Name    string   `bson:"name" json:"name"`
	Authors []string `bson:"authors" json:"authors"`
}

type ArticleModelWithoutID struct {
	Name    string   `bson:"name" json:"name"`
	Authors []string `bson:"authors" json:"authors"`
}

func GetArticleDBModel(input *ArticleModelWithoutID) *Article {
	return &Article{
		Id:      uuid.New().String(),
		Name:    input.Name,
		Authors: input.Authors,
	}
}
