package main

import (
	"section3/wallet/interfaces/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(cors.Default())
	http.Server(app)
}
