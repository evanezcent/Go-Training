package service

import (
	"log"

	"../dto"
	"../model"
	"../repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

// AuthService interface that cover all function needed
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) model.User
	FindEmail(email string) model.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

// NewAuthService is used to create new Instance
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (serv *authService) VerifyCredential(email string, pass string) interface{} {
	res := serv.userRepository.VerifyCredential(email, pass)
	if val, ok := res.(model.User); ok {
		comparedPass := comparePassword(val.Password, []byte(pass))
		if val.Email == email && comparedPass {
			return res
		}

		return false
	}

	return false
}

func comparePassword(hashed string, pass []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, pass)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (serv *authService) CreateUser(user dto.UserCreateDTO) model.User {
	newUser := model.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := serv.userRepository.InsertUser(newUser)
	return res
}

func (serv *authService) FindEmail(email string) model.User {
	return serv.userRepository.FindEmail(email)
}

func (serv *authService) IsDuplicateEmail(email string) bool {
	res := serv.userRepository.IsDuplicateEmail(email)

	return !(res.Error == nil)
}
