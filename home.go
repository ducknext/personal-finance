package main

import (
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Info("In the HOME")

	t := time.Now()
	yearNow := t.Year()
	monthNow := t.Month()

	firstYear := expenses[0].EntryDate.Year()
	firstMonth := expenses[0].EntryDate.Month()

	// Find all unique years
	yearsList := make([]int, 0)
	for _, val := range expenses {
		yearsList = append(yearsList, val.EntryDate.Year())
	}
	uniqueYearsList := unique(yearsList)

	// Find all unique months
	monthList := make([]string, 0)
	for _, val := range expenses {
		yearMonth := val.EntryDate.Month().String()
		monthList = append(monthList, yearMonth)
	}
	uniqueMonthList := uniqueStr(monthList)

	// Find all unique year and month combinations
	yearMonthList := make([]string, 0)
	for _, val := range expenses {
		yearMonth := val.EntryDate.Month().String() + " " + strconv.Itoa(val.EntryDate.Year())
		yearMonthList = append(yearMonthList, yearMonth)
	}
	uniqueYMList := uniqueStr(yearMonthList)

	// To make the catogory list
	// allCatNames := make([]string, 0)
	// allCatNames = categoryNames()

	// --------- SUMS ---------

	// --------- For the current month ---------
	incomeSumNowMonth := sumNowMonth(income)
	expensesSumNowMonth := sumNowMonth(expenses)
	leftNowMonth := incomeSumNowMonth - expensesSumNowMonth

	// --------- For all period ---------
	exSumAll := sumAll(expenses)
	inSumAll := sumAll(income)
	leftSumAll := inSumAll - exSumAll

	sumByYearEx := sumByYear(expenses, uniqueYearsList)
	sumByYearIn := sumByYear(income, uniqueYearsList)

	savingsByYear := make([]int, 0)
	for i := 0; i < len(sumByYearIn); i++ {
		dif := sumByYearIn[i] - sumByYearEx[i]
		savingsByYear = append(savingsByYear, dif)
	}

	sumByMonthEx := sumByMonth(expenses, uniqueYearsList, uniqueMonthList)
	sumByMonthIn := sumByMonth(income, uniqueYearsList, uniqueMonthList)

	savingsByMonth := make([]int, 0)
	for i := 0; i < len(sumByMonthIn); i++ {
		dif := sumByMonthIn[i] - sumByMonthEx[i]
		savingsByMonth = append(savingsByMonth, dif)
	}

	// fmt.Println(catSumAllMonths)

	money := map[string]interface{}{
		"YearNow":             yearNow,
		"MonthNow":            monthNow,
		"FirstYear":           firstYear,
		"FirstMonth":          firstMonth,
		"IncomeSumNowMonth":   incomeSumNowMonth,
		"ExpensesSumNowMonth": expensesSumNowMonth,
		"LeftNowMonth":        leftNowMonth,
		"ExSumAll":            exSumAll,
		"InSumAll":            inSumAll,
		"LeftSumAll":          leftSumAll,
		"UniqueYearsList":     uniqueYearsList,
		"UniqueYMList":        uniqueYMList,
		"SumByMonthEx":        sumByMonthEx,
		"SumByMonthIn":        sumByMonthIn,
		"SavingsByMonth":      savingsByMonth,
		"SumByYearEx":         sumByYearEx,
		"SumByYearIn":         sumByYearIn,
		"SavingsByYear":       savingsByYear,
	}

	render(w, r, tpl, "home.html", money)

}
