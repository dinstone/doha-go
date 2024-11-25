package service

import (
	"doha-single-module/business/model"
	"doha-single-module/business/repository"
	"errors"
	"strconv"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(rep repository.UserRepository) *UserService {
	return &UserService{repository: rep}
}

func (us UserService) FindUser(id int64) (*model.User, error) {
	return us.repository.Get(id)
}

func (us UserService) GetAll() (*[]model.User, error) {
	return us.repository.All()
}

func (us UserService) DeleteUser(uid int64) error {
	c, err := us.repository.Delete(uid)
	if err != nil {
		return err
	}
	if c == 0 {
		return errors.New("not fund user by id=" + strconv.FormatInt(uid, 10))
	}
	return nil
}

func (us UserService) SaveUser(user *model.User) error {
	if user.Id == 0 {
		uid, err := us.repository.Save(user)
		if err != nil {
			return err
		}
		user.Id = uid
	} else {
		c, err := us.repository.Update(user)
		if err != nil {
			return err
		}
		if c == 0 {
			return errors.New("not fund user by id=" + strconv.FormatInt(user.Id, 10))
		}
	}

	return nil
}
