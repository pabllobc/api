package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutTest(c *gin.Context) {
	//Rota de teste raiz
	c.JSON(http.StatusOK, gin.H{
		"message": "Minha primeira API funcionando!",
	})
}
