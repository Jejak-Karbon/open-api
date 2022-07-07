package http

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/news"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	news.NewHandler(f).Route(e.Group("/news"))
}
