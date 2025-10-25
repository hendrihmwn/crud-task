package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/crud-task-backend/model"
	"net/http"
)

func registerAuthHandler(route *gin.Engine) {
	route.POST("/login", ValidationErrorHandler(model.LoginParam{}), InstanceHandler.login)
}

func (i MainInstance) login(c *gin.Context) {
	var param model.LoginParam

	err := c.ShouldBind(&param)
	if err != nil {
		errMessage := FormatValidationError(err, param)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	res, err := i.authUseCase.Login(c, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
