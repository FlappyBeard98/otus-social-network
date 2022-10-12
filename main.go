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


func main() {


	cfg := loadConfig("config.json")


	dbase := database.InitMysql(cfg.Database.ConnectionString)
	defer func() { _ = dbase.Close() }()

	a := app.NewApp(dbase,cfg.AesKey)
	e := echo.New()

	api.SetupRoutes(e, a)
	http.StartHttpServer(e, cfg.GetEchoPort())
}

type Config struct {
	ApiPort  int `json:"apiPort" env:"API_PORT" env-default:"1323"`
	AesKey string ` env:"AES_KEY" env-default:"0de62d7d94e24fb18065f4693d3fa1d3"`
	Database struct {
		ConnectionString string `json:"connectionString" env:"CONNECTION_STRING"`
	}
}

func loadConfig(path string) Config {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}

func (receiver *Config) GetEchoPort() string {
	return fmt.Sprintf(":%d", receiver.ApiPort)
}
