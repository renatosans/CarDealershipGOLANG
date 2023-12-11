package handlers;

import (
	"net/http"
	"strconv"
	"cardealership/utils"
	"cardealership/prisma/db"
	"github.com/gin-gonic/gin"
)

func GetSalespeople(c *gin.Context) {
	// var salespeople []db.InnerSalesperson

	client := utils.GetPrisma(c)

	salesPeople, err := client.Salesperson.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": salesPeople})
}

func PostSalesperson(c *gin.Context) {
	var payload db.InnerSalesperson

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	insertedSp, err := client.Salesperson.CreateOne(
		db.Salesperson.FirstName.Set(payload.FirstName),
		db.Salesperson.Commission.Set(payload.Commission),
		db.Salesperson.LastName.SetOptional(payload.LastName),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Salesperson created successfully", "salesperson": insertedSp})
}

func PatchSalesperson(c *gin.Context) {
	var payload db.InnerSalesperson

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
	updatedSp, err := client.Salesperson.FindUnique(
		db.Salesperson.ID.Equals(id),
	).Update(
		db.Salesperson.FirstName.Set(payload.FirstName),
		db.Salesperson.Commission.Set(payload.Commission),
		db.Salesperson.LastName.SetOptional(payload.LastName),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Salesperson patched", "salesperson": updatedSp})
}

func DeleteSalesperson(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := utils.GetPrisma(c)
	deletedSp, err := client.Salesperson.FindUnique(
		db.Salesperson.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Salesperson deleted successfully", "id": deletedSp.ID})
}
