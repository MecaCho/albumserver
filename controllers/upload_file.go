package controllers

import (
	"net/http"
	"io"
	"fmt"
	"os"
	"github.com/astaxie/beego"
	"album_server/models/db"
)

type FileController struct {
	beego.Controller
}

const (
	upload_path string = "./static/img/"
)


func (this *FileController) UploadImage(){
	user_id := this.GetString(":user_id")

	r := this.Ctx.Request

	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//创建文件
	if _,err := os.Stat(upload_path + user_id);err != nil{
		//fmt.Println(sta.IsDir())
		err = os.Mkdir(upload_path + user_id,0777)
		if err != nil {
			fmt.Println(err.Error())
			this.Ctx.Output.Body([]byte("文件夹创建失败 : "+upload_path + user_id+err.Error()))
			return
		}
	}

	filePath := upload_path + user_id + "/" + head.Filename
	fW, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
		this.Ctx.Output.Body([]byte("文件创建失败 : "+err.Error()))
		return
	}
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		this.Ctx.Output.Body([]byte("文件保存失败 : "+err.Error()))
		return
	}
	fileInfo,err := os.Stat(filePath)
	if err != nil{
		fmt.Println("Read file Info error : ",err.Error())
	}
	photo := db.Photo{}
	photo.AlbumID = "test"
	photo.FilePath = filePath
	photo.CreateAt = fileInfo.ModTime().String()
	photo.UserID = user_id
	//photolist := []*Photo{&photo}
	//album_ := db.Album{1,"edf","qwq","","","",123,0,photolist,user_id}
	//photo.Album = &album_
	//db.InsertAlbum(album_)
	db.InsertPhoto(photo)
	//io.WriteString(w, head.Filename+" 保存成功")
	//http.Redirect(w, r, "/success", http.StatusFound)
	this.Ctx.Output.Body([]byte(head.Filename+" \n Upload Success"))

}

func (this *FileController) GetUploadPage(){
	this.TplName = "upload.tpl"
}

func load_success(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "上传成功!")
}

//上传
func uploadHandle(w http.ResponseWriter, r *http.Request) {
	//从请求当中判断方法
	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传图片</label>"+":"+
			"<input type=\"file\" name='file'  /><br/><br/>&nbsp;&nbsp;&nbsp;&nbsp;"+
			"<label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		//获取文件内容 要这样获取
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//创建文件
		fW, err := os.Create(upload_path + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}
		//io.WriteString(w, head.Filename+" 保存成功")
		http.Redirect(w, r, "/success", http.StatusFound)
		//io.WriteString(w, head.Filename)
	}
}

func UploadFile() {
	fmt.Println("OK!")
	//启动一个http 服务器
	http.HandleFunc("/success", load_success)
	//上传
	http.HandleFunc("/upload", uploadHandle)
	http.Handle("/",http.FileServer(http.Dir("./upload/")))
	//mux := http.NewServeMux()
	//mux.Handle("/view/",http.StripPrefix("/view/",http.FileServer(http.Dir("./upload/"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
		return
	}
	fmt.Println("服务器启动成功")
}

