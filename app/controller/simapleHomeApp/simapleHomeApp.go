package simapleHomeApp

import (
	"io/ioutil"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
)

var colArr = [15]string{
	"title",
	"link",
	"button",
	"richText",
	"contentItemActions",
	"comments",
	"ContentItemActionTime",
	"ContentItemMore",
	"titleDesc",
	"unknow",
	"follow",
	"readNum",
	"answerNum",
	"answerAllLink",
	"goodQuestion",
}

func XlsxAnalysis(c *gin.Context) {
	// files, _ := ioutil.ReadDir("./uploads/xlsx/")
	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }
	rootPath, err := os.Getwd()
	dirname := rootPath + "/uploads/xlsx/simapleHomeAppliance/"

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	var filenameSlice []string
	for _, file := range files {
		filenameSlice = append(filenameSlice, file.Name())
	}

	fileDataMap := make(map[string]interface{})
	for _, filename := range filenameSlice {
		excel, err := excelize.OpenFile(dirname + filename)
		if err != nil {
			panic(err)
		}

		// cell, err := excel.GetCellValue("Sheet1", "B2")
		// if err != nil {
		// panic(err)
		// }

		// 获取 Sheet1 上所有单元格
		rows, err := excel.GetRows("sheet1")
		var rowDataMap map[string]interface{}
		var dataMap []map[string]interface{}

		for rowNum, row := range rows {
			rowDataMap = make(map[string]interface{})
			if rowNum == 0 {
				continue
			}

			if row[0] == "" {
				continue
			}

			for colNum, colCell := range row {
				rowDataMap[colArr[colNum]] = colCell
				// fmt.Print(dataMap, "\t")
			}
			// fmt.Println()
			dataMap = append(dataMap, rowDataMap)
		}

		fileDataMap[filename] = dataMap
	}

	c.JSON(200, fileDataMap)
}
