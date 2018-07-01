package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"math/rand"
	"time"
	"fmt"
	"strconv"
)



type PhotoController struct {
	beego.Controller
}



func (c *PhotoController) Photopage() {

	c.Data["Username"] = "loongc"

	c.Data["Password"] = "123"

	c.Data["Error"] = ""

	c.TplName = "login.tpl"

}

type PhotoList struct {

}

type Photo struct {
	Id string
	Name string
	UploadTime string
	Size int
}

var PhotosList []Photo


func (c *PhotoController) ListPhotos() {
	fmt.Println("Photo list : ",PhotosList)
	ret,err := json.Marshal(&PhotosList)
	if err != nil{
		fmt.Println("Json marshal error : ",err.Error())
	}
	fmt.Println(string(ret))
	c.Ctx.Output.Body(ret)
	return
}

func (c PhotoController) PostPhoto(){
	photoID := strconv.Itoa(rand.Int())
	photoUploadTime := time.Now().String()
	newPhoto := Photo{photoID,"photo name",photoUploadTime,0}
	fmt.Println(newPhoto)
	PhotosList = append(PhotosList, newPhoto)
	fmt.Println(PhotosList)
	c.Ctx.Output.Body([]byte(photoID))
	return

}