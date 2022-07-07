package weather

import (
	"github.com/labstack/echo/v4"
)

func Route(g *echo.Group) {
	g.GET("", Get)
	g.GET("/:id", GetByID)

}
