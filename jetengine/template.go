package jetengine

import (
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo"
	"github.com/lionelvon/e-shop/vm"
	"io"
)

type Template struct {
	View *jet.Set
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	template, err := t.View.GetTemplate(name)
	if err == nil {
		var result = data.(vm.TxViewModelGroup)
		template.Execute(w, result.ViewModels, result.Data)
	}
	return err
}
