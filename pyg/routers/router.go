package routers

import (
	"github.com/astaxie/beego"
	"pygproject/pyg/controller"
)

func init() {
	beego.Router("/login",&controller.UserControler{},"get:ShowLogin")
}
