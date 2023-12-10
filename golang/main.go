package main

import (
	"cardealership/utils"
	"cardealership/handlers"
	"cardealership/prisma/db"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func getCars(c *gin.Context) {
	// var cars []db.InnerCarsForSale

	client := utils.GetPrisma(c)

	pets, err := client.CarsForSale.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func postCar(c *gin.Context) {
	var payload db.InnerCarsForSale

	// Bind JSON body to the Pet struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	insertedCar, err := client.CarsForSale.CreateOne(
		db.CarsForSale.Brand.Set(payload.Brand),
		db.CarsForSale.Model.Set(payload.Model),
		db.CarsForSale.Year.Set(payload.Year),
		db.CarsForSale.Price.Set(payload.Price),
		db.CarsForSale.Color.SetOptional(payload.Color),
		db.CarsForSale.Mileage.SetOptional(payload.Mileage),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car created successfully", "car": insertedCar})
}

func deleteCar(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := utils.GetPrisma(c)
	deletedCar, err := client.CarsForSale.FindUnique(
		db.CarsForSale.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully", "car id": deletedCar.ID})
}

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

func main() {
	// Variáveis de ambiente passadas no docker compose,  remover o comentário caso necessite
	// godotenv.Load(".env")

	router := gin.Default()
	router.Use(cors.Default())    // CORS - Default() allows all origins

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	rGroup.GET("/cars", getCars)
	rGroup.POST("/cars", postCar)
	// rGroup.PATCH("/cars/:id", patchCar)
	rGroup.DELETE("/cars/:id", deleteCar)

	addCustomerHandlers(rGroup)
	addSalespersonHandlers(rGroup)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
