package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os/signal"
	"syscall"
	"time"

	"godgifu/config"
	"godgifu/modules/account"
	"godgifu/modules/auth"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	router, err := config.DevLoadConfig()
	if err != nil {
		log.Printf("LoadConfig failed, error: %v", err)
	}

	auth.InitAuth(router.Router, router.Postgres, router.Redis, router.JWT.PrivateKey, router.JWT.PublicKey, router.JWT.RefreshSecretKey, router.JWT.RefreshTokenExpirationSecs, router.JWT.IDTokenExpirationSecs)
	account.InitAccount(router.Router, router.Postgres)

	log.Println("Server port:", router.Port)
	log.Println("Postgres connection:", router.Postgres)
	log.Println("Redis connection:", router.Redis)

	// context listens for the server kill cmd ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// server := http.Server{
	// 	Addr:    fmt.Sprintf(":%s", router.Port),
	// 	Handler: router.Router,
	// }

	go func() {
		// if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	log.Fatalf("error with ListenAndServe: %v\n", err)
		// }
		if err := router.Router.Start(":" + router.Port); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// pprof profile and analysis
	go func() {
		// go to: http://localhost:6060/debug/pprof/
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// -----Gracefull Shutdown-----
	<-ctx.Done()
	stop()
	log.Println("Graceful Shutdown initiated, press Ctrl+C again to force Shutdown")

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// close data storage connections here
	if err := router.CloseDataStorageConnections(); err != nil {
		log.Fatalf("Possible error or Graceful Shutdown initiated. Closing data storage connections %v\n", err)
	}

	if err := router.Router.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown due to: ", err)
	}

	log.Print("Server exiting...")
}
