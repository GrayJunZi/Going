package db

const (
	DBURI        = "mongodb://localhost:27017"
	DBNAME       = "hotel-reservation"
	TEST_DB_NAME = "hotel-reservation-test"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}
