package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	adminUsername = "siful"
	adminPassword = "admin123"
)

// EncodeBasicAuth encodes username and password as Basic Auth header
func EncodeBasicAuth(username, password string) string {
	credentials := fmt.Sprintf("%s:%s", username, password)
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	return fmt.Sprintf("Basic %s", encoded)
}

// AdminAuthMiddleware validates admin credentials from Basic Auth header
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Parse Basic Auth header
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Decode base64
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Parse username:password
		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		username := credentials[0]
		password := credentials[1]

		// Validate credentials
		if username != adminUsername || password != adminPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set admin context
		c.Set("admin", true)
		c.Next()
	}
}
