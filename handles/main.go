package handles

import (
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo"
	"github.com/lionelvon/e-shop/vm"
	"net/http"
)

func Main(c echo.Context) (err error) {
	vars := make(jet.VarMap)
	vars.Set("title", "ttitle")
	return c.Render(http.StatusOK, "home.jet", vm.TxViewModelGroup{ViewModels: vars, Data: "Hello"})
}
