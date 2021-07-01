package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kuwerin/fibonacci/internal/server"
	"github.com/Kuwerin/fibonacci/internal/service"
	"github.com/Kuwerin/fibonacci/pkg/cache"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")
	return viper.ReadInConfig()
}

func Run() {
	if err := initConfig(); err != nil {
		log.Fatalf("an error occured while trying to initialize configs: %s", err.Error())
	}

	ctx := context.Background()

	rdbConfig := cache.NewRedisConfig(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASS"), 0)
	rdbClient := rdbConfig.Connect()
	_, err := rdbClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("an error occured while trying to connect to redis: %s", err.Error())
	}

	rdbCache := cache.NewRedisStore(rdbClient, ctx)
	services := service.NewService(rdbCache)
	srv := server.NewHttpServer(services)
	grpcsrv := server.NewGRPCServer(services)
	go func() {
		if err := http.ListenAndServe(viper.GetString("http.port"), srv); err != nil {
			log.Fatalf("an error occured while trying to run http server: %s", err.Error())
		}
	}()

	go func() {
		if err := grpcsrv.Run(services); err != nil {
			log.Fatalf("an error occured while trying to run gprc server: %s", err.Error())
		}
	}()
	log.Print("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("app shutting down")
	grpcsrv.Shutdown()

}
