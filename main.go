package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// TODO: Replace it with flags
const ColorschemePath = "./data/colors.json"
const SchemaPath = "./data/scheme.json"

func getMarker(key string) string {
	return fmt.Sprintf("{{%s}}", key)
}

func main() {
	// Parsing colors
	colorsRaw, err := os.ReadFile(ColorschemePath)
	if err != nil {
		panic(err)
	}

	colors := make(map[string]string)
	if err := json.Unmarshal(colorsRaw, &colors); err != nil {
		panic(err)
	}

	// Open schema
	schema, err := os.ReadFile(SchemaPath)
	if err != nil {
		panic(err)
	}

	// Process
	result := string(schema)
	for k, v := range colors {
		result = strings.ReplaceAll(result, getMarker(k), v)
	}

	if err := os.WriteFile("./themes/gruvbox-material.json", []byte(result), 0600); err != nil {
		panic(err)
	}
}
