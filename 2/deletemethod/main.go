package main

import (
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := http.NewRequest("DELETE", "http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
