package service

import (
	"log"

	"../dto"
	"../model"
	"../repository"
	"github.com/mashingan/smapping"
)

// UserService interface for user service
type UserService interface {
	UpdateUser(user dto.UserUpdateDTO) model.User
	GetUser(userID string) model.User
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService is new Instance
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) UpdateUser(user dto.UserUpdateDTO) model.User {
	newUser := model.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(user))

	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}

	updatedUser := service.userRepository.UpdateUser(newUser)
	return updatedUser
}

func (service *userService) GetUser(id string) model.User {
	return service.userRepository.ProfileUser(id)
}
