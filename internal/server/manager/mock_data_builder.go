package manager

import (
	"fmt"

	"github.com/taalhach/aroundhome-challennge/internal/server/database"
	"github.com/xuri/excelize/v2"
)

func InitMockDataBuilder(db *database.DbSession, dataFilePath string) {
	f, err := excelize.OpenFile(dataFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}
