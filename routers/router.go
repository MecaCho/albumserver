package routers

import (
	"album_server/controllers"
	"github.com/astaxie/beego"
)

func init() {

	AlbumNS := beego.NewNamespace("/v1",
			beego.NSNamespace("/:user_id",
				beego.NSRouter("/albums",&controllers.AlbumController{},"GET:ListAlbums"),
				beego.NSRouter("/albums",&controllers.AlbumController{},"POST:PostAlbum"),
				beego.NSRouter("/albums/:album_id",&controllers.AlbumController{},"GET:GetAlbum"),
				beego.NSRouter("/albums/:album_id",&controllers.AlbumController{},"DELETE:DeleteAlbum"),
				beego.NSRouter("/albums",&controllers.AlbumController{},"PUT:UploadAlbum"),
			),
		)

	PhotoNS := beego.NewNamespace("/v1",
		beego.NSNamespace("/:user_id/albums/:album_id",
			beego.NSRouter("/photos",&controllers.PhotoController{},"get:ListPhotos"),
			beego.NSRouter("/photos",&controllers.PhotoController{},"post:PostPhoto"),
			beego.NSRouter("/photos/:photo_id",&controllers.PhotoController{},"get:GetPhoto"),
			beego.NSRouter("/photos/:photo_id",&controllers.PhotoController{},"delete:DeletePhoto"),
			beego.NSRouter("/photos",&controllers.PhotoController{},"put:UploadPhoto"),
		),
	)
	beego.AddNamespace(AlbumNS)
	beego.AddNamespace(PhotoNS)
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/albums",&controllers.PhotoController{},"*:ListPhotos")
	//beego.Router("/album/post",&controllers.PhotoController{},"*:PostPhoto")
}
