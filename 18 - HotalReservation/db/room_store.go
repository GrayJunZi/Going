package db

import (
	"context"

	"github.com/grayjunzi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	Insert(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]*types.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection

	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("rooms"),
		HotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	resp, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var rooms []*types.Room
	if err := resp.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *MongoRoomStore) Insert(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := s.collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.Id = resp.InsertedID.(primitive.ObjectID)

	s.HotelStore.Update(ctx, room.HotelId, &types.UpdateHotelParams{
		RoomId: room.Id,
	})
	return room, nil
}
