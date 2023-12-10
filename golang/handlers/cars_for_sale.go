package handlers;

import (
	"net/http"
	"cardealership/utils"
	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	// var cars []db.InnerCarsForSale

	client := utils.GetPrisma(c)

	pets, err := client.CarsForSale.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

// func PostCar(c *gin.Context) {
// ...
// ...
