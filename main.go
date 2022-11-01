package main

import (
	"fmt"
	"social-network/api"
	"social-network/app"
	"social-network/common/database"
	"social-network/common/http"
	"social-network/tests"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

const (
	config    = "config.json"  //configuration file
	migration = "sql/init.sql" //file with sql migrations
)

func main() {

	cfg := loadConfig(config)

	dbase := database.InitMysql(cfg.Database.ConnectionString)
	defer func() { _ = dbase.Close() }()

	retries := []time.Duration{2 * time.Second, 3 * time.Second, 5 * time.Second, 8 * time.Second, 13 * time.Second}

	err := Retry(func() error { return database.Migrate(dbase, migration) }, retries)

	if err != nil {
		panic(err) 
	}

	a := app.NewApp(dbase, cfg.AesKey)

	tests.GenerateUsers(a,1000000)

	e := echo.New()

	api.SetupRoutes(e, a)
	http.StartHttpServer(e, cfg.GetEchoPort())
}

// Config stores application settings
type Config struct {
	ApiPort  int    `json:"apiPort" env:"API_PORT" env-default:"1323"`              //the port that the application will use
	AesKey   string ` env:"AES_KEY" env-default:"0de62d7d94e24fb18065f4693d3fa1d3"` //secret key that is used for encryption
	Database struct {
		ConnectionString string `json:"connectionString" env:"CONNECTION_STRING"` //database connection string
	}
}

// GetEchoPort returns formatted port for echo http server
func (receiver *Config) GetEchoPort() string {
	return fmt.Sprintf(":%d", receiver.ApiPort)
}

// loadConfig load configuration in Config struct
func loadConfig(path string) Config {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}

func Retry(fn func() error, delays []time.Duration) error {

	f := func() error {
		defer func() {
			if r := recover(); r != nil {
				println(r)
			}
		}()

		return fn()
	}

	var err error
	for _, d := range delays {

		err = f()
		if err != nil {
			println(err.Error())
		}
		time.Sleep(d)
	}
	return err
}
