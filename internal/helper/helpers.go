package helper

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatMACAddress(mac string) ([]string, error) {

	if !isValidMAC(mac) {
		return nil, fmt.Errorf("invalid MAC address format: %s", mac)
	}
	separators := []string{":", "-", "."}
	result := make([]string, 0)
	mac = strings.ReplaceAll(mac, ":", "")

	for _, sep := range separators {
		result = append(result, strings.Join(splitEvery(mac, 2), sep))
	}

	return result, nil
}

func isValidMAC(mac string) bool {
	re := regexp.MustCompile("^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$")
	return re.MatchString(mac)
}

func splitEvery(s string, chunkSize int) []string {
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
