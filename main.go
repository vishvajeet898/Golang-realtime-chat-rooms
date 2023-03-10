package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	go h.run()
	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		fmt.Println("The URL: ", c.Request.Host+c.Request.URL.Path)

		fmt.Printf("roooommmm = %v\n", roomId)
		serveWs(c.Writer, c.Request, roomId)

	})

	port := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	err = router.Run(port)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
}
