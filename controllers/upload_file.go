package controllers

import (
	"net/http"
	"io"
	"fmt"
	"os"
)

const (
	upload_path string = "./upload/"
)

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

