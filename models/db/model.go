package db

type User struct {
	IncID     int    `orm:"pk;column(inc_id)";json:"inc_id"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreateAt  string `json:"create_at"`
	UpdatedAt string `json:"updated_at"`
}

type Album struct {
	IncID      int    `orm:"pk;column(inc_id)";json:"inc_id"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	UploadTime string   `json:"upload_time"`
	CreateAt   string `json:"create_at"`
	UpdatedAt  string `json:"updated_at"`
	Size       int      `json:"size"`
	PhotoNum   int      `json:"photo_num"`
	//Photos     []*Photo `orm:"rel(m2m)";json:"photos"`
	UserID     string   `json:"user_id"`
}

type Photo struct {
	IncID      int    `orm:"pk;column(inc_id)";json:"inc_id"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	UserID string `orm:"column(user_id)";json:"user_id"`
	Size       int    `json:"size"`
	UploadTime string `json:"upload_time"`
	CreateAt   string `json:"create_at"`
	UpdatedAt  string `json:"updated_at"`
	FilePath string `json:"file_path"`
	AlbumID string `orm:"column(album_id)";json:"album_id"`
	//Album      *Album `orm:"rel(fk)";json:"album"`
}
