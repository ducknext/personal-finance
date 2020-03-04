package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

func registerIncome(w http.ResponseWriter, r *http.Request) {

	incomeNames := readNames(income)

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		log.Info("If in the GET in /register.")

		money := map[string]interface{}{
			"IncomeNames": incomeNames,
		}

		render(w, r, tpl, "registerIncome.html", money)

		return
	}

	log.Info("If in the POST in /register.")
	r.ParseForm()

	log.Println(r.Form)

	incomeSource := r.Form.Get("incomeSource")
	newSource := r.Form.Get("name")
	amount := r.Form.Get("amount")
	comment := r.Form.Get("comment")

	a, _ := strconv.Atoi(amount)

	t := time.Now()
	layout := "2006-01-02 15:04:05"
	str := t.Format("2006-01-02 15:04:05")
	incomeTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}

	if newSource != "" {
		income = append(income, entries{
			Category:  newSource,
			Amount:    a,
			EntryDate: incomeTime,
			Comment:   comment,
		})
	} else {
		income = append(income, entries{
			Category:  incomeSource,
			Amount:    a,
			EntryDate: incomeTime,
			Comment:   comment,
		})
	}
	saveToCSV(income, "income")

	// incomeNames := readNames(income)

	money := map[string]interface{}{
		"IncomeNames": incomeNames,
	}

	render(w, r, tpl, "registerIncome.html", money)
}
