package common
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string, param map[string]string) string {
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