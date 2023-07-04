package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Gather() gin.HandlerFunc {
	return func(c *gin.Context) {

		/**
		* Check if the agent id header is present in the request
		* header and validate it for valid UUIDv4 before proceed
		* into the next handler function
		 */
		AgentID := c.GetHeader("opsel-agent-id")
		uuid, err := uuid.Parse(AgentID)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(400, gin.H{
				"status": "failed",
				"error": gin.H{
					"ref": "invalid_agent_id",
				},
			})
			return
		}

		/**
		* We have to pass the AgentID information to the next
		* handler function of the application to identify the
		* instance that data are coming from
		 */
		c.Set("agent-id", uuid)
		c.Next()

	}
}
