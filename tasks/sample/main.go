package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var err error

type Foo struct {
	ID int `json:"id"`
}

func main() {
	url := "https://dotstamp.com/api/contributions/15"
	r, _ := http.Get(url)
	log.Println(r.Body)
	defer r.Body.Close()

	// Bodyを読み込む
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	output := Foo{}
	err = json.Unmarshal(body, &output)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Println(output)

}
