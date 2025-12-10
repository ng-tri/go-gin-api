package service

import (
	"errors"
	"strings"

	"go-gin-api/internal/database"
	"go-gin-api/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUser() ([]model.User, error)
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByPhone(phone string) (*model.User, error)
	CheckEmailOrPhoneExists(email, phone string) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetAllUser() ([]model.User, error) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	// Nếu muốn check trường hợp không có user nào
	if len(users) == 0 {
		return nil, errors.New("không có user nào trong hệ thống")
	}

	return users, nil
}

func (s *userService) CreateUser(user *model.User) error {

	if user.Email == "" && user.Phone == "" {
		return errors.New("email hoặc phone phải có ít nhất 1 giá trị")
	}

	if err := s.CheckEmailOrPhoneExists(user.Email, user.Phone); err != nil {
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	if err := database.DB.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("email hoặc số điện thoại đã tồn tại")
		}
		return err
	}

	return nil
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user không tồn tại")
		}
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) CheckEmailOrPhoneExists(email, phone string) error {
	var count int64

	if email != "" {
		database.DB.Model(&model.User{}).Where("email = ?", email).Count(&count)
		if count > 0 {
			return errors.New("email đã tồn tại")
		}
	}

	count = 0
	if phone != "" {
		database.DB.Model(&model.User{}).Where("phone = ?", phone).Count(&count)
		if count > 0 {
			return errors.New("số điện thoại đã tồn tại")
		}
	}

	return nil
}
