package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/db"
	"github.com/grayjunzi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRoomParams struct {
	FromDate   time.Time `json:"fromDate"`
	TillDate   time.Time `json:"tillDate"`
	PersonNums int       `json:"personNums"`
}

func (p BookRoomParams) validate() error {
	now := time.Now()
	if now.After(p.FromDate) || now.After(p.TillDate) {
		return fmt.Errorf("cannot book a room in the past")
	}

	return nil
}

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (h *RoomHandler) GetRooms(c *fiber.Ctx) error {
	rooms, err := h.store.Room.GetRooms(c.Context(), bson.M{})
	if err != nil {
		return err
	}

	return c.JSON(rooms)
}

func (h *RoomHandler) BookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.validate(); err != nil {
		return err
	}

	roomId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResponse{
			Type:    "error",
			Message: "internal server error",
		})
	}

	ok, err = h.isRoomAvailableForBooking(c.Context(), roomId, params)

	if err != nil {
		return nil
	}

	if !ok {
		return c.Status(http.StatusBadRequest).JSON(genericResponse{
			Type:    "error",
			Message: fmt.Sprintf("room %s already booked", c.Params("id")),
		})
	}

	booking := types.Booking{
		RoomId:     roomId,
		UserId:     user.Id,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
		PersonNums: params.PersonNums,
	}

	inserted, err := h.store.Booking.Insert(c.Context(), &booking)
	if err != nil {
		return err
	}

	return c.JSON(inserted)
}

func (h *RoomHandler) isRoomAvailableForBooking(ctx context.Context, roomId primitive.ObjectID, params BookRoomParams) (bool, error) {
	filter := bson.M{
		"roomId": roomId,
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$lte": params.TillDate,
		},
	}
	bookings, err := h.store.Booking.GetBookings(ctx, filter)
	if err != nil {
		return false, nil
	}

	ok := len(bookings) == 0
	return ok, nil
}
