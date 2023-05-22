package model

type Client struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number;unique;not null"`
	Password    string `json:"password" gorm:"column:password;not null"`
}
