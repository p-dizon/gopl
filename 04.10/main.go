package main

import (
	"fmt"
	"log"
	"os"

	"./issues"
)

func main() {
	searchQuery := os.Args[1:]
	result, err := issues.SearchIssues(searchQuery)
	if err != nil {
		log.Fatal(err)
	}
	issuesLessThanAMonthOld, issuesLessThanAYearOld, issuesMoreThanAYearOld := getIssuesByDate(result)
	displayIssuesLessThanAMonthOld(issuesLessThanAMonthOld)
	displayIssuesLessThanAYearOld(issuesLessThanAYearOld)
	displayIssuesMoreThanAYearOld(issuesMoreThanAYearOld)
}

func getIssuesByDate(result *issues.IssuesSearchResult) (*issues.IssuesSearchResult, *issues.IssuesSearchResult, *issues.IssuesSearchResult) {
	var issuesLessThanAMonthOld, issuesLessThanAYearOld, issuesMoreThanAYearOld issues.IssuesSearchResult
	for _, item := range result.Items {
		switch {
		case isIssueLessThanOneMonthOld(item.CreatedAt):
			issuesLessThanAMonthOld.Items = append(issuesLessThanAMonthOld.Items, item)
			issuesLessThanAMonthOld.TotalCount++
		case isIssueOlderThanOneMonth(item.CreatedAt) && isIssueLessThanOneYearOld(item.CreatedAt):
			issuesLessThanAYearOld.Items = append(issuesLessThanAYearOld.Items, item)
			issuesLessThanAYearOld.TotalCount++
		default:
			issuesMoreThanAYearOld.Items = append(issuesMoreThanAYearOld.Items, item)
			issuesMoreThanAYearOld.TotalCount++
		}
	}
	return &issuesLessThanAMonthOld, &issuesLessThanAYearOld, &issuesMoreThanAYearOld
}

func displayIssuesLessThanAMonthOld(results *issues.IssuesSearchResult) {
	printHeader("Issues less than a month old", results.TotalCount)
	printIssues(results.Items)
}

func displayIssuesLessThanAYearOld(results *issues.IssuesSearchResult) {
	printHeader("Issues less than a year old", results.TotalCount)
	printIssues(results.Items)
}

func displayIssuesMoreThanAYearOld(results *issues.IssuesSearchResult) {
	printHeader("Issues more than a year old", results.TotalCount)
	printIssues(results.Items)
}

func printHeader(message string, totalCount int) {
	fmt.Printf("******************** %d %s ********************\n", totalCount, message)
}

func printIssues(i []*issues.Issue) {
	if len(i) == 0 {
		fmt.Printf("No results found\n\n")
		return
	}

	for _, item := range i {
		fmt.Printf("%-12s #%-5d %10.9s %.55s \n",
			item.CreatedAt.Format("2006-01-02"), item.Number, item.User.Login, item.Title)
	}
	fmt.Println()
}
