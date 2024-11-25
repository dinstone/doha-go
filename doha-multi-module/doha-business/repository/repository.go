package repository

import "doha-business/model"

type UserRepository interface {
	All() (*[]model.User, error)
	Get(id int64) (*model.User, error)
	Save(user *model.User) (int64, error)
	Update(user *model.User) (int64, error)
	Delete(id int64) (int64, error)
}
