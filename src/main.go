package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"social-network/api"
	"social-network/app"
	"social-network/common/database"
	"social-network/common/http"
)

const configFileName = "config.json"

func main() {

	var cfg Config
	cfg.Load(configFileName)

	db := database.InitMysql(cfg.Database.ConnectionString)
	defer func() { _ = db.Close() }()

	a := app.NewApp(db)
	e := echo.New()

	api.SetupRoutes(e, a)
	http.StartHttpServer(e, cfg.GetEchoPort())
}

type Config struct {
	ApiPort  int `json:"apiPort" env:"API_PORT" env-default:"1323"`
	Database struct {
		ConnectionString string `json:"connectionString" env:"CONNECTION_STRING"`
	}
}

func (receiver *Config) Load(path string) {
	if err := cleanenv.ReadConfig(path, receiver); err != nil {
		panic(err.Error())
	}
}

func (receiver *Config) GetEchoPort() string {
	return fmt.Sprintf(":%d", receiver.ApiPort)
}
