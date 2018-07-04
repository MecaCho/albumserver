package main

import (
	_ "album_server/routers"
	"album_server/controllers"
)

func main() {
	//beego.Run()
	controllers.UploadFile()
}

