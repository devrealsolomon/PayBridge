package entities

type User struct {
	Cod      int64  `json:"cod" gorm:"column:user_cod; primaryKey; autoIncrement"`
	Name     string `json:"name" gorm:"column:user_name"`
	Username string `json:"username" gorm:"column:user_username; unique"`
	Phone    string `json:"phone" gorm:"column:user_phone"`
	Password string `json:"password" gorm:"column:user_password"`
}
