package controller
import "github.com/astaxie/beego"

type UserControler struct {
	beego.Controller
}
func(this*UserControler)ShowLogin(){
	this.TplName="login.html"
}