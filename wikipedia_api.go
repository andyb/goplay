package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Started...")
	resp, err := http.Get("http://en.wikipedia.org/w/api.php?action=opensearch&search=Te")
	if err != nil {
		log.Fatal("Failed to get URL")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("%s", body)
}
