package config

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"godgifu/logger"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

type JWTFields struct {
	PrivateKey                 *rsa.PrivateKey
	PublicKey                  *rsa.PublicKey
	RefreshSecretKey           string
	IDTokenExpirationSecs      int64
	RefreshTokenExpirationSecs int64
}

type Router struct {
	API             string
	APIVersion      string
	BaseURL         string
	Enviroment      string
	Port            string
	TimeoutDuration time.Duration
}

// Configuration struct fields are used for configuring the server
type DevConfiguration struct {
	Echo     *echo.Echo
	Router   *Router
	Postgres *sqlx.DB
	// NoSQL
	Redis *redis.Client
	JWT   *JWTFields
	// FilesDir   string
	// LoggerDir  string
}

func DevLoadConfig() (*DevConfiguration, error) {

	// ----- Dev only: remove before building Production -----
	devENV, err := devLoadENV()
	if err != nil {
		log.Fatal("Error: Config cannot find .env.dev file")
	}
	// ----- Debug only: remove before building Production -----

	// ----- Begin RSA Keys -----
	log.Println("Loading RSA keys")

	privateKeyFile := os.Getenv("PRIVATE_KEY_FILE")
	privateKeyValue, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read private key file: %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyValue)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	publicKeyFile := os.Getenv("PUBLIC_KEY_FILE")
	publicKeyValue, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("could not read public key file: %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyValue)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}
	// ----- End RSA Keys -----

	// ----- Begin JWT values
	id_token_exp := os.Getenv("ID_TOKEN_EXP")
	refresh_token_exp := os.Getenv("REFRESH_TOKEN_EXP")
	refreshSecretKey := os.Getenv("REFRESH_SECRET_KEY")
	idTokenExpiration, err := strconv.ParseInt(id_token_exp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse JWT token expiration ENV to int: %w", err)
	}
	refreshTokenExpiration, err := strconv.ParseInt(refresh_token_exp, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse JWT refresh expiration ENV to int: %w", err)
	}
	log.Print("idToken expiration time: ", idTokenExpiration)
	log.Print("refreshToken expiration time: ", refreshTokenExpiration)
	// ----- End JWT values

	// ----- Begin Postgres verify connection -----
	log.Printf("Config: Connecting to Postgres")
	pgConnectionString := devENV.pgConn
	postgres, err := sqlx.Connect("pgx", pgConnectionString)
	if err != nil {
		return nil, fmt.Errorf("error with sqlx.Connect: %v", err)
	}
	postgres.SetMaxIdleConns(5)
	// ----- End Postgres verify connection

	// ----- Begin Redis verify connection -----
	log.Printf("Connecting to redis")
	redisDB := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", devENV.redisHost, devENV.redisPort),
		Password: "",
		DB:       0,
	})
	// ----- End Redis verify connection -----

	echoEngine := echo.New()
	echoEngine.Use(middleware.RecoverWithConfig(
		middleware.RecoverConfig{
			Skipper:           middleware.DefaultSkipper,
			StackSize:         3 << 10,
			DisableStackAll:   true,
			DisablePrintStack: false,
			LogLevel:          0,
			LogErrorFunc:      middleware.DefaultRecoverConfig.LogErrorFunc,
		},
	))
	echoEngine.Use(logger.RequestLogger())
	echoEngine.Use(middleware.CORS())

	echoEngine.Validator = &CustomValidator{
		validator: validator.New(),
	}

	config := &DevConfiguration{
		Echo: echoEngine,
		Router: &Router{
			API:             devENV.api,
			APIVersion:      devENV.apiVersion,
			BaseURL:         devENV.baseURL,
			Enviroment:      devENV.enviroment,
			Port:            devENV.port,
			TimeoutDuration: 5 * time.Second,
		},
		JWT: &JWTFields{
			PrivateKey:                 privateKey,
			PublicKey:                  publicKey,
			RefreshSecretKey:           refreshSecretKey,
			IDTokenExpirationSecs:      idTokenExpiration,
			RefreshTokenExpirationSecs: refreshTokenExpiration,
		},
		// Port:            devENV.port,
		Postgres: postgres,
		Redis:    redisDB,
		// TimeoutDuration: 5 * time.Second,
	}

	return config, nil
}

type dev_env struct {
	api        string
	apiVersion string
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

	api := os.Getenv("API")
	apiVersion := os.Getenv("API_VERSION")
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
		api:        api,
		apiVersion: apiVersion,
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
func (config *DevConfiguration) CloseDataStorageConnections() error {
	if err := config.Postgres.Close(); err != nil {
		return fmt.Errorf("error, closing Postgres database: %w", err)
	}

	//Redis
	if err := config.Redis.Close(); err != nil {
		return fmt.Errorf("error, closing Redis database:%w", err)
	}

	//Cloud Storage(?)

	//CDN for files(?)

	return nil
}
