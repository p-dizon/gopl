package main

import (
	"time"
)

func isIssueLessThanOneMonthOld(createDate time.Time) bool {
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	return createDate.After(oneMonthAgo)
}

func isIssueOlderThanOneMonth(createDate time.Time) bool {
	return !isIssueLessThanOneMonthOld(createDate)
}

func isIssueLessThanOneYearOld(createDate time.Time) bool {
	oneYearAgo := time.Now().AddDate(-1, 0, 0)
	return createDate.After(oneYearAgo)
}

func isIssueOlderThanOneYear(createDate time.Time) bool {
	return !isIssueLessThanOneYearOld(createDate)
}
