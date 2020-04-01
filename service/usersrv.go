package service

import (
	"gowork/model"
	"sync"
)

//UserSrv service.
var UserSrv = &userService{
	mutex: &sync.Mutex{},
}

type userService struct {
	mutex *sync.Mutex
}

func (srv *userService) UpdateAvatar(userID uint, avatarURL string) error {
	if err := db.Model(&model.User{}).Where("id = ?", userID).Update("avatar_url", avatarURL).Error; err != nil {
		return err
	}
	return nil
}

func (srv *userService) SaveUser(user *model.User) error {
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (srv *userService) GetUserByUserID(userID uint) *model.User {
	ret := &model.User{}
	if err := db.Where("id = ?", userID).First(&ret).Error; err != nil {
		return nil
	}

	return ret
}

func (srv *userService) GetUserByUsername(username string) *model.User {
	ret := &model.User{}
	if err := db.Where("name = ?", username).First(&ret).Error; err != nil {
		return nil
	}

	return ret
}

func (srv *userService) AddUser(user *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *userService) UpdateUser(user *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}
