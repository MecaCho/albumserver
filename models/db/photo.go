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

func ListPhotos(projectID string) ([]Photo, error) {
	var photo Photo
	nodeList := []Photo{}
	num, err := ORM_.QueryTable(photo).Filter("project_id", projectID).All(&nodeList)
	fmt.Println(num, err, nodeList)
	return nodeList, nil
}

func DeletePhoto(photoId string) error {
	photo, err := GetPhoto(photoId)
	if err != nil {
		return err
	}
	id, err := ORM_.Delete(&photo)
	if err != nil {
		return err
	}
	fmt.Printf("Delete Group : %#v , %s", photo, id)
	return nil
}

func UpdatePhoto(photoId string, photo Photo) error {
	fmt.Println(photoId, photo.Id)
	if photoId == photo.Id {
		photo.UpdatedAt = time.Now().String()
		id, err := ORM_.Update(&photo)
		if err != nil {
			fmt.Printf("UPdate Group %s Error,detail : %#v \n", photoId, photo)
			return err
		}
		fmt.Printf("UPdate Group %s Success ,detail : %#v ; %q \n", photoId, photo, id)
		return nil
	}
	return fmt.Errorf("Update edgenode error , no such photo (%q) \n", photoId)

}
