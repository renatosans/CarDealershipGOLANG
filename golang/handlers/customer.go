package handlers;

import (
	"net/http"
	"strconv"
	"cardealership/utils"
	"cardealership/prisma/db"
	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	// var customers []db.InnerCustomer

	client := utils.GetPrisma(c)

	customers, err := client.Customer.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func PostCustomer(c *gin.Context) {
	var payload db.InnerCustomer

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	insertedCustomer, err := client.Customer.CreateOne(
		db.Customer.FirstName.Set(payload.FirstName),
		db.Customer.LastName.Set(payload.LastName),
		db.Customer.BirthDate.Set(payload.BirthDate),
		db.Customer.Email.SetOptional(payload.Email),
		db.Customer.Phone.SetOptional(payload.Phone),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer created successfully", "customer": insertedCustomer})
}

func PatchCustomer(c *gin.Context) {
	var payload db.InnerCustomer

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
	updatedCustomer, err := client.Customer.FindUnique(
		db.Customer.ID.Equals(id),
	).Update(
		db.Customer.FirstName.Set(payload.FirstName),
		db.Customer.LastName.Set(payload.LastName),
		db.Customer.BirthDate.Set(payload.BirthDate),
		db.Customer.Email.SetOptional(payload.Email),
		db.Customer.Phone.SetOptional(payload.Phone),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer patched", "customer": updatedCustomer})
}

func DeleteCustomer(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := utils.GetPrisma(c)
	deletedCustomer, err := client.Customer.FindUnique(
		db.Customer.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully", "id": deletedCustomer.ID})
}
