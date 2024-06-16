package auth

import (
	"crypto/rsa"
	"godgifu/modules/auth/db"
	"godgifu/modules/auth/handlers"
	"godgifu/modules/auth/routes"
	"godgifu/modules/auth/services"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type Auth struct {
}

func InitAuth(server *echo.Echo, database *sqlx.DB, redis *redis.Client, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, refreshSecretKey string, idTokenExpiration int64, refreshTokenExpiration int64) {
	postgresDB := db.NewPostgresDB(database)
	redisDB := db.NewRedisTokenRepository(redis)
	auth := services.NewAuthServices(postgresDB)
	jwt := services.NewJWTService(&services.ConfigTokenService{
		TokenRepository:            redisDB,
		PrivateKey:                 privateKey,
		PublicKey:                  publicKey,
		RefreshSecretKey:           refreshSecretKey,
		IDTokenExpirationSecs:      idTokenExpiration,
		RefreshTokenExpirationSecs: refreshTokenExpiration,
	})
	handlers := handlers.NewAuthHandlers(auth, jwt)
	routes.InitRoutes(server, handlers)
}
