package gif

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
)

var colArr = [5]string{
	"title",
	"link",
	"src",
	"label",
	"rng",
}

func XlsxAnalysis(c *gin.Context) {
	rootPath, err := os.Getwd()
	dirname := rootPath + "/uploads/xlsx/gif/"

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

	c.JSON(http.StatusOK, fileDataMap)
}

func DownloadGif(c *gin.Context) {
	rootPath, err := os.Getwd()
	dirname := rootPath + "/uploads/xlsx/gif/"

	downloadDir := rootPath + "/uploads/gif/"

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	var filenameSlice []string
	for _, file := range files {
		filenameSlice = append(filenameSlice, file.Name())
	}

	// fileDataMap := make(map[string]interface{})
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
		// var rowDataMap map[string]interface{}
		// var dataMap []map[string]interface{}

		for rowNum, row := range rows {
			// rowDataMap = make(map[string]interface{})
			if rowNum == 0 {
				continue
			}

			if row[0] == "" {
				continue
			}

			if row[2] == "https://www.soogif.com/images/img/img-home-page/default.png" {
				continue
			}

			dir := downloadDir + strings.Replace(row[3], "#", "", 1)
			gifFilename := dir + "/" + row[4] + "_" + strings.Replace(row[0], "\"", "", -1) + ".gif"

			response, err := http.Get(row[2])

			if err != nil {
				panic(err)
			}

			body, err := ioutil.ReadAll(response.Body)

			if err != nil {
				panic(err)
			}

			if !isExist(dir) {
				os.MkdirAll(dir, os.ModePerm)
			}

			out, err := os.Create(gifFilename)

			if err != nil {
				panic(err)
			}

			_, err = io.Copy(out, bytes.NewReader(body))

			if err != nil {
				panic(err)
			}
			// row[3] //label
			// row[2] //src
			// row[0] //title
			// row[4] //rng

			// for colNum, colCell := range row {
			// 	rowDataMap[colArr[colNum]] = colCell
			// 	// fmt.Print(dataMap, "\t")
			// }
			// // fmt.Println()
			// dataMap = append(dataMap, rowDataMap)
		}

		// fileDataMap[filename] = dataMap
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
