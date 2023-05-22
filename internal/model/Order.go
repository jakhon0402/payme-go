package model

type Order struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Client    Client `json:"client" gorm:"column:client;not null"`
	OrderSum  int    `json:"order_sum" gorm:"column:order_sum;not null"`
	Cancelled bool   `json:"cancelled"`
	Paid      bool   `json:"paid"`
	Delivered bool   `json:"delivered"`
}

func (Order) TableName() string {
	return "orders"
}
