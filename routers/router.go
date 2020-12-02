package routers

import (
	"github.com/astaxie/beego"
	"go-youku/controllers"
)

func init() {
    beego.Include(&controllers.UserController{})
    beego.Include(&controllers.VideoController{})
    beego.Include(&controllers.TopController{})
    beego.Include(&controllers.BaseController{})
}
