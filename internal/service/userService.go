package service

import (
	"Go-Ecom-Aws/internal/domain"
	"Go-Ecom-Aws/internal/dto"
	"log"
)

type UserService struct{}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println(input)
	// perform some db operation
	// business logic
	return "this is my token as of now", nil
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	// perform some db operation
	// business logic
	return nil, nil
}

func (s UserService) Login(input any) (string, error) {
	// perform some db operation
	// business logic
	return "", nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	// perform some db operation
	// business logic
	return 0, nil
}

func (s UserService) VerifyCode(id uint) error {
	// perform some db operation
	// business logic
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	// perform some db operation
	// business logic
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	// perform some db operation
	// business logic
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	// perform some db operation
	// business logic
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	// perform some db operation
	// business logic
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	// perform some db operation
	// business logic
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	// perform some db operation
	// business logic
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	// perform some db operation
	// business logic
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	// perform some db operation
	// business logic
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	// perform some db operation
	// business logic
	return nil, nil
}
