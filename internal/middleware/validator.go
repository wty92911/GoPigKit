package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

//	func InitValidator() error {
//		validate = validator.New()
//
//		err := errors.Join(
//			validate.RegisterValidation("role", func(fl validator.FieldLevel) bool {
//				role := fl.Field().String()
//				return role == "cook" || role == "diner"
//			}),
//		)
//		if err != nil {
//			return err
//		}
//
// }
func ValidateStruct(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(obj); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := validate.Struct(obj); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.StructNamespace()+" "+err.Tag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			c.Abort()
			return
		}

		c.Set("validatedData", obj)
		c.Next()
	}
}
