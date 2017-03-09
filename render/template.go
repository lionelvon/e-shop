package render

import (
	"../vm"
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo"
	"io"
)

type Template struct {
	View *jet.Set
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	template, err := t.View.GetTemplate(name)
	if err == nil {
		if data != nil {
			var result = data.(vm.TxViewModelGroup)
			template.Execute(w, result.ViewModels, result.Data)
		} else {
			template.Execute(w, nil, nil)
		}
	}
	return err
}
