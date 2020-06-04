package main

import (
	"testing"
	"time"
)

func TestIssueLessThanAMonthOld(t *testing.T) {
	input := time.Now()
	want := true
	output := isIssueLessThanOneMonthOld(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}

func TestIssueLessThanAMonthOld2(t *testing.T) {
	input := time.Now().AddDate(-1, 0, 0)
	want := false
	output := isIssueLessThanOneMonthOld(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}

func TestIssueLessThanAMonthOld3(t *testing.T) {
	input := time.Now().AddDate(-1, 0, 0)
	want := true
	output := isIssueOlderThanOneMonth(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}

func TestIssueLessThanAMonthOld4(t *testing.T) {
	input := time.Now()
	want := false
	output := isIssueOlderThanOneMonth(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}

func TestIssueLessThanAYearOld(t *testing.T) {
	input := time.Now()
	want := true
	output := isIssueLessThanOneYearOld(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}

func TestIssueLessThanAYearOld2(t *testing.T) {
	input := time.Now().AddDate(-2, 0, 0)
	want := false
	output := isIssueLessThanOneYearOld(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}

func TestIssueLessThanAYearOld3(t *testing.T) {
	input := time.Now().AddDate(-2, 0, 0)
	want := true
	output := isIssueOlderThanOneYear(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}

func TestIssueLessThanAYearOld4(t *testing.T) {
	input := time.Now()
	want := false
	output := isIssueOlderThanOneYear(input)
	if want != output {
		t.Errorf("Wanted %v, got %v", want, output)
	}
}
