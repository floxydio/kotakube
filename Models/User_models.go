package models

type User struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Umur      int    `json:"umur"`
	Username  string `json:"username"`
	Email     string `json:"email" gorm:"unique"`
	Kota      string `json:"kota"`
	NoTelepon string `json:"no_telepon"`
	Password  string `json:"password"`
	Point     string `json:"point"`
	Role      string `json:"role" default:"0"`
	CreatedAt string `json:"created_at"`
}

func (User) TableName() string {
	return "user"
}
