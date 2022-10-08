package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"social-network/app"
	"social-network/common/application"
)

func SetupRoutes(echo *echo.Echo, app *app.App) {
	test := echo.Group("test")
	test.GET("/:name", wrap(app.Queries.Hello))

	echo.POST("/register", wrap(app.Commands.Register))
	echo.POST("/login", wrap(app.Commands.Register))
	echo.GET("/profiles", wrap(app.Queries.ProfilesByFilter))

	authed := echo.Group("public", middleware.BasicAuth(auth))
	authed.GET("/profile/:id", wrap(app.Queries.Profile))
	authed.POST("/profile/:id", wrap(app.Commands.SaveProfile))
	authed.GET("/profile/:id/friends", wrap(app.Queries.Friends))
	authed.POST("/profile/:id/friends", wrap(app.Commands.AddFriend))
	authed.DELETE("/profile/:id/friends", wrap(app.Commands.RemoveFriend))
	authed.POST("/logout", wrap(app.Commands.Logout))
}

func wrap[In any, Out any](handler application.Handler[In, Out]) echo.HandlerFunc {
	return func(c echo.Context) error {

		var arg In
		err := c.Bind(&arg)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		result, err := handler.Handle(c.Request().Context(), arg)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}

func auth(login string, password string, context echo.Context) (bool, error) {
	//todo setup basic auth
	return true, nil
}
