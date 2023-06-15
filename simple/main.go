package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rrojan/enforcer"
)

func main() {
	r := gin.Default()

	r.POST("/signup", func(c *gin.Context) {
		type User struct {
			Name         string    `enforce:"required between:2,32 wordCount:1,4"`
			Password     string    `enforce:"required match:password"`
			UserType     string    `enforce:"default:user enum:user,admin"`
			CreatedAt    time.Time `enforce:"default:timeNow"`
			VerifyBefore time.Time `enfore:"default:timeNow+3_days"`
		}
		// Create and bind request data to User instance
		u := User{}
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(400, gin.H{"error": "Error binding data"})
			return
		}

		// Validate and set default values for this user
		errors := enforcer.Validate(&u)

		if len(errors) > 0 {
			c.JSON(400, gin.H{"errors": errors})
			return
		}
		c.JSON(200, gin.H{"message": "Success!", "user": u})
	})
	r.Run(":6969")
}
