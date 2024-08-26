package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost     = 12
	minNameLen     = 2
	minPasswordLen = 7
)

type User struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name" json:"name"`
	Email             string             `bso:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPasword" json:"-"`
}

type CreateUserParams struct {
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `json:"password"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:              params.Name,
		Email:             params.Email,
		EncryptedPassword: string(password),
	}, nil
}

func (params *CreateUserParams) Validate() map[string]string {
	var errors = make(map[string]string)
	if len(params.Name) < minNameLen {
		errors["name"] = fmt.Sprintf("name length should be at least %d characters", minNameLen)
	}

	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}

	if !isEmailValid(params.Email) {
		errors["email"] = "email is invalid"
	}

	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z-9._%]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

type UpdateUserParams struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}
