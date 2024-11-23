// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"doha-business/service"
	"doha-handler/controllers"
	"doha-invoker/dao"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	daos := dao.NewDaos(beego.AppConfig)
	userService := service.NewUserService(daos.UserDao)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{UserService: userService},
			),
		),
	)
	beego.AddNamespace(ns)
}
