package service

import (
	"backend/internal/entity"
	"backend/pkg/helpers"

	"time"

	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepository entity.UserRepository
}

func NewUserService(userRepository entity.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func checkCredentials(userReg *entity.UserRegister) error {
	return nil
}

func (s *UserService) userToInfo(user *entity.User) (*entity.UserInfo, error) {
	var view entity.UserView
	err := copier.Copy(&view, &user)
	if err != nil {
		return nil, err
	}

	info, err := s.userRepository.FillTracks(&view)
	return info, err
}

func (s *UserService) Register(userReg *entity.UserRegister) (*entity.UserInfo, error) {
	err := checkCredentials(userReg)
	if err != nil {
		return nil, err
	}

	passwordHash, err := helpers.GeneratePasswordHash(userReg.Password)
	if err != nil {
		return nil, err
	}
	userReg.Password = passwordHash

	user := entity.User{UserRegister: *userReg, LastLoginAt: time.Now()}
	err = s.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}

	info, err := s.userToInfo(&user)
	return info, err
}

func (s *UserService) Login(userLog *entity.UserLogin) (*entity.UserInfo, error) {
	user, err := s.userRepository.GetByLogin(userLog.Login)
	if err != nil {
		return nil, err
	}
	err = helpers.ComparePasswordWithHash(userLog.Password, user.Password)
	if err != nil {
		return nil, err
	}

	user.LastLoginAt = time.Now()
	err = s.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	info, err := s.userToInfo(user)
	return info, err
}

func (s *UserService) Load(login string) (*entity.UserInfo, error) {
	user, err := s.userRepository.GetByLogin(login)
	if err != nil {
		return nil, err
	}

	info, err := s.userToInfo(user)
	return info, err
}

func (s *UserService) Change(view *entity.UserView) error {
	var user entity.User
	err := copier.Copy(&user, &view)
	if err != nil {
		return err
	}
	err = s.userRepository.Update(&user)
	copier.Copy(&view, &user)
	return err
}
