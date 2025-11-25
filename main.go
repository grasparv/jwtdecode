package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func decodeSegment(seg string) (map[string]any, error) {
	b, err := base64.RawURLEncoding.DecodeString(seg)
	if err != nil {
		return nil, err
	}

	var out map[string]any
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: jwtdecode <token>")
		return
	}

	parts := strings.Split(os.Args[1], ".")
	if len(parts) < 2 {
		fmt.Println("Invalid JWT")
		return
	}

	header, err := decodeSegment(parts[0])
	if err != nil {
		fmt.Println("Header decode error:", err)
		return
	}

	payload, err := decodeSegment(parts[1])
	if err != nil {
		fmt.Println("Payload decode error:", err)
		return
	}

	fmt.Println("Header:")
	h, _ := json.MarshalIndent(header, "", "  ")
	fmt.Println(string(h))

	fmt.Println("\nPayload:")
	p, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println(string(p))
}
