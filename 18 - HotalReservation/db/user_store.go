package db

import (
	"context"

	"github.com/grayjunzi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	dbName     string
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(userCollection),
	}
}

func (m *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := m.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
