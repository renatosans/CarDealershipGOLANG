package handlers;

import (
	"net/http"
	"strconv"
	"cardealership/utils"
	"cardealership/prisma/db"
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

func PostCar(c *gin.Context) {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car created successfully", "car": insertedCar})
}

func PatchCar(c *gin.Context) {
	var payload db.InnerCarsForSale

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	updatedCar, err := client.CarsForSale.FindUnique(
		db.CarsForSale.ID.Equals(id),
	).Update(
		db.CarsForSale.Brand.Set(payload.Brand),
		db.CarsForSale.Model.Set(payload.Model),
		db.CarsForSale.Year.Set(payload.Year),
		db.CarsForSale.Price.Set(payload.Price),
		db.CarsForSale.Color.SetOptional(payload.Color),
		db.CarsForSale.Mileage.SetOptional(payload.Mileage),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car patched", "car": updatedCar})
}

func DeleteCar(c *gin.Context) {
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
