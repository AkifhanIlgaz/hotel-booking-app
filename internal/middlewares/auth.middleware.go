package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/token"
	"github.com/gin-gonic/gin"
)

const bearerPrefix = "Bearer "

type AuthMiddleware struct {
	tokenManager *token.Manager
}

func NewAuthMiddleware(tokenManager *token.Manager) *AuthMiddleware {
	return &AuthMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *AuthMiddleware) AccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return

		}

		// 2. Check if the header starts with "Bearer "
		//    The expected format is: "Bearer <token>"
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			// return "", errors.New("authorization header format must be Bearer {token}")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		// 3. Extract the token part by trimming the prefix
		token := strings.TrimSpace(authHeader[len(bearerPrefix):])
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
			// return "", errors.New("authorization token is empty")
		}

		claims, err := m.tokenManager.ParseAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("uid", claims.Subject)
		c.Set("role", claims.Role)

		fmt.Println(c.Get("uid"))
		fmt.Println(c.Get("role"))

		c.Next()
	}
}
