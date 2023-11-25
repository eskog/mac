package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	mac := strings.Join(os.Args[1:], "")
	macurl := "https://api.macvendors.com/"

	resp, err := http.Get(macurl + mac)
	if err != nil {
		log.Fatalf("Error accessing %v %v", macurl, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading the response", err)
	}

	fmt.Println(string(body))
}
