package dao

import (
	"doha-business/model"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"

	_ "github.com/mattn/go-sqlite3"
)

type Daos struct {
	UserDao *UserDao
}

func NewDaos(c config.Configer) *Daos {
	c.DefaultString("sql.driver", "sqlite3")

	// need to register models in init
	orm.RegisterModel(new(model.User))
	orm.RegisterModel(new(model.Profile))

	// need to register db driver
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// need to register default database
	orm.RegisterDataBase("default", "sqlite3", "db/doha.db")

	// automatically build table
	orm.RunSyncdb("default", false, true)

	// create orm object
	ormer := orm.NewOrm()

	return &Daos{UserDao: &UserDao{ormer: ormer}}
}
