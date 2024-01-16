package repository

import (
	"go-wallet/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	UsernameExist(username string) bool
	Register(req *models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) UsernameExist(email string) bool {
	var user models.User

	err := r.db.First(&user, "username = ?", user.Username).Error

	return err == nil
}

func (r *authRepository) Register(user *models.User) error {
	err := r.db.Create(&user).Error

	return err
}
