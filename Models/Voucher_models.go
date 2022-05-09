package models

type Voucher struct {
	Id              uint   ` gorm:"primaryKey" json:"id"`
	Title           string `json:"title"`
	Deskripsi       string `json:"deskripsi"`
	Image           string `json:"image"`
	Vendor          string `json:"vendor"`
	Kode            string `json:"kode"`
	VoucherQuantity int    `json:"voucher_quantity"`
	CreatedAt       string `json:"created_at"`
}

func (Voucher) TableName() string {
	return "voucher"
}
