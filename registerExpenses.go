package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

func registerExpenses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		log.Info("If in the GET in /expenses.")
		render(w, r, tpl, "registerExpenses.html", nil)
		return
	}

	// defensive programming

	log.Info("If in the POST in /expenses.")
	r.ParseForm()

	log.Println(r.Form) // Log all data. Form is a map[]

	category := r.Form.Get("category")
	subCategory := r.Form.Get("categorySelect")
	amount := r.Form.Get("amount")
	comment := r.Form.Get("comment")
	date := r.Form.Get("date")

	a, _ := strconv.Atoi(amount) // TODO: what if it is a decimal?

	layout := "2006-01-02 15:04:05"
	dateStr := date + " 20:20:20"
	readTime, err := time.Parse(layout, dateStr)

	str := time.Now().Format("2006-01-02 15:04:05")
	entryTime, err := time.Parse(layout, str)
	if err != nil { // TODO: what this does?
		fmt.Println(err)
		render(w, r, tpl, "registerExpenses.html", nil)
		return
	}

	var entryDate time.Time
	if date != "" {
		entryDate = readTime
	} else {
		entryDate = entryTime
	}

	expenses = append(expenses, entries{
		Category:    category,
		SubCategory: subCategory,
		Amount:      a,
		EntryDate:   entryDate,
		Comment:     comment,
	})

	saveToCSV(expenses, "expenses")

	render(w, r, tpl, "registerExpenses.html", nil)

}
