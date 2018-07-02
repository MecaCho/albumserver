package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"math/rand"
	"time"
	"fmt"
	"strconv"
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
	Id string
	Name string
	UploadTime string
	Size int
	Photos []Photo
}

var TempAlbumList []Album
var AlbumNum int64


func (c *AlbumController) ListAlbums() {
	fmt.Println("Album list : ",TempAlbumList)
	ret,err := json.Marshal(&TempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) GetAlbum() {
	fmt.Println("Album list : ",TempAlbumList)
	ret,err := json.Marshal(&TempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) DeleteAlbum() {
	fmt.Println("Album list : ",TempAlbumList)
	ret,err := json.Marshal(&TempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c *AlbumController) UploadAlbum() {
	fmt.Println("Album list : ",TempAlbumList)
	ret,err := json.Marshal(&TempAlbumList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c AlbumController) PostAlbum(){
	photoID := strconv.Itoa(rand.Int())
	photoUploadTime := time.Now().String()
	AlbumNum = int64(len(TempAlbumList))
	AlbumNum++
	newAlbum := Album{AlbumNum,photoID,"album name",photoUploadTime,0,nil}
	fmt.Println(newAlbum)
	TempAlbumList = append(TempAlbumList, newAlbum)
	fmt.Println(TempAlbumList)
	c.Ctx.Output.Body([]byte(photoID))

	return

}
