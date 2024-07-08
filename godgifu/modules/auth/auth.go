package auth

import (
	"godgifu/config"
	"godgifu/modules/auth/db"
	"godgifu/modules/auth/handlers"
	"godgifu/modules/auth/routes"
	"godgifu/modules/auth/services"
)

type Auth struct {
	AuthHandlers handlers.AuthHandlers
	AuthServices services.AuthService
	DB           db.PostgresDB
	JWTHandlers  handlers.JWTHandler
	JWTServices  services.JWTService
	Redis        db.RedisDB
}

func InitAuth(server *config.DevConfiguration) Auth {
	postgresDB := db.NewPostgresDB(server.Postgres)
	redisDB := db.NewRedisTokenRepository(server.Redis)
	auth := services.NewAuthServices(postgresDB)
	jwt := services.NewJWTService(&services.ConfigTokenService{
		TokenRepository:            redisDB,
		PrivateKey:                 server.JWT.PrivateKey,
		PublicKey:                  server.JWT.PublicKey,
		RefreshSecretKey:           server.JWT.RefreshSecretKey,
		IDTokenExpirationSecs:      server.JWT.IDTokenExpirationSecs,
		RefreshTokenExpirationSecs: server.JWT.RefreshTokenExpirationSecs,
	})
	authHandlers := handlers.NewAuthHandlers(auth, jwt)
	jwtHandlers := handlers.NewJWTHandlers(auth, jwt)
	routes.InitRoutes(server.Router, authHandlers, jwtHandlers, jwt)

	return Auth{
		AuthHandlers: authHandlers,
		AuthServices: auth,
		DB:           postgresDB,
		JWTHandlers:  jwtHandlers,
		JWTServices:  jwt,
		Redis:        redisDB,
	}
}
