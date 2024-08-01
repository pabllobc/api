package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Rota de teste raiz
func RoutTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Minha primeira API funcionando!",
	})
}

// Buscando todas as tarefas
func GetAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, taskList)
}
