package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()
	router := gin.New()
	//router.LoadHTMLFiles("index.html")

	/*router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})*/

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		fmt.Printf("roooommmm = %v\n", roomId)
		serveWs(c.Writer, c.Request, roomId)
	})

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}
