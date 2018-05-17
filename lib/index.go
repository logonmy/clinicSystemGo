package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

// ReadXlsx 读取Excel文件
func ReadXlsx(filePath string) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}

	for _, row := range xlFile.Sheets[0].Rows {
		for _, cell := range row.Cells {
			text := cell.String()
			fmt.Printf("%s\t", text)
		}
		fmt.Printf("\n")
	}
}



func main() {
	excelFileName := "lab.xlsx"
	ReadXlsx(excelFileName)
	// read(excelFileName)
}
