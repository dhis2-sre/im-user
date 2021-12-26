package user

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/dhis2-sre/im-user/internal/apperror"
	"github.com/dhis2-sre/im-user/pkg/model"
	"golang.org/x/crypto/scrypt"
	"strings"
)

type Service interface {
	Signup(email string, password string) (*model.User, error)
	SignIn(email string, password string) (*model.User, error)
	FindById(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindOrCreate(email string, password string) (*model.User, error)
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
		return nil, apperror.NewInternal(message)
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = s.repository.Create(user)

	if err != nil && err.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" {
		return nil, apperror.NewBadRequest(err.Error())
	}

	if err != nil {
		return nil, apperror.NewInternal(err.Error())
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

func (s service) SignIn(email string, password string) (*model.User, error) {
	unauthorizedMessage := "Invalid email and password combination"

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, apperror.NewUnauthorized(unauthorizedMessage)
	}

	match, err := comparePasswords(user.Password, password)
	if err != nil {
		message := fmt.Sprintf("Password hashing failed: %s", err)
		return nil, apperror.NewInternal(message)
	}

	if !match {
		return nil, apperror.NewUnauthorized(unauthorizedMessage)
	}

	return user, nil
}

func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	passwordAndSalt := strings.Split(storedPassword, ".")
	if len(passwordAndSalt) != 2 {
		return false, fmt.Errorf("wrong password/salt format: %s", storedPassword)
	}

	salt, err := hex.DecodeString(passwordAndSalt[1])
	if err != nil {
		return false, fmt.Errorf("unable to verify user password")
	}

	hash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)
	if err != nil {
		return false, err
	}

	return hex.EncodeToString(hash) == passwordAndSalt[0], nil
}

func (s service) FindById(id uint) (*model.User, error) {
	return s.repository.FindById(id)
}

func (s service) FindByEmail(email string) (*model.User, error) {
	return s.repository.FindByEmail(email)
}

func (s service) FindOrCreate(email string, password string) (*model.User, error) {
	hashedPassword, err := hashPassword(password)

	if err != nil {
		message := fmt.Sprintf("Password hashing failed: %s", err)
		return nil, apperror.NewInternal(message)
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	return s.repository.FindOrCreate(user)
}
