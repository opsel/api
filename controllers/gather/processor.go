package gather

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

func Processor(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	/**
	* We need to bind JSON data values into struct
	* before start processing them with the application
	* logic.
	 */
	var Request struct {
		Cores []struct {
			ID    uint64   `json:"id" binding:"required"`
			Name  string   `json:"name" binding:"required"`
			MHz   float64  `json:"mhz" binding:"required"`
			Flags []string `json:"flags" binding:"required"`
		} `json:"cores" binding:"required"`
	}
	if err := c.ShouldBindBodyWith(&Request, binding.JSON); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": "failed",
			"error": gin.H{
				"ref": "malformed",
			},
		})
		return
	}

	// fmt.Println(Request)
	fmt.Println(db)

	c.AbortWithStatusJSON(200, gin.H{
		"status": "success",
	})
}
