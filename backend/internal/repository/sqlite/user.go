package repository_sqlite

import (
	"backend/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type UserSQLite struct {
	db *gorm.DB
}

func NewUserSQLite(db *gorm.DB) *UserSQLite {
	return &UserSQLite{db: db}
}

func (r *UserSQLite) Create(user *entity.User) error {
	err := r.db.Create(user).Error

	if err != nil {
		if checkPrimaryKeyError(err) {
			return entity.ErrUserAlreadyExists
		} else {
			return err
		}
	}
	return nil
}

func (r *UserSQLite) Update(user *entity.User) error {
	result := r.db.Model(user).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return entity.ErrUserNotFound
	}
	return nil
}

func (r *UserSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return entity.ErrUserNotFound
	}
	return nil
}

func (r *UserSQLite) Get(id uint) (*entity.User, error) {
	var user entity.User

	result := r.db.First(&user, id)
	if result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrUserNotFound
	} else {
		return nil, result.Error
	}
}

func (r *UserSQLite) GetByLogin(login string) (*entity.User, error) {
	var user entity.User

	result := r.db.Model(&entity.User{}).Where("login = ?", login).First(&user)
	if result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrUserNotFound
	} else {
		return nil, result.Error
	}
}

func (r *UserSQLite) GetViewByLogin(login string) (*entity.UserView, error) {
	var user entity.UserView

	result := r.db.Model(&entity.User{}).Where("login = ?", login).First(&user)
	if result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrUserNotFound
	} else {
		return nil, result.Error
	}
}

func (r *UserSQLite) FillTracks(view *entity.UserView) (*entity.UserInfo, error) {
	tracks, err := NewMediaTrackSQLite(r.db).LoadAll(view.ID)
	if err != nil {
		return nil, err
	}
	return &entity.UserInfo{User: *view, Tracks: *tracks}, nil
}
