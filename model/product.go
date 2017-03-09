package model

type Product struct {
	Id          int    `xorm:"not null pk autoincr unique INT(11)"`
	ProductName string `xorm:"VARCHAR(50)"`
	ImageUrl    string `xorm:"VARCHAR(200)"`
	Description string `xorm:"TEXT"`
	Price       int    `xorm:"INT(11)"`
}
