package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

// --------- Expenses and income variables ---------

var catNames []string
var catNamesFood []string
var catNamesRegular []string
var catNamesIrregular []string

type entries struct {
	Category    string
	SubCategory string
	Amount      int
	EntryDate   time.Time
	Comment     string
}

var expenses []entries
var income []entries
var regular []entries

var tpl *template.Template

func main() {

	// --------- Setting the logfile ---------
	var filename string = "logfile.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.TextFormatter)

	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		fmt.Println(err) // Cannot open log file. Logging to stderr
	} else {
		log.SetOutput(f)
	}
	log.Info("Log file set.")

	//

	readCategories()
	income = readFile("income")
	expenses = readFile("expenses")
	regular = readFile("regular")

	tpl = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", home)
	http.HandleFunc("/expenses", registerExpenses)
	http.HandleFunc("/income", registerIncome)
	http.HandleFunc("/regular", regularExpenses)
	http.HandleFunc("/show", showInAndOut)
	http.HandleFunc("/add", addRegular)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
