package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// --------- Read Income, Expenses files and get expense category names---------

func readFile(folderName string) []entries {
	file := lastFile(folderName)

	lines := readLines(file)

	entry := make([]entries, 0)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		sline := strings.Split(line, ",")
		convertAmount, _ := strconv.Atoi(sline[2])
		convertTime, err := time.Parse("02-01-2006  15:04:05", sline[3])
		if err != nil {
			panic(err)
		}
		entry = append(entry, entries{
			Category:    sline[0],
			SubCategory: sline[1],
			Amount:      convertAmount,
			EntryDate:   convertTime,
			Comment:     sline[4],
		})
	}
	log.Info("An information from a file read into slice.", file)

	return entry
}

// The expenses category file
func readCategories() {

	pwd, _ := os.Getwd()
	catNamesF := pwd + "/datafiles/catnames.csv"

	linesCatNames := readLines(catNamesF)

	for i := 0; i < len(linesCatNames); i++ {
		line := linesCatNames[i]
		if i == 0 {
			catNames = strings.Split(line, ",")
		} else if i == 1 {
			catNamesFood = strings.Split(line, ",")
		} else if i == 2 {
			catNamesRegular = strings.Split(line, ",")
		} else if i == 3 {
			catNamesIrregular = strings.Split(line, ",")
		}

	}
}

// --------- Find unique category names - Income ---------

func readNames(income []entries) []string {
	var list []string
	for _, val := range income {
		list = append(list, val.Category)
	}
	names := uniqueStr(list)
	return names
}

// --------- Find unique category and subcategory names - Expenses ---------

func readNamesCat(expenses []entries) []string {
	var list []string
	for _, val := range expenses {
		list = append(list, val.Category)
	}
	names := uniqueStr(list)
	return names
}

func readNamesSubCat(expenses []entries) []string {
	var list []string
	for _, val := range expenses {
		list = append(list, val.SubCategory)
	}
	names := uniqueStr(list)
	return names
}

func readNamesReg(regular []entries) []string {
	var list []string
	for _, val := range regular {
		list = append(list, val.Comment)
	}
	names := uniqueStr(list)
	return names
}

// --------- Save to CSV file ---------

func saveToCSV(data []entries, folder string) {

	log.Println("In saveToCSV received data")

	prDetCol := 5 // somehow automatic

	inStr := make([][]string, len(data)) // initialize an empty 2D slice
	for i := range inStr {
		inStr[i] = make([]string, prDetCol)
		inStr[i][0] = data[i].Category
		inStr[i][1] = data[i].SubCategory
		inStr[i][2] = strconv.Itoa(data[i].Amount)
		inStr[i][3] = data[i].EntryDate.Format("02-01-2006 15:04:05")
		inStr[i][4] = data[i].Comment
	}
	// fmt.Println("In saveToCSV inStr ", inStr)
	// CAN  index out of range error be fixed with the right way to parse the form?

	t := time.Now()
	formattedTime := fmt.Sprintf("%d-%02d-%02d-%02d-%02d-%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	var fileAddress = "datafiles/" + folder + "/" + formattedTime + ".csv"

	file, err := os.Create(fileAddress)
	checkError(err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, value := range inStr {
		err := writer.Write(value)
		checkError(err)
	}

	defer writer.Flush()
}

// used in saveToCSV
func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
