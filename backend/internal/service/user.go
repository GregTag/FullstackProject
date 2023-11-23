package service

import (
	"backend/internal/entity"
	"backend/pkg/hash"
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

func checkCredentials(user_reg *entity.UserRegister) error {
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

func (s *UserService) Register(user_reg *entity.UserRegister) (*entity.UserInfo, error) {
	err := checkCredentials(user_reg)
	if err != nil {
		return nil, err
	}

	passwordHash, err := hash.GeneratePasswordHash(user_reg.Password)
	if err != nil {
		return nil, err
	}
	user_reg.Password = passwordHash

	user := entity.User{UserRegister: *user_reg, LastLoginAt: time.Now()}
	err = s.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}

	info, err := s.userToInfo(&user)
	return info, err
}

func (s *UserService) Login(user_log *entity.UserLogin) (*entity.UserInfo, error) {
	user, err := s.userRepository.GetByLogin(user_log.Login)
	if err != nil {
		return nil, err
	}
	err = hash.ComparePasswordWithHash(user_log.Password, user.Password)
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
	return err
}
