package db

type User struct {
	Id int `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Password string `json:"password"`
}
