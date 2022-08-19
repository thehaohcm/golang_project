package pkg

import "github.com/gin-gonic/gin"

// ErrorResponse: handle error when creating a response
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
