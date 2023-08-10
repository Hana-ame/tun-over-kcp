package main

import (
	"github.com/Hana-ame/tun-over-kcp/utils"
	"github.com/gin-gonic/gin"
)

type ServerInfo struct {
	IP   string
	Info string
}

var (
	infoMap = utils.NewLockedMap()
	listMap = utils.NewLockedMap()
)

// usage:
// POST /[key]?addr=[server_ip]
// GET /[key]?addr=[server_ip]
// GET /[key]/list

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route handler
	router.POST("/:key", func(c *gin.Context) {
		var o struct {
			Key string `uri:"key"`
		}
		if err := c.ShouldBindUri(&o); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		key := o.Key
		addr := c.Query("addr")
		infoMap.Put(key, addr)
		listMap.Put(key, []string{})
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	router.GET("/:key", func(c *gin.Context) {
		var o struct {
			Key string `uri:"key"`
		}
		if err := c.ShouldBindUri(&o); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		key := o.Key
		addr := c.Query("addr")
		if list, ok := listMap.Get(key); ok {
			if l, ok := list.([]string); ok {
				if len(l) > 128 { // 简易限制
					l = []string{}
				}
				l = append(l, addr)
				listMap.Put(key, l)
			}
		}
		info, _ := infoMap.Get(key)
		c.JSON(200, info)
	})

	router.GET("/:key/list", func(c *gin.Context) {
		var o struct {
			Key string `uri:"key"`
		}
		if err := c.ShouldBindUri(&o); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		key := o.Key
		if list, ok := listMap.Get(key); ok {
			if l, ok := list.([]string); ok {
				listMap.Put(key, []string{})
				c.JSON(200, l)
				return
			}
		}
		c.JSON(404, gin.H{"error": "not found"})
	})

	// Run the server
	router.Run("127.99.0.1:8080")
}
