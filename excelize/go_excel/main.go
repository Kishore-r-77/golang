package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func main() {
	xlsx, err := excelize.OpenFile("kishore.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sourceSheetName := "Sheet1"

	filteredXLSX := excelize.NewFile()

	filteredSheetName := "FilteredSheet"

	index, err := filteredXLSX.NewSheet(filteredSheetName)
	if err != nil {
		fmt.Println("Error creating new sheet:", err)
		os.Exit(1)
	}

	filteredXLSX.SetActiveSheet(index)

	rows, err := xlsx.GetRows(sourceSheetName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rowIndex := 1

	for _, row := range rows {

		if rowIndex == 1 {
			for colIndex, cell := range row {
				cellName := fmt.Sprintf("%c%d", 'A'+colIndex, rowIndex)
				filteredXLSX.SetCellValue(filteredSheetName, cellName, cell)
			}
		} else {

			age, err := strconv.Atoi(row[1])

			if err != nil {
				fmt.Printf("Error parsing age in row %d: %v. Skipping row.\n", rowIndex, err)
				continue
			}

			if age > 25 {
				for colIndex, cell := range row {
					cellName := fmt.Sprintf("%c%d", 'A'+colIndex, rowIndex)
					filteredXLSX.SetCellValue(filteredSheetName, cellName, cell)
				}
			}
		}

		rowIndex++
	}

	if err := filteredXLSX.SaveAs("filtered_kishore.xlsx"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Filtered data saved to filtered_kishore.xlsx")
}
