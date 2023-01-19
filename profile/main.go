package main

import (
	"github.com/swaggo/echo-swagger"
	"social-network/lib/http"
	"social-network/lib/mysql"
	service "social-network/profile/internal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func createRouter(app service.App) *echo.Echo{
	r := echo.New()

	r.POST("/register", app.Register)
	r.GET("/profiles", app.Profiles)

	authed := r.Group("profile", middleware.BasicAuth(app.NewUserBasicAuth()))
	authed.GET("/:userId", app.Profile)
	authed.POST("/:userId", app.SaveProfile)
	authed.GET("/:userId/friends", app.Friends)
	authed.POST("/:userId/friends/:friendUserId", app.AddFriend)
	authed.DELETE("/:userId/friends/:friendUserId", app.DeleteFriend)

	admin := r.Group("admin", http.NewKeyMiddleware(qaKey))
	admin.GET("/swagger/*", echoSwagger.WrapHandler)

	return r
}

var (
	cfgFile = "config.json"
	qaKey = "0567904c9b85418084917772d29d0e6d"
)

func main() {
	cfg := service.LoadConfig(cfgFile)
	app := service.App{
		Db: mysql.Connect(cfg.Db),
	}

	r := createRouter(app)

	mysql.Migrate(cfg.Db)
	http.StartHttpServer(r, cfg.Http)
}