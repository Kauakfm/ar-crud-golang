package main

import (
	"github.com/ar-crud-golang/go-standard/controller"
	"github.com/ar-crud-golang/go-standard/db"
	"github.com/ar-crud-golang/go-standard/repository"
	"github.com/ar-crud-golang/go-standard/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de respository
	productRepository := repository.NewProductRepository(dbConnection)
	//Camada de useCase
	ProductUseCase := usecase.NewProductUseCase(productRepository)
	//Camada de controllers
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8000")
}
