package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"medods"
)

type Token interface {
	SetRefreshToken(guid string, refresh []byte) error
	CheckUser(guid string) error
	GetUser(refresh string) (medods.User, error)
	UpdateTokens(user medods.User) error
}

type Repository struct {
	Token
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Token: NewTokenMongo(db),
	}
}
