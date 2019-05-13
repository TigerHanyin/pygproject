package controller

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"pygproject/pyg/models"
	"regexp"
)

type UserControler struct {
	beego.Controller
}

func (this *UserControler) ShowLogin() {
	this.TplName = "login.html"
}
func (this *UserControler) HandleLogin() {
	//获取数据   注册的时候要求用户名必须为字母加数字
	userName := this.GetString("userName")
	pwd := this.GetString("pwd")
	if userName == "" || pwd == "" {
		beego.Error("用户名或密码错误")
		this.Data["errmsg"] = "用户名或密码错"
		this.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	reg, _ := regexp.Compile(`^\w[\w\.-]*@[0-9a-z][0-9a-z-]*(\.[a-z]+)*\.[a-z]{2,6}$`)
	result := reg.FindString(userName)
	if result != "" {
		user.Email = userName
		err := o.Read(&user, "Email")
		if err != nil {
			this.Data["errmsg"] = "邮箱未注册"
			this.TplName = "login.html"
			return
		}
		if user.Pwd != pwd {
			this.Data["errmsg"] = "密码错误"
			this.TplName = "login.html"
			return
		}

	} else {
		user.Name = userName
		err := o.Read(&user, "Name")
		if err != nil {
			this.Data["errmsg"] = "用户名不存在"
			this.TplName = "login.html"
			return
		}

	}
	if user.Active == false {
		this.Data["errmsg"] = "当前用户未激活，请去目标邮箱激活！"
		this.TplName = "login.html"
		return
	}
	m1 := this.GetString("m1")
	if m1 == "2" {
		this.Ctx.SetCookie("LoginName", user.Name, 60*60)
	} else {
		this.Ctx.SetCookie("LoginName", user.Name, -1)
	}
	this.SetSession("name", user.Name)
	this.Redirect("/index", 302)

}

