package user

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/dhis2-sre/im-users/pgk/helper"
	"github.com/dhis2-sre/im-users/pgk/model"
	"golang.org/x/crypto/scrypt"
)

type Service interface {
	Signup(email string, password string) (*model.User, error)
	//	Signin(email string, password string) (*model.User, error)
	//	FindByEmail(email string) (*model.User, error)
	//	FindById(id uint) (*model.User, error)
}

func ProvideService(repository Repository) Service {
	return &service{repository}
}

type service struct {
	repository Repository
}

func (s service) Signup(email string, password string) (*model.User, error) {
	hashedPassword, err := hashPassword(password)

	if err != nil {
		message := fmt.Sprintf("Password hashing failed: %s", err)
		return nil, helper.NewInternal(message)
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = s.repository.Create(user)

	if err != nil && err.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" {
		return nil, helper.NewBadRequest(err.Error())
	}

	if err != nil {
		return nil, helper.NewInternal(err.Error())
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	// example for making salt - https://play.golang.org/p/_Aw6WeWC42I
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// using recommended cost parameters from - https://godoc.org/golang.org/x/crypto/scrypt
	hash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	hashedPassword := fmt.Sprintf("%s.%s", hex.EncodeToString(hash), hex.EncodeToString(salt))

	return hashedPassword, nil
}
