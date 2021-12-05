package user

import (
	"github.com/dhis2-sre/im-users/pkg/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindById(id uint) (*model.User, error)
}

func ProvideRepository(DB *gorm.DB) Repository {
	return &userRepository{db: DB}
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) Create(u *model.User) error {
	return r.db.Create(&u).Error
}

func (r userRepository) FindByEmail(email string) (*model.User, error) {
	var u *model.User
	err := r.db.Where("email = ?", email).First(&u).Error
	return u, err
}

func (r userRepository) FindById(id uint) (*model.User, error) {
	var u *model.User
	err := r.db.
		Preload("Groups").
		First(&u, id).Error
	return u, err
}
