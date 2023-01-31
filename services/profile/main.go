// Package: main
package main

import (
	"database/sql"
	"social-network/lib/http"
	"social-network/lib/mysql"
	"social-network/lib/utils"
	service "social-network/services/profile/internal"
	"time"

	"github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// createRouter creates a new router
func createRouter(app service.App) *echo.Echo {
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
	cfgFile = "config.json"                      // Path to config file
	qaKey   = "0567904c9b85418084917772d29d0e6d" // QA key
)

// @title           Swagger Social-network API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	cfg := service.LoadConfig(cfgFile)
	delays := []time.Duration{3 * time.Second, 5 * time.Second, 8 * time.Second}

	db, err := utils.Retry(func() (*sql.DB, error) { return mysql.Connect(cfg.Db) }, delays...)

	if err != nil {
		panic(err)
	}

	app := service.App{
		Db: db,
	}

	r := createRouter(app)

	http.StartHttpServer(r, cfg.Http)
}
