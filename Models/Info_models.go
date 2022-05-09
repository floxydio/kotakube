package models

type InfoKota struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	Image     string `json:"image"`
	Tag       string `json:"tag"`
	Category  string `json:"category"`
	Kota      string `json:"kota"`
	View      string `json:"view" default:"0"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (InfoKota) TableName() string {
	return "info_kota"
}
