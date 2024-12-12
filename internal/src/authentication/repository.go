package authentication

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"live-code-2-XioweL/internal/models"
	"log"
)

type AuthenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *AuthenticationRepository {
	return &AuthenticationRepository{db}
}

func (repo *AuthenticationRepository) RegisterUser(ctx context.Context, request models.User) (response models.User, err error) {
	tx := repo.db.Begin()

	if err = tx.WithContext(ctx).Model(&models.User{}).Create(&request).Error; err != nil {
		log.Printf(fmt.Sprintf("Error when insert user: %v", err))
		tx.Rollback()

		return
	}

	tx.Commit()
	return request, nil
}

func (repo *AuthenticationRepository) RegisterCustomer(ctx context.Context, request models.Customer) (err error) {
	tx := repo.db.Begin()

	if err = tx.WithContext(ctx).Model(&models.Customer{}).Create(&request).Error; err != nil {
		log.Printf(fmt.Sprintf("Error when insert customer: %v", err))
		tx.Rollback()

		return
	}

	tx.Commit()
	return
}
