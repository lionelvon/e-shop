package vm

import "../conf"

type TxProduct struct {
	Id          int
	ProductName string
	ImagePath   string
	Description string
	Price       int
}

func (product *TxProduct) ImageUrl(imageUrl string) {
	product.ImagePath = conf.ITEM_IMAGE_PATH + imageUrl
}
