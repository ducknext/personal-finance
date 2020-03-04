package main

import (
	"time"
)

// *********** Calculations for the HOME page ***********

// --------- Current month income and expenses ---------

func sumNowMonth(data []entries) int {
	t := time.Now()
	yearNow := t.Year()
	monthNow := t.Month()
	var monthSum = 0
	for _, val := range data {
		monthEntry := val.EntryDate.Month()
		yearEntry := val.EntryDate.Year()

		if monthNow == monthEntry && yearNow == yearEntry {
			monthSum = monthSum + val.Amount
		}
	}
	return monthSum
}

// --------- Expenses and income for all recorded period ---------

func sumAll(data []entries) int {
	var monthSum = 0
	for _, val := range data {
		monthSum = monthSum + val.Amount
	}
	return monthSum
}

// --------- Expenses and income per year for all recorded period ---------

func sumByYear(data []entries, years []int) []int {
	sum := make([]int, 0)
	for _, year := range years {
		var yearSum = 0
		for _, val := range data {
			yearEntry := val.EntryDate.Year()
			if year == yearEntry {
				yearSum = yearSum + val.Amount
			}
		}
		sum = append(sum, yearSum)
	}
	return sum
}

// --------- Expenses and income per month for all recorded period ---------

func sumByMonth(data []entries, years []int, months []string) []int {
	sum := make([]int, 0)
	for _, year := range years {
		for _, month := range months {
			var monthSum = 0
			for _, val := range data {
				monthEntry := val.EntryDate.Month().String()
				yearEntry := val.EntryDate.Year()
				if year == yearEntry && month == monthEntry {
					monthSum = monthSum + val.Amount
				}
			}
			sum = append(sum, monthSum)
		}
	}
	return sum
}

// *********** Calculation for the SHOW page ***********

func catSumNowMonth(expenses []entries, categories []string) []int {
	catSum := make([]int, 0)
	for _, cat := range categories {
		var sum = 0
		for _, val := range expenses {
			if cat == val.Category || cat == val.SubCategory {
				sum = sum + val.Amount
			}
		}
		catSum = append(catSum, sum)
	}
	return catSum
}
