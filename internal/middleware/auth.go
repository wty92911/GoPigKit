package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wty92911/GoPigKit/internal/dao"
	"net/http"
	"strings"
)

type Claims struct {
	OpenID string `json:"openid"`
	jwt.RegisteredClaims
}

// AuthToken Token 验证中间件
func AuthToken(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("openid", claims.OpenID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AuthFamily 验证是否加入家庭，
func AuthFamily(mustOwner bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		openid := c.GetString("openid")
		if openid == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		user, err := dao.GetUser(openid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if user.FamilyID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not in family"})
			c.Abort()
			return
		}

		if mustOwner {
			family, err := dao.GetFamily(user.FamilyID)
			if err != nil || family.OwnerOpenID != user.OpenID {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": fmt.Sprintf("User is not owner of family %d", family.ID),
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
