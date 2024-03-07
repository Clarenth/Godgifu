package config

// import (
// 	"log"
// 	"time"

// 	// auth_handlers "godgifu-server/modules/auth/handlers"
// 	// auth_repository "godgifu-server/modules/auth/repository"
// 	// auth_services "godgifu-server/modules/auth/services"
// 	"godgifu-server/modules/auth/routes"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type server struct {
// 	Router          *gin.Engine
// 	BaseURL         string
// 	TimeoutDuration time.Duration
// 	Postgres *pgxpool.Config
// 	//NoSQL *sqlx.DB

// 	//DocumentDirectory string
// }

// func Injection() *gin.Engine {
// 	config, err := LoadConfig()
// 	if err != nil {
// 		log.Print("Error loading configuration")
// 	}

// 	router := gin.New()
// 	router.Use(gin.Recovery(), gin.Logger())

// 	configServer := &server{
// 		Router: router,
// 		BaseURL: "localhost",
// 		TimeoutDuration: 5 * time.Second,
// 	}

// 	auth := routes.InitRoutes(config.Postgres.DB, router)

// 	panic("not done yet")
// }
