//go:build ignore
// +build ignore

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	f, err := os.Open("../assets/world-cities.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	outFile, err := os.Create("world-cities.go")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Write the package and Slice declaration
	outFile.WriteString("package assets\n\n")
	outFile.WriteString("var WorldCities = [][]string{\n")

	for i, record := range records {
		if i == 0 {
			// Skip the header row
			continue
		}
		outFile.WriteString(fmt.Sprintf("\t{\"%s\", \"%s\", \"%s\"},\n", record[0], record[2], record[3]))
	}
	outFile.WriteString("}\n")
}
