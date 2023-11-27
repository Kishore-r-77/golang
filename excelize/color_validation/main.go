package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Open the Excel file
	file := "kishore.xlsx"
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the sheet name
	sheetName := "Sheet1"

	// Get the rows from the sheet
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create a new style for red fill
	redFillStyle, _ := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#f283a6"}, Pattern: 1},
	})
	greenFillStyle, _ := xlsx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#83f2a1"}, Pattern: 1},
	})

	// Get the header row to find the index of the "Age" and "isIntern" columns
	headerRow := rows[0]

	// Find the index of the "Age" and "isIntern" columns
	ageColumnIndex := -1
	isInternColumnIndex := -1
	for i, cellValue := range headerRow {
		switch cellValue {
		case "Age":
			ageColumnIndex = i
		case "isIntern":
			isInternColumnIndex = i
		}
	}

	// Check if the "Age" and "isIntern" columns were found
	if ageColumnIndex == -1 || isInternColumnIndex == -1 {
		fmt.Println("Column 'Age' or 'isIntern' not found in the header row.")
		os.Exit(1)
	}

	// Iterate through the rows and update "isIntern" based on "Age"
	for rowIndex, row := range rows {
		// Skip the header row
		if rowIndex == 0 {
			continue
		}

		// Parse the age value
		age := row[ageColumnIndex]

		// Convert the age to an integer
		ageInt := 0
		fmt.Sscanf(age, "%d", &ageInt)

		// Update the "isIntern" value based on age
		var isInternValue bool
		if ageInt < 25 {
			isInternValue = true
		} else {
			isInternValue = false
		}

		// Update the cell in the "isIntern" column
		isInternCell := fmt.Sprintf("%c%d", 'A'+isInternColumnIndex, rowIndex+1)
		xlsx.SetCellValue(sheetName, isInternCell, isInternValue)

		// Highlight the "isIntern" cell based on its value
		if isInternValue {
			xlsx.SetCellStyle(sheetName, isInternCell, isInternCell, greenFillStyle)
		} else {
			xlsx.SetCellStyle(sheetName, isInternCell, isInternCell, redFillStyle)
		}

		// Highlight the "Age" cell based on the original condition
		if ageInt < 25 {
			cell := fmt.Sprintf("%c%d", 'A'+ageColumnIndex, rowIndex+1)
			xlsx.SetCellStyle(sheetName, cell, cell, redFillStyle)
		} else if ageInt > 25 {
			cell := fmt.Sprintf("%c%d", 'A'+ageColumnIndex, rowIndex+1)
			xlsx.SetCellStyle(sheetName, cell, cell, greenFillStyle)
		}
	}

	// Save the modified file
	err = xlsx.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Excel file updated and saved.")
}
