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
		db.Invoice.Amount.Set(payload.Amount),
		db.Invoice.CustomerID.SetIfPresent(&payload.CustomerID),
		db.Invoice.SalespersonID.SetIfPresent(&payload.SalespersonID),
		db.Invoice.CarID.SetIfPresent(&payload.CarID),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice created successfully", "invoice": insertedInvoice})
}

func PatchInvoice(c *gin.Context) {
	var payload db.InnerInvoice

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
	updatedInvoice, err := client.Invoice.FindUnique(
		db.Invoice.ID.Equals(id),
	).Update(
		db.Invoice.Amount.Set(payload.Amount),
		db.Invoice.CustomerID.SetIfPresent(&payload.CustomerID),
		db.Invoice.SalespersonID.SetIfPresent(&payload.SalespersonID),
		db.Invoice.CarID.SetIfPresent(&payload.CarID),
	).Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice patched", "invoice": updatedInvoice})
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
	deletedInvoice, err := client.Invoice.FindUnique(
		db.Invoice.ID.Equals(id),
	).Delete().Exec(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully", "id": deletedInvoice.ID})
}
