package utils;

import (
	"net/http"
	"cardealership/prisma/db"
	"github.com/gin-gonic/gin"
)

// TODO:  verficar se o prisma trabalha com Pool de Conexões
func GetPrisma(c *gin.Context) *db.PrismaClient {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	return client
}
