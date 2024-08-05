package main

import (
	"github.com/gin-gonic/gin"
)

// Suprimido para consumir no BD SQLite
/*

type Tasks struct {
	Id    int    `json:"id`
	Title string `json:"title`
}

var taskList = []Tasks{
	{Id: 1, Title: "Task 1"},
	{Id: 2, Title: "Task 2"},
}

*/

func RegisterRoutes(router *gin.Engine) {

	//Rota de teste raiz
	router.GET("/", RoutTest)

	//Buscando todas as tarefas
	router.GET("/tarefas", GetAllTasks)

	//Cadastrando nova tarefa
	router.POST("/tarefas", AddNewTask)

	//Buscando tarefa por Id
	router.GET("/tarefas/:id", GetTaskById)

	//Deletando tarefa por Id
	router.DELETE("/tarefas/:id", DeleteTaskById)

	//Editando tarefa por Id
	router.PUT("/tarefas/:id", UpdateTaskById)

}
