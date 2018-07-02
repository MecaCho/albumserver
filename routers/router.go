package routers

import (
	"album_server/controllers"
	"github.com/astaxie/beego"
)

func init() {

	AlbumNS := beego.NewNamespace("/v1",
			beego.NSNamespace("/:user_id",
				beego.NSRouter("/albums",&controllers.PhotoController{},"*:ListAlbums"),
				beego.NSRouter("/albums",&controllers.PhotoController{},"*:PostAlbum"),
				beego.NSRouter("/albums/:album_id",&controllers.PhotoController{},"*:GetAlbum"),
				beego.NSRouter("/albums/:album_id",&controllers.PhotoController{},"*:DeleteAlbum"),
				beego.NSRouter("/albums",&controllers.PhotoController{},"*:UploadAlbum"),
			),
		)

	PhotoNS := beego.NewNamespace("/v1",
		beego.NSNamespace("/:user_id/albums/:album_id",
			beego.NSRouter("/photos",&controllers.PhotoController{},"*:ListPhotos"),
			beego.NSRouter("/photos",&controllers.PhotoController{},"*:PostPhoto"),
			beego.NSRouter("/photos/:photo_id",&controllers.PhotoController{},"*:GetPhoto"),
			beego.NSRouter("/photos/:photo_id",&controllers.PhotoController{},"*:DeletePhoto"),
			beego.NSRouter("/photos",&controllers.PhotoController{},"*:UploadPhoto"),
		),
	)
	beego.AddNamespace(AlbumNS)
	beego.AddNamespace(PhotoNS)
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/albums",&controllers.PhotoController{},"*:ListPhotos")
	//beego.Router("/album/post",&controllers.PhotoController{},"*:PostPhoto")
}
