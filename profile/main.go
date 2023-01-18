package main

import (
	"fmt"
	"social-network/lib/http"
	"social-network/lib/mysql"
	service "social-network/profile/internal"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
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
	// r.POST("/register", wrap(app.Commands.Register))
	// r.GET("/profiles", wrap(app.Queries.ProfilesByFilter))

	// authed := r.Group("profile", middleware.BasicAuth(app.NewBasicAuth()))
	// authed.GET("/:userId", wrap(app.Queries.Profile))
	// authed.POST("/:userId", wrap(app.Commands.SaveProfile))
	// authed.GET("/:userId/friends", wrap(app.Queries.Friends))
	// authed.POST("/:userId/friends/:friendUserId", wrap(app.Commands.AddFriend))
	// authed.DELETE("/:userId/friends/:friendUserId", wrap(app.Commands.RemoveFriend))

	mysql.Migrate(cfg.Db)
	http.StartHttpServer(r, cfg.Http)
}
