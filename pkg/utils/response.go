package utils

import "github.com/gin-gonic/gin"

// RespondJSON success response bhejne ka standard tareeqa
func RespondJSON(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, data)
}

// RespondError error response bhejne ka standard tareeqa
func RespondError(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, gin.H{"error": message})
}