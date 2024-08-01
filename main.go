package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tasks struct {
	Id    int    `json:"id`
	Title string `json:"title`
}

var taskList = []Tasks{
	{Id: 1, Title: "Task 1"},
	{Id: 2, Title: "Task 2"},
}

func main() {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	//Rota de teste raiz
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Minha primeira API funcionando!",
		})
	})

	//Buscando todas as tarefas
	router.GET("/tarefas", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, taskList)
	})

	//Cadastrando nova tarefa
	router.POST("/tarefas", func(c *gin.Context) {
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

	})

	router.Run(":3001")
}
