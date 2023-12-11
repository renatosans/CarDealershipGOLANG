package main

import (
	"net/http"
	"cardealership/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	// "github.com/joho/godotenv"
)

func addCustomerHandlers(rGroup *gin.RouterGroup) {
	rGroup.GET("/customers", handlers.GetCustomers)
	rGroup.POST("/customers", handlers.PostCustomer)
	rGroup.PATCH("/customers/:id", handlers.PatchCustomer)
	rGroup.DELETE("/customers/:id", handlers.DeleteCustomer)
}

func addSalespersonHandlers(rGroup *gin.RouterGroup) {
	rGroup.GET("/salespeople", handlers.GetSalespeople)
	rGroup.POST("/salespeople", handlers.PostSalesperson)
	rGroup.PATCH("/salespeople/:id", handlers.PatchSalesperson)
	rGroup.DELETE("/salespeople/:id", handlers.DeleteSalesperson)
}

func addInvoiceHandlers(rGroup *gin.RouterGroup) {
	rGroup.GET("invoices", handlers.GetInvoices)
	rGroup.POST("invoices", handlers.PostInvoice)
	rGroup.PATCH("invoices/:id", handlers.PatchInvoice)
	rGroup.DELETE("invoices/:id", handlers.DeleteInvoice)
}

func main() {
	// Variáveis de ambiente passadas no docker compose,  remover o comentário caso necessite
	// godotenv.Load(".env")

	router := gin.Default()
	router.Use(cors.Default()) // CORS - Default() allows all origins

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	rGroup.GET("/cars", handlers.GetCars)
	rGroup.POST("/cars", handlers.PostCar)
	rGroup.PATCH("/cars/:id", handlers.PatchCar)
	rGroup.DELETE("/cars/:id", handlers.DeleteCar)

	addCustomerHandlers(rGroup)
	addSalespersonHandlers(rGroup)
	addInvoiceHandlers(rGroup)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
