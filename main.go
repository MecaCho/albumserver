package main

import (
	_ "album_server/routers"
	//"album_server/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	//controllers.UploadFile()
}

