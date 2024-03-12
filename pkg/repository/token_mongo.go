package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"medods"
)

type TokenMongo struct {
	db *mongo.Database
}

func NewTokenMongo(db *mongo.Database) *TokenMongo {
	return &TokenMongo{db: db}
}

func (r *TokenMongo) SetRefreshToken(guid string, refresh []byte) error {
	var user = medods.User{GUID: guid, Refresh: string(refresh)}

	_, err := r.db.Collection(usersTable).InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (r *TokenMongo) CheckUser(guid string) error {
	var result medods.User
	filter := bson.M{"guid": guid}

	err := r.db.Collection(usersTable).FindOne(context.TODO(), filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil
	}

	if err != nil {
		return err
	}
	return errors.New("guid already exist")

}

func (r *TokenMongo) GetUser(refresh string) (medods.User, error) {
	var result medods.User
	filter := bson.M{"refresh": refresh}

	err := r.db.Collection(usersTable).FindOne(context.TODO(), filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return result, err
	}
	if err != nil {
		return result, err
	}
	return result, nil

}

func (r *TokenMongo) UpdateTokens(user medods.User) error {
	filter := bson.M{"guid": user.GUID}
	update := bson.M{"$set": bson.M{"refresh": user.Refresh}}

	_, err := r.db.Collection(usersTable).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
