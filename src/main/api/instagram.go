package api
import (
	rest "lostAndFound/src/main/common"
)

func GetAccessTokenByCode(code string) { 
	url := "https://api.instagram.com/oauth/access_token"
	urlParams := make(map[string]string)
	urlParams["cliet_id"] = "809042689924382"
	urlParams["client_secret"] = "00a814d0e83caadde969132ddf703d9b"
	urlParams["grant_type"] = "authorization_code"
	urlParams["redirect_uri"] = "https://35.194.194.103/ig/accessToken"
	urlParams["code"] = code
	rest.Get(url, urlParams)
}