package user

import (
	"errors"
	"fmt"

	"github.com/dhis2-sre/im-user/internal/errdef"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindById(id uint) (*model.User, error)
	FindOrCreate(email *model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(DB *gorm.DB) *userRepository {
	return &userRepository{db: DB}
}

type duplicateError struct{ error }

func (e duplicateError) Duplicate() {}

func (e duplicateError) Unwrap() error {
	return e.error
}

func (r userRepository) Create(u *model.User) error {
	err := r.db.Create(&u).Error

	var perr *pgconn.PgError
	const uniqueKeyConstraint = "23505"
	if errors.As(err, &perr) && perr.Code == uniqueKeyConstraint {
		return duplicateError{error: fmt.Errorf("user already exists: %v", err)}
	}

	return err
}

func (r userRepository) FindByEmail(email string) (*model.User, error) {
	var u *model.User
	err := r.db.Where("email = ?", email).First(&u).Error
	return u, err
}

func (r userRepository) FindOrCreate(user *model.User) (*model.User, error) {
	var u *model.User
	err := r.db.Where(model.User{Email: user.Email}).Attrs(model.User{Password: user.Password}).FirstOrCreate(&u).Error
	return u, err
}

func (r userRepository) FindById(id uint) (*model.User, error) {
	var u *model.User
	err := r.db.
		Preload("Groups").
		First(&u, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := fmt.Errorf("failed to find user with id %d: %v", id, err)
			return u, errdef.NotFound(err)
		}
		return u, err
	}

	return u, nil
}
