package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func retreive(uri string)  {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
