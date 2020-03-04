package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

func showInAndOut(w http.ResponseWriter, r *http.Request) {
	log.Info("In the showInAndOut")

	t := time.Now()
	yearNow := t.Year()
	monthNow := t.Month()

	// --------- To make needed category lists ---------

	// To make the catogory lists for the expenses overview table
	allCatNames := make([]string, 0)
	allCatNames = categoryNames()

	// To make the catogory lists for the filtering expenses table
	catNames := catNames
	subCatNames := make([]string, 0)
	subCatNames = append(subCatNames, catNamesFood...)
	subCatNames = append(subCatNames, catNamesRegular...)
	subCatNames = append(subCatNames, catNamesIrregular...)

	// --------- To filter Income for this year and months ---------

	tmpIncome := make([]entries, 0)

	for _, val := range income {
		monthEntry := val.EntryDate.Month()
		yearEntry := val.EntryDate.Year()
		if monthNow == monthEntry && yearNow == yearEntry {
			tmpIncome = append(tmpIncome, val)
		}
	}

	// --------- To filter Expenses for this year and months ---------

	tmpExpenses := make([]entries, 0)

	for _, val := range expenses {
		monthEntry := val.EntryDate.Month()
		yearEntry := val.EntryDate.Year()
		if monthNow == monthEntry && yearNow == yearEntry {
			tmpExpenses = append(tmpExpenses, val)
		}
	}

	//
	catSumCalc := catSumNowMonth(tmpExpenses, allCatNames)

	money := map[string]interface{}{
		"YearNow":     yearNow,
		"MonthNow":    monthNow,
		"Income":      tmpIncome,
		"Expenses":    tmpExpenses,
		"CatNames":    catNames,
		"SubCatNames": subCatNames,
		"AllCatNames": allCatNames,
		"CatSumCalc":  catSumCalc,
	}

	render(w, r, tpl, "showInAndOut.html", money)
}
