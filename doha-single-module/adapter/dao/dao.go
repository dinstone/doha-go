package dao

import (
	"doha-single-module/business/model"

	"github.com/beego/beego/v2/client/orm"
)

type UserDao struct {
	ormer orm.Ormer
}

func (dao *UserDao) All() (*[]model.User, error) {
	qt := dao.ormer.QueryTable("user")
	var users []model.User
	_, error := qt.All(&users)
	if error != nil {
		return nil, error
	}
	return &users, nil
}

func (dao *UserDao) Get(id int64) (*model.User, error) {
	qt := dao.ormer.QueryTable("user")
	var users []model.User
	count, error := qt.Filter("id", id).All(&users)
	if error != nil {
		return nil, error
	}

	if count > 0 {
		return &users[0], nil
	}
	return nil, nil
}

func (dao *UserDao) Save(user *model.User) (int64, error) {
	return dao.ormer.Insert(user)
}

func (dao *UserDao) Update(user *model.User) (int64, error) {
	return dao.ormer.Update(user)
}

func (dao *UserDao) Delete(id int64) (int64, error) {
	return dao.ormer.Delete(&model.User{Id: id})
}
