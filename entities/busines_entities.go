package entities

type Business struct {
	Cod          int64  `json:"cod" gorm:"column:business_cod; primaryKey; autoIncrement"`
	BusinessName string `json:"business_name" gorm:"column:business_name"`
	Address      string `json:"address" gorm:"column:address"`
	IsVerified   bool   `json:"is_verified" gorm:"column:is_verified; default:false"`
	BusinessURL  string `json:"business_url" gorm:"column:business_url"`
	BusinessPic  string `json:"business_pic" gorm:"column:business_pic"`
}
