package user

import (
	"errors"
	"fmt"

	"github.com/dhis2-sre/im-user/internal/errdef"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

func NewRepository(DB *gorm.DB) *repository {
	return &repository{db: DB}
}

type repository struct {
	db *gorm.DB
}

func (r repository) create(u *model.User) error {
	err := r.db.Create(&u).Error

	var perr *pgconn.PgError
	const uniqueKeyConstraint = "23505"
	if errors.As(err, &perr) && perr.Code == uniqueKeyConstraint {
		return errdef.NewDuplicated("user already exists: %v", err)
	}

	return err
}

func (r repository) findByEmail(email string) (*model.User, error) {
	var u *model.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := fmt.Errorf("failed to find user with email %s: %v", email, err)
			return u, errdef.NotFound(err)
		}
	}
	return u, err
}

func (r repository) findOrCreate(user *model.User) (*model.User, error) {
	var u *model.User
	err := r.db.Where(model.User{Email: user.Email}).Attrs(model.User{Password: user.Password}).FirstOrCreate(&u).Error
	return u, err
}

func (r repository) findById(id uint) (*model.User, error) {
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
