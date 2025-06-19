// @title Product Trace API
// @version 1.0
// @description API for managing products on blockchain with QR code support
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:5000
// @BasePath /api
// @schemes http

package main

import (
	//"fmt"
	"log"
	"os"
	"product_trace/config"
	"product_trace/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "product_trace/docs"
)

func main() {
	wd, _ := os.Getwd()
	log.Println("Current working directory:", wd)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Không thể load file .env: ", err)
	}
	// fmt.Println("INFURA_KEY:",os.Getenv("INFURA_KEY"))
	// fmt.Println("CONTRACT_ADDRESS:",os.Getenv("CONTRACT_ADDRESS"))
	// fmt.Println("PRIVATE_KEY:",os.Getenv("PRIVATE_KEY"))

	config.ConnectDB()
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/product/add", handlers.AddProduct)
	r.POST("/api/product/:id/status", handlers.UpdateProductStatus)
	r.GET("/api/product/:id", handlers.GetProduct)
	r.GET("/api/products", handlers.GetProducts)

	r.Run(":5000")
}