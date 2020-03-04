package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

func addRegular(w http.ResponseWriter, r *http.Request) {
	log.Info("In the regular expenses")

	t := time.Now()
	layout := "2006-01-02 15:04:05"
	str := t.Format("2006-01-02 15:04:05")
	entryTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	}

	for _, val := range regular {
		expenses = append(expenses, entries{
			Category:    "Regular",
			SubCategory: "Bills",
			Amount:      val.Amount,
			EntryDate:   entryTime,
			Comment:     val.Comment,
		})
	}

	saveToCSV(expenses, "expenses")

	money := map[string]interface{}{
		"RegularExp": regular,
	}

	render(w, r, tpl, "addRegular.html", money)
}
