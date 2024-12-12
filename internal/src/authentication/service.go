package authentication

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"live-code-2-XioweL/internal/models"
	"log"
)

type AuthenticationService struct {
	AuthenticationRepository *AuthenticationRepository
}

func NewAuthenticationService(authenticationRepository *AuthenticationRepository) *AuthenticationService {
	return &AuthenticationService{AuthenticationRepository: authenticationRepository}
}

func (service *AuthenticationService) RegisterUser(ctx context.Context, request models.RegisterRequest) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	user := models.User{
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
	}

	userResponse, err := service.AuthenticationRepository.RegisterUser(ctx, user)
	if err != nil {
		log.Println(err)
		return
	}

	customer := models.Customer{
		UserID:      userResponse.UserID,
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
	}

	if err = service.AuthenticationRepository.RegisterCustomer(ctx, customer); err != nil {
		log.Println(err)
		return
	}

	return
}
