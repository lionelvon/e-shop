package vm

type TxProduct struct {
	Id          int    `json:"id"`
	ProductName string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
