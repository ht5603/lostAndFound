package api
import (
	"fmt"
	rest "lostAndFound/src/main/common"
)

func GetAccessTokenByCode(code string) { 
	url := "https://api.instagram.com/oauth/access_token"
	urlParams := make(map[string]string)
	urlParams["cliet_id"] = "776181923221122"
	urlParams["client_secret"] = "ff94d60aaf19a49c2b90177fa8a4b639"
	urlParams["grant_type"] = "authorization_code"
	urlParams["redirect_uri"] = "https://35.194.194.103/ig/auth/"
	urlParams["code"] = code
	fmt.Printf("執行取得accessToken, code:%v \n url:%v \n", code, url)
	rest.Post(url, urlParams)
}