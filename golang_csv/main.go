package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	type Person struct {
		Name  string
		Age   int
		Email string
	}

	people := []Person{
		{Name: "Alice", Age: 25, Email: "alice@example.com"},
		{Name: "Bob", Age: 30, Email: "bob@example.com"},
		// Add more data as needed
	}

	// Open a CSV file for writing.
	file, err := os.Create("people.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row.
	header := []string{"Name", "Age", "Email"}
	writer.Write(header)

	// Write each person's data to the CSV file.
	for _, person := range people {
		row := []string{person.Name, fmt.Sprintf("%d", person.Age), person.Email}
		writer.Write(row)
	}
}
