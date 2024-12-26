package db

import (
	"context"

	"github.com/grayjunzi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper
	GetUsers(context.Context) ([]*types.User, error)
	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUserById(context.Context, string) (*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	UpdateUser(context.Context, string, *types.UpdateUserParams) error
	DeleteUser(context.Context, string) error
}

type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(userCollection),
	}
}

func (m *MongoUserStore) Drop(ctx context.Context) error {
	return m.collection.Drop(ctx)
}

func (m *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*types.User
	if err = cur.All(ctx, &users); err != nil {
		return []*types.User{}, err
	}
	return users, nil
}

func (m *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	if err := m.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
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

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := m.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUserStore) UpdateUser(ctx context.Context, id string, values *types.UpdateUserParams) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}
	update := bson.M{"$set": values}
	_, err = m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
