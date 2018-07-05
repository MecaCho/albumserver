package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"time"
	"fmt"
	"strconv"
	"math/rand"
)



type AlbumController struct {
	beego.Controller
}



func (c *AlbumController) Albumpage() {

	c.Data["Username"] = "loongc"

	c.Data["Password"] = "123"

	c.Data["Error"] = ""

	c.TplName = "login.tpl"

}

type AlbumList struct {

}

type Album struct {
	IncID int64
	Id string `json:"id"`
	Name string `json:"name"`
	UploadTime string `json:"upload_time"`
	Size int `json:"size"`
	Photos []Photo `json:"photos"`
	UserID string `json:"user_id"`
}

var TempAlbumList []Album
var AlbumNum int64


func (c *AlbumController) ListAlbums() {
	userId := c.GetString(":user_id")
	fmt.Printf("User (%s) Album list : %#v",userId,TempAlbumList)
	tempAlbumList := []Album{}
	for _,v := range(TempAlbumList){
		if v.UserID == userId{
			tempAlbumList = append(tempAlbumList, v)
		}
	}
	ret,err := json.Marshal(&tempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) GetAlbum() {
	fmt.Println("Get Album list : ",TempAlbumList)
	ret,err := json.Marshal(&TempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) DeleteAlbum() {
	fmt.Println("Delete Album list : ",TempAlbumList)
	ret,err := json.Marshal(&TempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) UploadAlbum() {
	fmt.Println("Upload Album list : ",TempAlbumList)
	ret,err := json.Marshal(&TempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c AlbumController) PostAlbum(){
	userId := c.GetString(":user_id")

	InputBody := c.Ctx.Input.RequestBody
	album := Album{}

	fmt.Println(string(InputBody))
	err := json.Unmarshal(InputBody,&album)
	if err != nil{
		ret := fmt.Sprintf("Unmarshal RequestBody Error : %s",err.Error())
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
	ret,err := json.Marshal(&album)
	if err != nil{
		fmt.Println("album struct marshal Error",err.Error())
	}
	fmt.Println(TempAlbumList)
	c.Ctx.Output.Body(ret)

	return
}
