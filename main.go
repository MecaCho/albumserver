package main

import (
	_ "album_server/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()

	//upload_path := "static/"
	//user_id := "qwqty"
	//sta,err := os.Stat(upload_path + user_id)
	//if err != nil{
	//	fmt.Println(sta,err.Error())
	//	//err = os.Mkdir(upload_path + user_id,0777)
	//	//if err != nil {
	//	//	fmt.Println(err.Error())
	//	//	this.Ctx.Output.Body([]byte("文件夹创建失败 : "+upload_path + user_id+err.Error()))
	//	//	return
	//	//}
	//}
	//fmt.Println(sta)

	//album := db.Album{123,"","","",123,123,nil,"hjhj"}
	//ph := db.Photo{123,"fkfjkldfjls","name","",123,"","",&album}
	//db.InsertPhoto(ph)
	//controllers.UploadFile()
}
