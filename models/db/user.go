package db

import (
	"fmt"
	"time"
)


func InsertUser(user User) error {
	var maxIncID int
	rawQuery := fmt.Sprintf("SELECT max(inc_id) from %s",GetTableName(user))
	err := ORM_.Raw(rawQuery).QueryRow(&maxIncID)
	user.IncID = maxIncID + 1
	user.UpdatedAt = time.Now().String()

	fmt.Printf("Insert Photo : %#v", user)
	id, err := ORM_.Insert(&user)
	if err != nil {
		fmt.Printf("Insert Photo (%q) Error (%d) : (%q)\n", user.Id, id, err)
		return err
	}
	fmt.Printf("Insert Photo (%q) Success : %d", user.Id, id)
	return nil
}

func GetUser(userName string) (string, error) {
	var user User
	qs := ORM_.QueryTable(user)
	num, err := qs.Filter("Name", userName).All(&user)
	if err != nil {
		fmt.Printf("Get User (%s) Error : (%q) , %s \n", userName, err, num)
		return "", err
	}
	fmt.Printf("Get PWD (%q) Success: %#v , %s", userName, user, num)
	return user.Password, nil
}
