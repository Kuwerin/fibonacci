package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Kuwerin/fibonacci/pkg/repository"
	"github.com/Kuwerin/fibonacci/pkg/service"
	grpctransport "github.com/Kuwerin/fibonacci/pkg/transport/grpc"
	"github.com/Kuwerin/fibonacci/pkg/transport/grpc/fibonaccipb"
	httptransport "github.com/Kuwerin/fibonacci/pkg/transport/http"
)

const (
	defaultHTTPPort  int = 5010
	defaultGRPCPort  int = 5000
	defaultRedisPort int = 6379

	defaultRedisHost string = "redis"
)

var (
	cfgHTTPPort  int
	cfgGRPCPort  int
	cfgRedisPort int

	cfgRedisHost string
)

var rootCmd = &cobra.Command{
	Use:   "svc-fibonacci",
	Short: "svc-fibonacci",
	Long:  "svc-fibonacci",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().IntVarP(&cfgHTTPPort, "http-port", "p", defaultHTTPPort, "http port to connect")
	rootCmd.PersistentFlags().IntVarP(&cfgGRPCPort, "grpc-port", "g", defaultGRPCPort, "grpc port to connect")
	rootCmd.PersistentFlags().IntVarP(&cfgRedisPort, "redis-port", "r", defaultRedisPort, "redis port to connect")

	rootCmd.PersistentFlags().StringVarP(&cfgRedisHost, "redis-host", "s", defaultRedisHost, "redis port to connect")

	viper.BindPFlag("http-port", rootCmd.PersistentFlags().Lookup("http-port"))   //nolint:errcheck
	viper.BindPFlag("grpc-port", rootCmd.PersistentFlags().Lookup("grpc-port"))   //nolint:errcheck
	viper.BindPFlag("redis-port", rootCmd.PersistentFlags().Lookup("redis-port")) //nolint:errcheck
	viper.BindPFlag("redis-host", rootCmd.PersistentFlags().Lookup("redis-host")) //nolint:errcheck
}

func initConfig() {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()
}

func run() {
	httpPort := viper.GetInt("http-port")
	redisHost := viper.GetString("redis-host")
	redisPort := viper.GetInt("redis-port")
	grpcPort := viper.GetInt("grpc-port")

	// Create the logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger.Log("app", os.Args[0], "event", "starting")
	}

	// Create the repository
	var rep *repository.Repository
	{
		log.With(logger, "repository", "redis")

		redisOpts := redis.Options{
			Addr: fmt.Sprintf("%s:%d", redisHost, redisPort),
		}

		redisClient := redis.NewClient(&redisOpts)

		rep = repository.MakeRepository(logger, redisClient)
	}

	service := service.NewService(rep)

	go func() {
		grpcServer := grpc.NewServer()

		reflection.Register(grpcServer)

		fibonaccipb.RegisterFibonacciServiceServer(grpcServer, grpctransport.NewGRPCServer(service))

		// Create gRPC-server
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			logger.Log("error", err)
		}

		if err := grpcServer.Serve(ln); err != nil {
			logger.Log("error", err)
		}
	}()

	srv := httptransport.NewHTTPServer(service)

	idleConnsClosed := make(chan struct{})

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		signal := <-quit

		logger.Log("os signal", signal)

		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}

		close(idleConnsClosed)
		logger.Log("event", "server stopped")
	}()

	logger.Log("event", "server started")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), srv); err != http.ErrServerClosed {
		logger.Log("error", err)
		os.Exit(1)
	}

	<-idleConnsClosed

	logger.Log("event", "server exited normally")
}
