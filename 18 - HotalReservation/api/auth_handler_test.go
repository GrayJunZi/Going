package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/grayjunzi/hotel-reservation/db"
	"github.com/grayjunzi/hotel-reservation/types"
)

func insertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Name:     "james",
		Email:    "james@email.com",
		Password: "test",
	})

	if err != nil {
		log.Fatal(err)
	}

	_, err = userStore.CreateUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func TestAuthenticate(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	insertedUser := insertTestUser(t, tdb.UserStore)

	app := fiber.New()

	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.Authenticate)

	params := AuthParams{
		Email:    "james@email.com",
		Password: "test",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected http status of 200 but got %d", resp.StatusCode)
	}

	var authResponse AuthResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		t.Error(err)
	}

	if authResponse.Token == "" {
		t.Fatalf("expected the JWT Token to be present in the auth response")
	}

	insertedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(insertedUser, authResponse.User) {
		t.Fatalf("expected the user to be the inserted user")
	}

	fmt.Println(resp)
}
