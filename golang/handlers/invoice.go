package handlers

import (
	"cardealership/prisma/db"
	"cardealership/utils"
	"net/http"
	"strings"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/signintech/gopdf"
)

func GetInvoices(c *gin.Context) {
	// var invoices []db.InnerInvoice

	client := utils.GetPrisma(c)

	invoices, err := client.Invoice.FindMany().Exec(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully", "id": deletedInvoice.ID})
}

func GerarPedido(c *gin.Context) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.SetMargins(40, 40, 40, 40)
	pdf.AddPage()
	err := pdf.AddTTFFontData("Inter", interFont)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = pdf.AddTTFFontData("Inter-Bold", interBoldFont)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.WriteLogo(&pdf, file.Logo, file.From)
	utils.WriteTitle(&pdf, file.Title, file.Id, file.Date)
	utils.WriteBillTo(&pdf, file.To)
	utils.WriteHeaderRow(&pdf)
	subtotal := 0.0
	for i := range file.Items {
		q := 1
		if len(file.Quantities) > i {
			q = file.Quantities[i]
		}

		r := 0.0
		if len(file.Rates) > i {
			r = file.Rates[i]
		}

		utils.WriteRow(&pdf, file.Items[i], q, r)
		subtotal += float64(q) * r
	}
	if file.Note != "" {
		writeNotes(&pdf, file.Note)
	}
	utils.WriteTotals(&pdf, subtotal, subtotal*file.Tax, subtotal*file.Discount)
	if file.Due != "" {
		writeDueDate(&pdf, file.Due)
	}
	utils.WriteFooter(&pdf, file.Id)
	output = strings.TrimSuffix(output, ".pdf") + ".pdf"
	err = pdf.WritePdf(output)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
