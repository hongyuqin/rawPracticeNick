package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func main() {
	xlsx, err := xlsx.OpenFile("./Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, sheet := range xlsx.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
