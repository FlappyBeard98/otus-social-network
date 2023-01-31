package service

import (
	"database/sql"
	"social-network/lib/http"
	"social-network/lib/mysql"
	"social-network/services/profile/internal/types"

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

func (o *App) NewUserBasicAuth() func(string, string, echo.Context) (bool, error) {
	return func(login string, password string, context echo.Context) (bool, error) {
		givenAuth, err := types.NewAuth(login, password)

		if err != nil {
			return false, err
		}

		dbAuth := new(types.Auth)
		err = givenAuth.ReadByLogin().QueryOne(context.Request().Context(), o.Db, dbAuth)

		if err != nil {
			return false, err
		}

		if dbAuth.Password == givenAuth.Password {
			return true, nil
		} else {
			return false, nil
		}

	}
}
