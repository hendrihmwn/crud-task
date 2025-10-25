package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"reflect"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}
		// parse token, verify signature
		jwtSecret := InstanceHandler.config.JWTSecret
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}

func ValidationErrorHandler(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			message := FormatValidationError(err, obj)

			c.JSON(http.StatusBadRequest, gin.H{
				"message": message,
			})
			c.Abort()
			return
		}
	}
}

func FormatValidationError(err error, obj interface{}) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var messages []string

		for _, e := range errs {
			field, _ := reflect.TypeOf(obj).FieldByName(e.Field())
			jsonTag := field.Tag.Get("json")

			switch e.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("%s is required", jsonTag))
			case "max":
				messages = append(messages, fmt.Sprintf("%s must be at most %s characters", jsonTag, e.Param()))
			default:
				messages = append(messages, fmt.Sprintf("Invalid %s", jsonTag))
			}
		}

		return strings.Join(messages, ", ")
	}

	return err.Error()
}

func MockToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("user_id", "test-user")
		c.Next()
	}
}
