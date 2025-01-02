package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/api"
	"github.com/grayjunzi/hotel-reservation/db"
	"github.com/grayjunzi/hotel-reservation/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	var (
		userStore    = db.NewMongoUserStore(client)
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			User:    userStore,
			Hotel:   hotelStore,
			Room:    roomStore,
			Booking: bookingStore,
		}
		userHandler    = api.NewUserHandler(userStore)
		hotelHandler   = api.NewHotelHandler(store)
		authHandler    = api.NewAuthHandler(userStore)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		app            = fiber.New(config)
		auth           = app.Group("/api")
		v1             = app.Group("/api/v1", middleware.JWTAuthentication(userStore))
		admin          = v1.Group("/admin", middleware.AdminAuth)
	)

	// Auth
	auth.Post("/auth", authHandler.Authenticate)

	// user handlers
	v1.Get("/users", userHandler.GetUsers)
	v1.Get("/users/:id", userHandler.GetUser)
	v1.Post("/users", userHandler.AddUser)
	v1.Put("/users/:id", userHandler.UpdateUser)
	v1.Delete("/users/:id", userHandler.DeleteUser)

	// hotel handlers
	v1.Get("/hotel", hotelHandler.GetHotels)
	v1.Get("/hotel/:id", hotelHandler.GetHotel)
	v1.Get("/hotel/:id/rooms", hotelHandler.GetRooms)

	// room handlers
	v1.Get("/room", roomHandler.GetRooms)
	v1.Post("/room/:id/book", roomHandler.BookRoom)

	// booking handlers
	admin.Get("/booking", bookingHandler.GetBookings)
	v1.Get("/booking/:id", bookingHandler.GetBooking)
	v1.Get("/booking/:id/cancel", bookingHandler.CancelBooking)

	app.Listen(*listenAddr)
}
