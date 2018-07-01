package routers

import (
	"album_server/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/albums",&controllers.PhotoController{},"*:ListPhotos")
	beego.Router("/album/post",&controllers.PhotoController{},"*:PostPhoto")
}
