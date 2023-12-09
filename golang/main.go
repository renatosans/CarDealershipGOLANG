package main

import (
	"strconv"
	"net/http"
	"cardealership/prisma/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetPrisma(c *gin.Context) *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return client
}

func getCars(c *gin.Context) {
	// var pets []db.InnerPet

	client := GetPrisma(c)

	pets, err := client.CarsForSale.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func deleteCar(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := GetPrisma(c)
	deletedCar, err := client.CarsForSale.FindUnique(
		db.CarsForSale.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully", "car id": deletedCar.ID})
}

func main() {
	godotenv.Load(".env")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	rGroup.GET("/cars", getCars)
	// rGroup.POST("/cars", postCar)
	// rGroup.PATCH("/cars/:id", patchCar)
	rGroup.DELETE("/cars/:id", deleteCar)

	router.Use(cors.Default())
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
