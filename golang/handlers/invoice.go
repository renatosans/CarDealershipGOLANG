package handlers

import (
	"cardealership/prisma/db"
	"cardealership/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInvoices(c *gin.Context) {
	// var invoices []db.InnerInvoice

	client := utils.GetPrisma(c)

	invoices, err := client.Invoice.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoices})
}

func PostInvoice(c *gin.Context) {
	var payload db.InnerInvoice

	// Bind JSON body to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	insertedInvoice, err := client.Invoice.CreateOne(
		db.
			db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.FlagRemoved.Set(payload.FlagRemoved),
		db.Pet.Age.SetOptional(payload.Age),
		db.Pet.Owner.SetOptional(payload.Owner),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet created successfully", "pet": insertedPet})
}

func PatchInvoice(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Bind JSON body to the Pet struct
	var payload db.InnerPet
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := utils.GetPrisma(c)
	updatedPet, err := client.Pet.FindUnique(
		db.Pet.ID.Equals(id),
	).Update(
		db.Pet.Name.Set(payload.Name),
		db.Pet.Breed.Set(payload.Breed),
		db.Pet.FlagRemoved.Set(payload.FlagRemoved),
		db.Pet.Age.SetOptional(payload.Age),
		db.Pet.Owner.SetOptional(payload.Owner),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet patched", "pet": updatedPet})
}

func DeleteInvoice(c *gin.Context) {
	// TODO: utilizar o flag_removed ao invés de apagar o registro na tabela

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client := utils.GetPrisma(c)
	deletedPet, err := client.Pet.FindUnique(
		db.Pet.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully", "pet id": deletedPet.ID})
}
