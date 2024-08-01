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

// Cadastrando nova tarefa
func AddNewTask(c *gin.Context) {
	var newTask Tasks

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	newTask.Id = len(taskList) + 1
	taskList = append(taskList, newTask)
	c.JSON(http.StatusOK, newTask)

}
