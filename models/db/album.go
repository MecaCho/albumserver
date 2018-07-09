package db

import (
	"fmt"
	"strings"
	"time"
	//"reflect"
)


func InsertAlbum(photo Album) error {
	var maxIncID int
	rawQuery := fmt.Sprintf("SELECT max(inc_id) from %s",GetTableName(photo))
	err := ORM_.Raw(rawQuery).QueryRow(&maxIncID)
	photo.IncID = maxIncID + 1
	photo.UploadTime = time.Now().String()

	fmt.Printf("Insert Album : %#v", photo)
	id, err := ORM_.Insert(&photo)
	if err != nil {
		fmt.Printf("Insert Album (%q) Error (%d) : (%q)\n", photo.Id, id, err)
		return err
	}
	fmt.Printf("Insert Album (%q) Success : %d", photo.Id, id)
	return nil
}

func GetAlbum(photoId string) (Album, error) {
	var photo Album
	qs := ORM_.QueryTable(photo)
	num, err := qs.Filter("id", photoId).All(&photo)
	if err != nil {
		fmt.Printf("Get Album (%s) Error : (%q) , %s \n", photoId, err, num)
		return photo, err
	}
	fmt.Printf("Get Album (%q) Success: %#v , %s", photoId, photo, num)
	return photo, nil
}

func ListAlbums(projectID string) ([]Album, error) {
	var photo Album
	queryKey := strings.ToLower(projectID)
	nodeList := []Album{}
	num, err := ORM_.QueryTable(photo).Filter(queryKey, projectID).All(&nodeList)
	fmt.Println(num, err, nodeList)
	return nodeList, nil
}

func ListAlbumsByUserID(userID string) ([]Album, error) {
	var photo Album
	nodeList := []Album{}
	num, err := ORM_.QueryTable(photo).Filter("user_id", userID).All(&nodeList)
	fmt.Println(num, err, nodeList)
	return nodeList, nil
}

func DeleteAlbum(photoId string) error {
	photo, err := GetAlbum(photoId)
	if err != nil {
		return err
	}
	id, err := ORM_.Delete(&photo)
	if err != nil {
		return err
	}
	fmt.Printf("Delete Group : %#v , %s" , photo, id)
	return nil
}

func UpdateAlbum(photoId string, photo Album) error {
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
