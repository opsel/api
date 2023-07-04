package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Database() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := gorm.Open(postgres.New(postgres.Config{
			DSN:                  "host=192.168.2.50 user=opsel_db password=p7rdSdGsAt3uqJVWotkNSGCk4VbUQMdh dbname=opsel_db port=5432 sslmode=disable TimeZone=Asia/Colombo",
			PreferSimpleProtocol: true,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		/**
		* Configure SQL connection pool settings to avoid any
		* funky things from happening with connection limites
		 */
		connection, _ := db.DB()
		connection.SetConnMaxLifetime(time.Second)

		c.Set("db", db)
		c.Next()
	}
}
