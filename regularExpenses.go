package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

var ind = 0 // THIS might cause big problems!!! 11.11.2019

func regularExpenses(w http.ResponseWriter, r *http.Request) {

	regularNames := readNamesReg(regular)

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		log.Info("If in the GET in /regular.")

		money := map[string]interface{}{
			"RegularExp":   regular,
			"RegularNames": regularNames,
		}

		render(w, r, tpl, "regularExp.html", money)

		return
	}

	log.Info("If in the POST in /regular.")
	r.ParseForm()

	log.Println(r.Form)

	regExp := r.Form.Get("regularExpense")
	amountCh, _ := strconv.Atoi(r.Form.Get("amountChange"))

	// use pointers to change the value
	ind = indexOf(regExp, regular)

	for _, name := range regularNames {
		if name == regExp {
			var pamountCh *int = &regular[ind].Amount
			ptf(pamountCh, amountCh)
		}
	}

	newRegExp := r.Form.Get("name")
	amountNew, _ := strconv.Atoi(r.Form.Get("amountNew"))

	t := time.Now()
	layout := "2006-01-02 15:04:05"
	str := t.Format("2006-01-02 15:04:05")
	entryTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}

	// Save only if there is something
	if amountNew != 0 {
		regular = append(regular, entries{
			Category:    "Regular",
			SubCategory: "Bills",
			Amount:      amountNew,
			EntryDate:   entryTime,
			Comment:     newRegExp,
		})
	}

	saveToCSV(regular, "regular")

	money := map[string]interface{}{
		"RegularExp":   regular,
		"RegularNames": regularNames,
	}

	render(w, r, tpl, "regularExp.html", money)
}
