package middleware

// RoleAuth 是一个简单的角色鉴权中间件
//func RoleAuth(requiredRole string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		username, exists := c.Get("username")
//		if !exists {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
//			c.Abort()
//			return
//		}
//
//		验证用户角色
//		if !service.HasRole(username.(string), requiredRole) {
//			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
//			c.Abort()
//			return
//		}
//
//		c.Next()
//	}
//}
