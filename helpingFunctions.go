package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
)

//----------------------------------------------------------------------------------------------------

func readCurrentDir(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	return list
}

func lastFile(folderName string) string {
	pwd, _ := os.Getwd()
	dir := pwd + "/datafiles/" + folderName
	fileList := readCurrentDir(dir)
	fileListLast := fileList[len(fileList)-1]
	file := dir + "/" + fileListLast
	return file
}

//----------------------------------------------------------------------------------------------------

func uniqueStr(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func indexOf(name string, data []entries) int { // very specific type
	for k, v := range data {
		if v.Comment == name {
			return k
		}
	}
	return 0 //not found. //may cause problems 11.11.2019
}

//----------------------------------------------------------------------------------------------------

// use pointers to change the value
func ptf(a *int, value int) {
	*a = value
}

//----------------------------------------------------------------------------------------------------

func readLines(f string) []string {
	// Get the file and split it in lines
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	var txtlines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	return txtlines
}

//----------------------------------------------------------------------------------------------------

// Render a template, or server error.
// From https://github.com/meshhq/golang-html-template-tutorial
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template,
	name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

//----------------------------------------------------------------------------------------------------

func categoryNames() []string {
	names := make([]string, 0)
	names = append(names, catNames[0])
	names = append(names, catNamesFood...)
	names = append(names, catNames[1])
	names = append(names, catNamesRegular...)
	names = append(names, catNames[2])
	names = append(names, catNamesIrregular...)

	return names
}

//----------------------------------------------------------------------------------------------------
// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
// https://programming.guide/go/find-search-contains-slice.html
func findIndex(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}
