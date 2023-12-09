package main

import (
	"cardealership/prisma/db"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func GetPrisma(c *gin.Context) *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	return client
}

func main() {
	// godotenv.Load(".env")

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin gonic API",
		})
	})

	rGroup := router.Group("/api")
	rGroup.GET("/pets", getCars)
	rGroup.POST("/pets", postCar)
	rGroup.PATCH("/pets/:id", patchCar)
	rGroup.DELETE("/pets/:id", deleteCar)

	router.Use(cors.Default())
	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
