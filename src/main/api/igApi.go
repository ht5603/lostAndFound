package api
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCreds() map[string]string { 
	creds := make(map[string]string)
	creds["access_token"] = "EAALf0e4kHR4BAG4QPnQbYNuuU9DpQZBELT8ozf4w7aAGAQmP7NTypyqoWGdVCUYd31HhjLpUaSZCKEEnOxD6mtzGSG6FQyDBgJqUCocpK8YY8ORQU0TznyhKBM2Ukq88CypZAw5dL6GCTX8tAJfCV8neIAJUQp0gW66n9QXAMTuYVyMM9g3vqQWQij8FrMGYtcOq5EkdgZDZD"
	creds["client_id"] = "809042689924382" 
	creds["client_secret"] = "00a814d0e83caadde969132ddf703d9b" 
	creds["graph_domain"] = "https://graph.facebook.com/"
	creds["graph_version"] = "v6.0" 
	creds["endpoint_base"] = "https://graph.facebook.com/v6.0/"
	creds["debug"] = "yes"
	fmt.Printf("getCreds : %v \n", creds)
	return creds
}


func DebugAccessToken() {
	creds := GetCreds()
	url := creds["graph_domain"] + "/debug_token"
	urlParams := make(map[string]string)
	urlParams["input_token"] = creds["access_token"]
	urlParams["access_token"] = creds["access_token"]
	HttpGet(url, urlParams)
}

func GetLongLiveAccessToken() {
	creds := GetCreds()
	url := creds["endpoint_base"] + "oauth/access_token"
	urlParams := make(map[string]string)
	urlParams["grant_type"] = "fb_exchange_token"
	urlParams["client_id"] = creds["client_id"]
	urlParams["client_secret"] = creds["client_secret"]
	urlParams["fb_exchange_token"] = creds["access_token"]
	HttpGet(url, urlParams)
}

func GetUserPages() {
	creds := GetCreds()
	url := creds["endpoint_base"] + "me/accounts"
	urlParams := make(map[string]string)
	urlParams["access_token"] = creds["access_token"]
	HttpGet(url, urlParams)
}

func HttpGet(url string, param map[string]string) string {
	url = url + "?"
	for key, value := range param {
        url = url + key + "=" + value + "&"
	}
	fmt.Printf("http get url:%v \n", url)
	resp, err := http.Get(url)
    if err != nil {
        print(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        print(err)
    }
	fmt.Printf("response:%v \n",string(body))
	return string(body)
}