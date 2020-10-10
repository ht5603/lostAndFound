package main
import (
	"fmt"
	"io/ioutil"
	"lostAndFound/src/main/api"
	"github.com/gin-gonic/gin"
)
func main() { 
	fmt.Printf("main start! \n")

	r := gin.Default()
	r.LoadHTMLGlob("www/*")

	r.GET("/ig/startPage", func(c *gin.Context) {
		c.HTML(200, "testIG.html", nil)
	})

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
		api.GetAccessTokenByCode(value)
        c.JSON(200, gin.H{
            "authCode": value,
        })
	})

	r.POST("/ig/accessToken", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
		  // Handle error
		  fmt.Printf("accessToken error!")
		}
		jsonValue := string(jsonData)
		fmt.Printf("accessToken json:", jsonValue)
        c.JSON(200, gin.H{
            "authCode": jsonValue,
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