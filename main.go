package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eskog/mac/internal/helper"
)

func getMac(mac string) (string, error) {
	const API = "https://api.macvendors.com/"

	resp, err := http.Get(API + mac)
	if err != nil {
		return "", fmt.Errorf("error in http request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response: %w", err)
	}

	return string(body), nil
}

func main() {
	mac := strings.Join(os.Args[1:], "")
	formatted, err := helper.FormatMACAddress(mac)
	if err != nil {
		log.Fatalf("error formatting input: %s", err)
	}

	vendor, err := getMac(mac)
	if err != nil {
		log.Fatalf("error calling API: %s", err)
	}

	fmt.Printf("%s\n\n", vendor)
	for _, result := range formatted {
		fmt.Println(result)
	}

}
