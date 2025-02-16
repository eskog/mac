package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetMac(mac string) string {
	const API = "https://api.macvendors.com/"

	resp, err := http.Get(API + mac)
	if err != nil {
		log.Fatalln("Error requesting http: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading the response: %w", err)
	}

	return string(body)
}

func FormatMACAddress(mac string) (colon, dashed, dotted string, err error) {

	mac = strings.ReplaceAll(mac, ":", "")
	mac = strings.ReplaceAll(mac, "-", "")

	if len(mac) != 12 {
		return "", "", "", fmt.Errorf("invalid MAC address format: %s", mac)
	}

	colon = strings.ToLower(strings.Join(SplitEvery(mac, 2), ":"))
	dashed = strings.ToLower(strings.Join(SplitEvery(mac, 2), "-"))
	dotted = strings.ToLower(strings.Join(SplitEvery(mac, 4), "."))

	return colon, dashed, dotted, nil
}

func SplitEvery(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func main() {
	mac := strings.Join(os.Args[1:], "")
	vendor := GetMac(mac)
	colon, dashed, dotted, err := FormatMACAddress(mac)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(vendor, "\n\n", colon, "\n", dashed, "\n", dotted)
}
