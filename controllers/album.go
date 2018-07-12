package controllers

import (
	"album_server/models/db"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type AlbumController struct {
	beego.Controller
}

func (c *AlbumController) Albumpage() {
	user_id := c.GetString(":user_id")
	//c.HandlerFunc("/upload")
	//http.Handle("/upload",http.FileServer(http.Dir("./upload/")))
	photoFilePathList, _ := db.ListPhotosByUserID(user_id)
	//ret = <img src="/static/img/1.jpg" alt="" />
	ret := ""
	for k, v := range photoFilePathList {
		fmt.Println("Photo Path : ", k, v)
		ret += "<img src=\""
		ret += strings.TrimLeft(v.FilePath, ".")
		ret += "\" alt=\"\" />"
	}
	fmt.Println(ret)
	c.Data["BackGroundImage"] = strings.TrimLeft(v.FilePath, ".")
	c.Data["ImgSrc"] = ret
	c.Data["Username"] = "qwq"
	c.TplName = "album.tpl"
}

type AlbumList struct {
}

type Album struct {
	IncID      int64
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	UploadTime string  `json:"upload_time"`
	Size       int     `json:"size"`
	Photos     []Photo `json:"photos"`
	UserID     string  `json:"user_id"`
}

var TempAlbumList []Album
var AlbumNum int64

func (c *AlbumController) ListAlbums() {
	userId := c.GetString(":user_id")
	fmt.Printf("User (%s) Album list : %#v", userId, TempAlbumList)
	tempAlbumList := []Album{}
	for _, v := range TempAlbumList {
		if v.UserID == userId {
			tempAlbumList = append(tempAlbumList, v)
		}
	}
	ret, err := json.Marshal(&tempAlbumList)
	if err != nil {
		fmt.Println("Json marshal error : ", err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) GetAlbum() {
	fmt.Println("Get Album list : ", TempAlbumList)
	ret, err := json.Marshal(&TempAlbumList)
	if err != nil {
		fmt.Println("Json marshal error : ", err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) DeleteAlbum() {
	fmt.Println("Delete Album list : ", TempAlbumList)
	ret, err := json.Marshal(&TempAlbumList)
	if err != nil {
		fmt.Println("Json marshal error : ", err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) UploadAlbum() {
	fmt.Println("Upload Album list : ", TempAlbumList)
	ret, err := json.Marshal(&TempAlbumList)
	if err != nil {
		fmt.Println("Json marshal error : ", err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c AlbumController) PostAlbum() {
	userId := c.GetString(":user_id")

	InputBody := c.Ctx.Input.RequestBody
	album := Album{}

	fmt.Println(string(InputBody))
	err := json.Unmarshal(InputBody, &album)
	if err != nil {
		ret := fmt.Sprintf("Unmarshal RequestBody Error : %s", err.Error())
		fmt.Println(ret)
		c.Ctx.Output.Body([]byte(ret))
		return
	}

	photoUploadTime := time.Now().String()

	AlbumNum = int64(len(TempAlbumList))
	AlbumNum++
	albumID := strconv.Itoa(rand.Int())
	album.IncID = AlbumNum
	album.Id = albumID
	album.UploadTime = photoUploadTime
	album.UserID = userId
	//newAlbum := Album{AlbumNum,albumID,"album name",photoUploadTime,0,nil,userId}
	fmt.Println(album)
	TempAlbumList = append(TempAlbumList, album)
	ret, err := json.Marshal(&album)
	if err != nil {
		fmt.Println("album struct marshal Error", err.Error())
	}
	fmt.Println(TempAlbumList)
	c.Ctx.Output.Body(ret)

	return
}
