// Package: main
package main

import (
	"fmt"
	"social-network/lib/http"
	"social-network/lib/pg"
	"social-network/lib/utils"
	service "social-network/services/profile/internal"
	"time"

	_ "social-network/services/profile/docs"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// createRouter creates a new router
func createRouter(app service.App) *echo.Echo {
	r := echo.New()

	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	r.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	r.POST("/register", app.Register)
	r.GET("/profiles", app.Profiles)

	authed := r.Group("profile", middleware.BasicAuth(app.NewUserBasicAuth()))
	authed.GET("/:userId", app.Profile)
	authed.POST("/:userId", app.SaveProfile)
	authed.GET("/:userId/friends", app.Friends)
	authed.POST("/:userId/friends/:friendUserId", app.AddFriend)
	authed.DELETE("/:userId/friends/:friendUserId", app.DeleteFriend)

	r.GET("/swagger*", echoSwagger.WrapHandler)

	_ = r.Group("admin", http.NewKeyMiddleware(qaKey))

	return r
}

var (
	cfgFile = "config.json"                      // Path to config file
	qaKey   = "0567904c9b85418084917772d29d0e6d" // QA key
)

// @title           Swagger social-network API
// @version         1.0
// @description     This is a sample social-network server.
// @host      localhost:1323
// @BasePath  /
// @securityDefinitions.basic  BasicAuth
func main() {
	cfg := service.LoadConfig(cfgFile)
	delays := []time.Duration{3 * time.Second, 5 * time.Second, 8 * time.Second}

	db, err := utils.Retry(createConnectRetry(cfg.Db), delays...)

	if err != nil {
		panic(err)
	}

	app := service.App{
		Db: db,
	}

	r := createRouter(app)

	http.StartHttpServer(r, cfg.Http)
}

func createConnectRetry(db pg.DbConfig) func() (*pgxpool.Pool, error) {
	return func() (*pgxpool.Pool, error) {
		c, err := pg.Connect(db)
		if err != nil {
			return nil, fmt.Errorf("%w; connection %s", err, db.ConnectionString)
		}
		return c, nil
	}
}
