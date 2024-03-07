package config

import (
	"fmt"
	"log"
	"time"

	// "godgifu-server/modules/account/routes"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/redis/go-redis/v9"
)

// Configuration struct fields are used for configuring the server
type Configuration struct {
	Router          *echo.Echo
	BaseURL         string
	Port            string
	TimeoutDuration time.Duration
	//Enviroment string
	Postgres *sqlx.DB
	// NoSQL
	// Redis      *redis.Client
	// FilesDir   string
	// LoggerDir  string
}

func LoadConfig() (*Configuration, error) {

	// Postgres verify connection
	log.Printf("Config: Connecting to Postgres")
	pgConnectionString := ""
	postgres, err := sqlx.Connect("pgx", pgConnectionString)
	postgres.SetMaxIdleConns(5)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	// Redis verify connection
	// log.Printf("Connecting to redis")
	// redisDB := redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%s:%s", redis_host, redis_port),
	// 	Password: "",
	// 	DB:       0,
	// })

	router := echo.New()
	router.Use(middleware.Recover(), middleware.Logger())

	config := &Configuration{
		Router:          router,
		Port:            "4000",
		Postgres:        postgres,
		BaseURL:         "localhost",
		TimeoutDuration: 5 * time.Second,
	}

	return config, nil
}
