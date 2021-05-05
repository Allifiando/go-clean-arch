package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS ...
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, accept, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// host := c.Writer.Header().Get("Access-Control-Allow-Origin")
		hostValid := os.Getenv("HOST")
		splitHost := strings.Split(hostValid, ",")
		validHost := false
		for _, s := range splitHost {
			if c.Request.Host == s {
				validHost = true
				break
			}
		}
		fmt.Println(c.Request.Host)
		fmt.Println(c.Request.Header.Get("ORIGIN"))

		originValid := os.Getenv("ORIGIN")
		origin := c.Request.Header.Get("Origin")
		validOrigin := false
		if origin != "" {
			splitOrigin := strings.Split(originValid, ",")
			for _, s := range splitOrigin {
				if origin == s {
					validOrigin = true
					break
				}
			}
		}
		if origin == "" {
			validOrigin = true
		}

		if !validHost || !validOrigin {
			c.AbortWithStatus(403)
			// c.Next()
			return
		} else {
			c.Next()
		}

	}
}
