package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Incluído para ser utilizado pelo BD
type Tasks struct {
	Id    int    `json:"id`
	Title string `json:"title`
}

var taskList = []Tasks{
	{Id: 1, Title: "Task 1"},
	{Id: 2, Title: "Task 2"},
}

// Rota de teste raiz
func RoutTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Minha primeira API funcionando!",
	})
}

// Buscando todas as tarefas
/*
func GetAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, taskList)

}
*/

func GetAllTasks(c *gin.Context) {
	rows, err := DB.Query("SELECT id, title FROM tasks")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	defer rows.Close()

	var tasks []Tasks

	for rows.Next() {
		var task Tasks
		if err := rows.Scan(&task.Id, &task.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		tasks = append(tasks, task)

	}

	c.JSON(http.StatusOK, tasks)

}

// Cadastrando nova tarefa
/*
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
*/

func AddNewTask(c *gin.Context) {
	var newTask Tasks

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := DB.Exec("INSERT INTO tasks (title) VALUES (?)", newTask.Title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTask.Id = int(id)
	c.JSON(http.StatusCreated, newTask)

}

// Buscando tarefa por Id
/*
func GetTaskById(c *gin.Context) {
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

}
*/
func GetTaskById(c *gin.Context) {
	id := c.Param("id")

	var task Tasks

	row := DB.QueryRow("SELECT id, title FROM tasks WHERE id = ?", id)

	if err := row.Scan(&task.Id, &task.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, task)

}

// Deletando tarefa por Id
func DeleteTaskById(c *gin.Context) {

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

}

// Editando tarefa por Id
func UpdateTaskById(c *gin.Context) {

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

}
