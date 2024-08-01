package main

import (
	"fmt"
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

func RegisterRoutes(router *gin.Engine) {

	//Rota de teste raiz
	router.GET("/", RoutTest)

	//Buscando todas as tarefas
	router.GET("/tarefas", GetAllTasks)

	//Cadastrando nova tarefa
	router.POST("/tarefas", AddNewTask)

	//Buscando tarefa por Id
	router.GET("/tarefas/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, task := range taskList {
			if fmt.Sprintf("%d", task.Id) == id {
				c.JSON(http.StatusOK, task)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Tarefa com id=" + string(c.Param("id")) + " não encontrada!",
		})

	})

	//Deletando tarefa por Id
	router.DELETE("/tarefas/:id", func(c *gin.Context) {
		id := c.Param("id")

		for index, task := range taskList {
			if fmt.Sprintf("%d", task.Id) == id {
				taskList = append(taskList[:index], taskList[index+1:]...)
				c.JSON(http.StatusOK, gin.H{
					"message": "Tarefa com id=" + string(c.Param("id")) + " deletada com sucesso!"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Tarefa com id=" + string(c.Param("id")) + " não encontrada!",
		})

	})

	//Editando tarefa por Id
	router.PUT("/tarefas/:id", func(c *gin.Context) {
		id := c.Param("id")

		var updateTask Tasks

		if err := c.BindJSON(&updateTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Tarefa com id=" + string(c.Param("id")) + " não encontrada!",
			})
			return

		}

		for index, task := range taskList {
			if fmt.Sprintf("%d", task.Id) == id {
				taskList[index] = updateTask
				c.JSON(http.StatusOK, gin.H{
					"message": "Tarefa com id=" + string(c.Param("id")) + " atualizada com sucesso!"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Tarefa com id=" + string(c.Param("id")) + " não encontrada!",
		})

	})

}
