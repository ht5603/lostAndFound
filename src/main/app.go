package main
import (
	"fmt"
	// "lostAndFound/src/main/api"
	"github.com/gin-gonic/gin"
)
func main() { 
	fmt.Printf("main start! \n")

	r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
		value := c.Query("id")
		fmt.Printf("id:%v \n", value)
        c.JSON(200, gin.H{
            "message": "pong",
        })
	})

	r.GET("/ig/auth", func(c *gin.Context) {
		value := c.Query("code")
		fmt.Printf("ig auth code:%v \n", value)
        c.JSON(200, gin.H{
            "authCode": value,
        })
	})

	r.GET("/ig/cancelAuth", func(c *gin.Context) {
		value := c.Query("code")
		fmt.Printf("ig cancel auth code:%v \n", value)
        c.JSON(200, gin.H{
            "IG cancel": value,
        })
	})
	


    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}