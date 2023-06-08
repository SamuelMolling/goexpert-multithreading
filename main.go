package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c2 := make(chan string)
	c1 := make(chan string)

	go getCEP("http://viacep.com.br/ws/93700000/json/", c1)
	go getCEP("https://cdn.apicep.com/file/apicep/93700-000.json", c2)

	select {
	case r := <-c1:
		fmt.Println("Resposta recebida da API 1\n", r)
	case r := <-c2:
		fmt.Println("Resposta recebida da API 2\n", r)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}

func getCEP(url string, c chan string) {
	URL := url
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	c <- string(body)
}
