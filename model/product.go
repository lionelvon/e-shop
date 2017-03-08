package model

type Product struct {
	Id          int    `xorm:"not null pk autoincr unique INT(11)"`
	ProductName string `xorm:"VARCHAR(50)"`
	Description string `xorm:"TEXT"`
	Price       int    `xorm:"INT(11)"`
}
