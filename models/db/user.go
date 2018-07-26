package db

import (
	"fmt"
	"strings"
	"time"
	"reflect"
)

func GetTableName(t interface{})string {
	structTypeName := reflect.TypeOf(t).Name()
	fmt.Println(structTypeName)
	ret := strings.ToLower(structTypeName)
	fmt.Println(ret)
	return ret
}

func InsertPhoto(photo Photo) error {
	var maxIncID int
	rawQuery := fmt.Sprintf("SELECT max(inc_id) from %s",GetTableName(photo))
	err := ORM_.Raw(rawQuery).QueryRow(&maxIncID)
	photo.IncID = maxIncID + 1
	photo.UploadTime = time.Now().String()

	fmt.Printf("Insert Photo : %#v", photo)
	id, err := ORM_.Insert(&photo)
	if err != nil {
		fmt.Printf("Insert Photo (%q) Error (%d) : (%q)\n", photo.Id, id, err)
		return err
	}
	fmt.Printf("Insert Photo (%q) Success : %d", photo.Id, id)
	return nil
}

func GetPhoto(photoId string) (Photo, error) {
	var photo Photo
	qs := ORM_.QueryTable(photo)
	num, err := qs.Filter("id", photoId).All(&photo)
	if err != nil {
		fmt.Printf("Get Photo (%s) Error : (%q) , %s \n", photoId, err, num)
		return photo, err
	}
	fmt.Printf("Get Photo (%q) Success: %#v , %s", photoId, photo, num)
	return photo, nil
}
