package controllers

import (
	"album_server/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Loginpage() {
	c.Data["Username"] = "qwq"
	c.Data["Password"] = "123456"
	c.Data["Error"] = ""
	c.TplName = "login.tpl"
}

func (c *LoginController) Login() {

	flash := beego.NewFlash()
	fmt.Println("Get ...")
	username, password := c.Input().Get("username"), c.Input().Get("password")
	fmt.Println("user info : ", username, password)
	if flag, _ := models.Login(username, password); flag {

		// c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.role"), user.Role, 1*1*20*60, beego.AppConfig.String("cookie.domain"), "/", false, true)
		fmt.Println("Login success ")
		c.Redirect("v1/"+username+"/index", 302)
		fmt.Println("Redirect to index ")
		//c.TplName = "index.tpl"

		return

	} else {

		flash.Error("username or password error")

		c.Redirect("/error", 302)

		return

	}

}
