package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"godgifu/logger"

	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/redis/go-redis/v9"
)

// Configuration struct fields are used for configuring the server
type DevConfiguration struct {
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

func DevLoadConfig() (*Configuration, error) {

	// ----- Dev only: remove before building Production -----
	devENV, err := devLoadENV()
	if err != nil {
		log.Fatal("Error: Config cannot find .env.dev file")
	}
	// ----- Debug only: remove before building Production -----

	// Postgres verify connection
	log.Printf("Config: Connecting to Postgres")
	pgConnectionString := devENV.pgConn
	postgres, err := sqlx.Connect("pgx", pgConnectionString)
	if err != nil {
		return nil, fmt.Errorf("error with sqlx.Connect: %v", err)
	}
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
	router.Use(middleware.RecoverWithConfig(
		middleware.RecoverConfig{
			Skipper:           middleware.DefaultSkipper,
			StackSize:         3 << 10,
			DisableStackAll:   true,
			DisablePrintStack: false,
			LogLevel:          0,
			LogErrorFunc:      middleware.DefaultRecoverConfig.LogErrorFunc,
		},
	))
	router.Use(logger.RequestLogger())
	router.Use(middleware.CORS())

	router.Validator = &CustomValidator{
		validator: validator.New(),
	}

	config := &Configuration{
		Router:          router,
		BaseURL:         devENV.baseURL,
		Port:            devENV.port,
		Postgres:        postgres,
		TimeoutDuration: 5 * time.Second,
	}

	return config, nil
}

type dev_env struct {
	baseURL    string
	port       string
	enviroment string
	pgConn     string
	redisHost  string
	redisPort  string
	filesDir   string
	loggerDir  string
}

// debugLoadENV uses .env files for loading config variables. Use flags in production
func devLoadENV() (*dev_env, error) {
	envLoadError := godotenv.Load(".env.dev")
	if envLoadError != nil {
		log.Fatal("Error loading env file: ", envLoadError)
	}

	baseURL := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	enviroment := os.Getenv("ENVIROMENT")
	filesDir := os.Getenv("FILES_DIR")
	loggerDir := os.Getenv("LOGGER_DIR")

	pgUser := os.Getenv("DB_USERNAME")
	pgPassword := os.Getenv("DB_PASSWORD")
	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgName := os.Getenv("DB_NAME")
	pgSSL := os.Getenv("DB_SSL")
	pgConnectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", pgUser, pgPassword, pgHost, pgPort, pgName, pgSSL)

	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")

	return &dev_env{
		baseURL:    baseURL,
		port:       port,
		enviroment: enviroment,
		pgConn:     pgConnectionString,
		redisHost:  redis_host,
		redisPort:  redis_port,
		filesDir:   filesDir,
		loggerDir:  loggerDir,
	}, nil
}

// Use in the server graceful shutdown to close all remote Data Storage connections (Postgres, Redis, Cloud, etc.)
func (config *Configuration) CloseDataStorageConnections() error {
	if err := config.Postgres.Close(); err != nil {
		return fmt.Errorf("error, closing Postgres database: %w", err)
	}

	//Redis
	// if err := config.Redis.Close(); err != nil {
	// 	return fmt.Errorf("error, closing Redis database:%w", err)
	// }

	//Cloud Storage(?)

	//CDN for files(?)

	return nil
}
