package service

import (
	"database/sql"
	"social-network/lib/http"
	"social-network/lib/mysql"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

type config struct {
	Db   mysql.DbConfig  `json:"db"`
	Http http.HttpConfig `json:"http"`
}

// LoadConfigFromJsonFile load configuration in Config struct
func LoadConfig(path string) config {
	var cfg config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}

	return cfg
}

type App struct {
	Db *sql.DB
}

func (o *App) NewBasicAuth() func(string, string, echo.Context) (bool, error) {
	return func(login string, password string, context echo.Context) (bool, error) {
		ok := true
		var err error = nil
		//ok, err := app.Queries.IsValidAuth.Handle(context.Request().Context(), query.IsValidAuthQuery{Login:    login,Password: password,})
		return ok, err
	}
}
