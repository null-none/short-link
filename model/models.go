package model

import "context"

// DbConnector database client abstraction with required methods
type DbConnector interface {
	Insert(ctx *context.Context, data *ShortUrl) error
	FindOne(ctx *context.Context, hashId string) (ShortUrl, error)
}

// ShortUrlReq incoming POST request struct to create short url
type ShortUrlReq struct {
	Url string `validate:"required"`
}

// ShortUrl model of key:value couple representation of unique hash id and URL
type ShortUrl struct {
	Url  string `bson:"Url" validate:"required"`
	Hash string `bson:"Hash" validate:"required"`
}
