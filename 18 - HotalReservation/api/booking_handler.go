package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/db"
	"github.com/grayjunzi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) GetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return err
	}

	return c.JSON(bookings)
}

func (h *BookingHandler) GetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingById(c.Context(), id)
	if err != nil {
		return err
	}

	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return err
	}

	if booking.UserId != user.Id {
		return c.Status(http.StatusUnauthorized).JSON(genericResponse{
			Type:    "error",
			Message: "not authorized",
		})
	}

	return c.JSON(booking)
}
