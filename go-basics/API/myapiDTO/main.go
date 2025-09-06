// dto.go
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks []Task
var lastTaskID = 0

func main() {
	r := gin.Default()

	r.GET("/tasks", getTasksHandler)
	r.POST("/tasks", createTaskHandler)
	r.PUT("/tasks/:id", updateTaskHandler)

	r.Run(":8080")
}

func getTasksHandler(c *gin.Context) {
	taskDTOs := make([]TaskDTO, len(tasks))

	for i, task := range tasks {
		taskDTOs[i] = TaskDTO{
			ID:    task.ID,
			Title: task.Title,
			Done:  task.Done,
		}
	}

	c.JSON(http.StatusOK, taskDTOs)
}

func createTaskHandler(c *gin.Context) {
	var taskDTO TaskDTO
	if err := c.ShouldBindJSON(&taskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lastTaskID++
	newTask := Task{
		ID:    lastTaskID,
		Title: taskDTO.Title,
		Done:  taskDTO.Done,
	}

	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, taskDTO)
}

func updateTaskHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var taskDTO TaskDTO
	if err := c.ShouldBindJSON(&taskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = taskDTO.Title
			tasks[i].Done = taskDTO.Done

			c.JSON(http.StatusOK, taskDTO)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
