package user

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/dhis2-sre/im-user/internal/errdef"

	"github.com/dhis2-sre/im-user/pkg/model"
	"golang.org/x/crypto/scrypt"
)

func NewService(repository userRepository) *service {
	return &service{repository}
}

type userRepository interface {
	create(user *model.User) error
	findByEmail(email string) (*model.User, error)
	findById(id uint) (*model.User, error)
	findOrCreate(email *model.User) (*model.User, error)
}

type service struct {
	repository userRepository
}

func (s service) SignUp(email string, password string) (*model.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("password hashing failed: %s", err)
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = s.repository.create(user)
	if err != nil {
		return nil, err
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
	unauthorizedError := fmt.Errorf("invalid email and password combination")

	user, err := s.repository.findByEmail(email)
	if err != nil {
		if errdef.IsNotFound(err) {
			return nil, errdef.NewUnauthorized(unauthorizedError)
		}
		return nil, err
	}

	match, err := comparePasswords(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("password hashing failed: %s", err)
	}

	if !match {
		return nil, errdef.NewUnauthorized(unauthorizedError)
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
	return s.repository.findById(id)
}

func (s service) FindOrCreate(email string, password string) (*model.User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("password hashing failed: %s", err)
	}

	user := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	return s.repository.findOrCreate(user)
}
