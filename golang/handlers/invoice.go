package handlers;

import (
	"net/http"
	"cardealership/utils"
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

// func PostInvoice(c *gin.Context) {
// ...
// ...
