package main

import (
	"fmt"

	"social-network/lib/http"
	"social-network/lib/mysql"
	service "social-network/profile/internal"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	cfgFile = "config.json"
	Retries = []time.Duration{2 * time.Second, 3 * time.Second, 5 * time.Second, 8 * time.Second, 13 * time.Second}
)

func main() {
	cfg := service.LoadConfig(cfgFile)
	app := service.App{
		Db: mysql.Connect(cfg.Db),
	}

	fmt.Printf("%v", app)

	r := echo.New()

	r.POST("/register", app.Register)
	r.GET("/profiles", app.Profiles)

	authed := r.Group("profile", middleware.BasicAuth(app.NewUserBasicAuth()))
	authed.GET("/:userId", app.Profile)
	authed.POST("/:userId", app.SaveProfile)
	authed.GET("/:userId/friends", app.Friends)
	authed.POST("/:userId/friends/:friendUserId", app.AddFriend)
	authed.DELETE("/:userId/friends/:friendUserId", app.DeleteFriend)

	mysql.Migrate(cfg.Db)
	http.StartHttpServer(r, cfg.Http)
}
