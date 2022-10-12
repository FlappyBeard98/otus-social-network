package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"social-network/app"
	"social-network/app/query"
	"social-network/common/application"
)

func SetupRoutes(echo *echo.Echo, app *app.App) {

	echo.POST("/register", wrap(app.Commands.Register))
	echo.GET("/profiles", wrap(app.Queries.ProfilesByFilter))

	authed := echo.Group("profile", middleware.BasicAuth(newBasicAuth(app)))
	authed.GET("/:userId", wrap(app.Queries.Profile))
	authed.POST("/:userId", wrap(app.Commands.SaveProfile))
	authed.GET("/:userId/friends", wrap(app.Queries.Friends))
	authed.POST("/:userId/friends", wrap(app.Commands.AddFriend))
	authed.DELETE("/:userId/friends", wrap(app.Commands.RemoveFriend))
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

func newBasicAuth( app *app.App) func (string,string,echo.Context) (bool, error) {
	return func (login string, password string, context echo.Context) (bool, error) {
		ok, err := app.Queries.IsValidAuth.Handle(context.Request().Context(),query.IsValidAuthQuery{
			Login:    login,
			Password: password,
		})
		return ok, err
	}

}
