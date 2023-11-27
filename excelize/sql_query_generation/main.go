package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <excel_file>")
	}

	excelFileName := os.Args[1]

	xlsx, err := excelize.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/policy")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sheetNames := xlsx.GetSheetList()

	for _, sheetName := range sheetNames {
		rows, err := xlsx.GetRows(sheetName)
		if err != nil {
			log.Printf("Error reading sheet %s: %v", sheetName, err)
			continue
		}

		columns := rows[0]

		valuePlaceholders := make([]string, len(columns))
		values := make([]string, len(rows)-1)

		for rowIndex, row := range rows {
			if rowIndex == 0 {
				continue
			}

			valuePlaceholders[rowIndex-1] = "(?" + strings.Repeat(", ?", len(columns)-1) + ")"
			values[rowIndex-1] = "(" + strings.Join(escapeStrings(row), ", ") + ")"
		}

		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", sheetName, strings.Join(columns, ", "), strings.Join(values, ", "))

		outputFile, err := os.Create(fmt.Sprintf("%s.sql", sheetName))
		if err != nil {
			log.Printf("Error creating output file for sheet %s: %v", sheetName, err)
			continue
		}
		defer outputFile.Close()

		_, err = outputFile.WriteString(query)
		if err != nil {
			log.Printf("Error writing SQL query for sheet %s to the file: %v", sheetName, err)
		}
	}
}

func escapeStrings(row []string) []string {
	escaped := make([]string, len(row))
	for i, value := range row {
		escaped[i] = "'" + strings.Replace(value, "'", "''", -1) + "'"
	}
	return escaped
}
