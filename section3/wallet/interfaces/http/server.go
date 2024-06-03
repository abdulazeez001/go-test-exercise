package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"section3/wallet/config"
	"section3/wallet/interfaces/http/router"

	"github.com/gin-gonic/gin"
)

// HTTP Server function call
func Server(app *gin.Engine) {
	// const headerTimeout = 10
	router.Version1RouteEntry(app)

	port := config.Get["app"].HTTPPort

	serve := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", "0.0.0.0", strconv.Itoa(port)),
		Handler:      app,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("[pid %d] REST server Listening on port %s", os.Getpid(), strconv.Itoa(port))
	// service connections
	if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
