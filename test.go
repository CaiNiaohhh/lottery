package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Do(id int) {
	url := "http://127.0.0.1:8080/get?uid=" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("id is ", id, " --- ", string(body))
}

func main() {
	resp, err := http.Get("http://127.0.0.1:8080/set?num=10&total=100")
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Duration(2)*time.Second)
	for i := 0; i < 100; i++ {
		go Do(i)
	}
	time.Sleep(time.Duration(5)*time.Second)
	defer resp.Body.Close()

}
