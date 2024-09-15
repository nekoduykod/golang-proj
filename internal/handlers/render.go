package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, t templ.Component) error {
	return t.Render(c.Request().Context(), c.Response())
}

// Example Render util func https://github.com/a-h/templ/blob/main/examples/integration-echo/main.go
// func Render(c echo.Context, statusCode int, t templ.Component) error {
// 	buf := templ.GetBuffer()
// 	defer templ.ReleaseBuffer(buf)

//		if err := t.Render(c.Request().Context(), buf); err != nil {
//			return err
//		}
//		return c.HTML(statusCode, buf.String())
//	}
