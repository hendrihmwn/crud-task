package handler

import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerTaskHandler(route *gin.Engine) {
	task := route.Group("/tasks", AuthMiddleware())
	task.GET("", InstanceHandler.listTask)
	task.GET("/:id", InstanceHandler.getTask)
	task.POST("", InstanceHandler.createTask)
	task.PUT("/:id", InstanceHandler.updateTask)
	task.DELETE("/:id", InstanceHandler.deleteTask)
}

func (i MainInstance) listTask(c *gin.Context) {
	var param model.TaskListParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, size, err := i.taskUseCase.ListTask(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": model.TaskMeta{
			Limit: int(param.Limit),
			Page:  int(param.Page),
			Total: size,
		},
		"data": res,
	})
}

func (i MainInstance) getTask(c *gin.Context) {
	var param model.TaskGetParam
	err := c.ShouldBindUri(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := i.taskUseCase.GetTask(c, param.ID)
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (i MainInstance) createTask(c *gin.Context) {
	var body model.TaskBodyParam
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := i.taskUseCase.CreateTask(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func (i MainInstance) updateTask(c *gin.Context) {
	var param model.TaskGetParam
	err := c.ShouldBindUri(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body model.TaskBodyParam
	err = c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := i.taskUseCase.UpdateTask(c, param.ID, body)
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (i MainInstance) deleteTask(c *gin.Context) {
	var param model.TaskGetParam
	err := c.ShouldBindUri(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = i.taskUseCase.DeleteTask(c, param.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
