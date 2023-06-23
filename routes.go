package main

import (
	"github.com/gin-gonic/gin"
)

func Routes(Router *gin.Engine) {

	/**
	* We have to define ping route to get health information about
	* the Zend API to intergrate with monitoring utilities.
	 */
	Router.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(200, gin.H{
			"status": "success",
		})
	})

	/*
	* We have to show resource not found error if some
	* application request undefined route.
	 */
	Router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{
			"status": "failed",
			"error": gin.H{
				"code":    404,
				"message": "Resource not found",
			},
		})
	})

}
