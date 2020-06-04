package main

import (
	"testing"
	"time"

	"./issues"
)

func TestGetIssuesByDate(t *testing.T) {
	searchResult := issues.IssuesSearchResult{
		TotalCount: 5,
		Items: []*issues.Issue{
			{Number: 1, CreatedAt: time.Now()},
			{Number: 2, CreatedAt: time.Now().AddDate(0, -2, 0)},
			{Number: 3, CreatedAt: time.Now().AddDate(0, -2, 0)},
			{Number: 4, CreatedAt: time.Now().AddDate(-3, 0, 0)},
			{Number: 5, CreatedAt: time.Now().AddDate(-3, 0, 0)},
			{Number: 6, CreatedAt: time.Now().AddDate(-3, 0, 0)},
		},
	}

	issuesLessThanAMonthOld, issuesLessThanAYearOld, issuesMoreThanAYearOld := getIssuesByDate(&searchResult)
	if issuesLessThanAMonthOld.TotalCount != 1 {
		t.Errorf("Wanted 1 issue, got %v", issuesLessThanAMonthOld.TotalCount)
	}
	if issuesLessThanAMonthOld.Items[0].Number != 1 {
		t.Errorf("Wanted number 1, got %v", issuesLessThanAMonthOld.Items[0].Number)
	}
	if issuesLessThanAYearOld.TotalCount != 2 {
		t.Errorf("Wanted 2 issue, got %v", issuesLessThanAYearOld.TotalCount)
	}
	if issuesLessThanAYearOld.Items[0].Number != 2 {
		t.Errorf("Wanted number 2, got %v", issuesLessThanAYearOld.Items[0].Number)
	}
	if issuesMoreThanAYearOld.TotalCount != 3 {
		t.Errorf("Wanted 2 issue, got %v", issuesMoreThanAYearOld.TotalCount)
	}
	if issuesMoreThanAYearOld.Items[0].Number != 4 {
		t.Errorf("Wanted number 4, got %v", issuesMoreThanAYearOld.Items[0].Number)
	}
}
