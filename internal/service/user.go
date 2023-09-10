package service

import (
	"fmt"
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/model"
	"github.com/formulationapp/formulationapp/internal/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type UserService interface {
	Login(email string, password string) (string, error)
	Register(email string, password string) (string, error)
	DecodeToken(token string) (uint, error)
}

type userService struct {
	userRepository repository.UserRepository
	config         dto.Config
}

func newUserService(userRepository repository.UserRepository, config dto.Config) UserService {
	return &userService{userRepository, config}
}

func (u userService) Login(email string, password string) (string, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return "", dto.AppError(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", dto.AppError(fmt.Errorf("invalid password"))
	}
	return u.generateJwt(user)
}

func (u userService) Register(email string, password string) (string, error) {
	if email == "" {
		return "", dto.AppError(fmt.Errorf("email is required"))
	}
	if password == "" {
		return "", dto.AppError(fmt.Errorf("password is required"))
	}
	if len(password) < 3 {
		return "", dto.AppError(fmt.Errorf("password must be at least 8 characters"))
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password")
	}
	userByEmail, err := u.userRepository.FindByEmail(email)
	if userByEmail.ID != 0 {
		return "", dto.AppError(fmt.Errorf("user with email %s already exists", email))
	}
	user := model.User{Email: email, Password: string(hashedPassword)}
	user, err = u.userRepository.Save(user)
	if err != nil {
		return "", fmt.Errorf("error saving user")
	}
	return u.generateJwt(user)
}

func (u userService) generateJwt(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(6 * time.Hour)
	claims["user_id"] = fmt.Sprint(user.ID)
	claims["email"] = user.Email
	tokenString, err := token.SignedString([]byte(u.config.SigningSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u userService) DecodeToken(token string) (uint, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.config.SigningSecret), nil
	})
	if err != nil {
		return 0, err
	}
	userIDString := claims["user_id"].(string)
	userID, err := strconv.ParseUint(userIDString, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}
