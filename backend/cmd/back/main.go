package main

import (
	"context"
	"flag"
	"loan-back-services/cmd/back/handler"
	"loan-back-services/pkg/ds"
	"loan-back-services/pkg/logger"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// for back services
func main() {
	port := flag.String("port", "8900", "back default port is 8900")
	flag.Parse()

	addr := net.JoinHostPort("", *port)

	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	ds, err := ds.NewDataSource()
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	h := handler.NewHandler(&handler.HConfig{
		R:  router,
		DS: ds,
	})

	h.Register()
	defer h.Destroy()

	server := http.Server{
		Addr:           addr,
		Handler:        h.R,
		ReadTimeout:    time.Duration(time.Minute * 3),
		WriteTimeout:   time.Duration(time.Minute * 3),
		MaxHeaderBytes: 10 << 20, // 10MB
	}

	go func() {
		logger.Sugar.Info("Server started listening on port : ", *port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Sugar.Fatal("Server Failed to initialed on port : ", *port)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	// shutdown close
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Sugar.Error("Failed to shutdown server: ", err.Error())
	}
}
