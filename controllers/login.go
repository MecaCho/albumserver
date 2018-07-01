package controllers

import (
	"github.com/astaxie/beego"
	"album_server/models"
)

type LoginController struct {

	beego.Controller

}



func (c *LoginController) Loginpage() {

	c.Data["Username"] = "loongc"

	c.Data["Password"] = "123"

	c.Data["Error"] = ""

	c.TplName = "login.tpl"

}



func (c *LoginController) Login() {

	flash := beego.NewFlash()

	username, password := c.Input().Get("username"), c.Input().Get("password")

	if flag, _ := models.Login(username, password); flag {

		// c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.role"), user.Role, 1*1*20*60, beego.AppConfig.String("cookie.domain"), "/", false, true)

		c.Redirect("/index", 302)

		// c.TplName = "index.tpl"

		return

	} else {

		flash.Error("username or password error")

		c.Redirect("/error", 302)

		return

	}

}