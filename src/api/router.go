package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"social-network/app/query"
	"social-network/common"
)

func SetupRoutes(echo *echo.Echo) {
	g := echo.Group("test")
	g.GET("/:name", wrap(query.NewHelloHandler()))
}

func wrap[In any, Out any](handler common.Handler[In, Out]) echo.HandlerFunc {
	return func(c echo.Context) error {

		var arg In
		err := c.Bind(&arg)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		result, err := handler.Handle(arg)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}
