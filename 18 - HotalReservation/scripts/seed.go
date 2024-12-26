package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grayjunzi/hotel-reservation/db"
	"github.com/grayjunzi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	hotelStore db.HotelStore
	roomStore  db.RoomStore
	userStore  db.UserStore
	ctx        = context.Background()
)

func seedUser(name, email, password string) {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:    email,
		Name:     name,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = userStore.CreateUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func seedHotel(name string, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rating:   rating,
		Rooms:    []primitive.ObjectID{},
	}

	rooms := []types.Room{
		{
			Size:  "small",
			Price: 99.9,
		},
		{
			Size:  "normal",
			Price: 199.9,
		},
		{
			Size:  "kingsize",
			Price: 299.9,
		},
	}

	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(insertedHotel)
	for _, room := range rooms {
		room.HotelId = insertedHotel.Id
		insertedRoom, err := roomStore.Insert(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(insertedRoom)
	}
}

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
		return
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
}

func main() {
	seedHotel("Bellucia", "France", 3)
	seedHotel("The cozy hotel", "The Nederlands", 5)
	seedUser("james", "james@email.com", "test")
}
