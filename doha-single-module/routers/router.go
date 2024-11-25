package routers

import (
	"doha-single-module/adapter/dao"
	"doha-single-module/business/service"
	"doha-single-module/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	daos := dao.NewDaos(beego.AppConfig)
	userService := service.NewUserService(daos.UserDao)

	beego.Router("/", &controllers.MainController{UserService: userService})
}
