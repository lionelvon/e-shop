package handles

import (
	"../db"
	"../model"
	"../vm"
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"net/http"
)

func Main(c echo.Context) (err error) {
	products := make([]model.Product, 0)
	err1 := db.MySQL().Find(&products)
	if err1 == nil {
		vars := make(jet.VarMap)
		txProducts := make([]vm.TxProduct, 0)
		copier.Copy(&txProducts, &products)
		vars.Set("products", txProducts)
		return c.Render(http.StatusOK, "home.jet", vm.TxViewModelGroup{ViewModels: vars, Data: "Hello"})
	}
	fmt.Println("eror:", err1)
	return c.Render(http.StatusBadRequest, "home.jet", nil)
}
