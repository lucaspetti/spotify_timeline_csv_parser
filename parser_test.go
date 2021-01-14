package main

import (
	"testing"
	"reflect"
)

func TestParse(t *testing.T) {
	csv_file := "timelines_test.csv"

	got := Parse(csv_file)
	want := []MonthlySum{
		MonthlySum{3, 2020, 300, 600, 50},
		MonthlySum{4, 2020, 100, 800, 20},
		MonthlySum{4, 2019, 50, 400, 20},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
