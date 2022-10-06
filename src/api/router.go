package api

import (
	"net/http"
	"social-network/app/command"
	"social-network/app/query"
	"social-network/common"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gogo/protobuf/vanity/command"
	"github.com/labstack/echo/v4"
	"github.com/zenazn/goji/web/middleware"
)

func SetupRoutes(echo *echo.Echo) {
	test := echo.Group("test")
	test.GET("/:name", wrap(query.NewHelloHandler()))

	echo.POST("/register",wrap(command.NewRegisterHandler()))
	echo.POST("/login",wrap(command.NewLoginHandler()))
	echo.GET("/profiles",wrap(query.NewProfilesByFilterHandler()))

	
	authed := echo.Group("public",middleware.BasicAuth()) //todo setup middleware
	authed.GET("/profile/:id",wrap(query.NewProfileHandler()))
	authed.POST("/profile/:id",wrap(command.NewSaveProfileHandler()))
	authed.GET("/profile/:id/friends",wrap(query.NewFriendsHandler()))
	authed.POST("/profile/:id/friends",wrap(command.NewAddFriendHandler()))
	authed.DELETE("/profile/:id/friends",wrap(command.NewRemoveFriendHandler()))
	authed.POST("/logout",wrap(command.NewLogoutHandler()))
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
