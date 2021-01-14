package main

import (
	"encoding/csv"
	"io"
	"os"
	// "fmt"
	"strconv"
	"log"
	"strings"
)

type MonthlySum struct {
	Month int
	Year int
	Listeners int
	Streams int
	Followers int
}

func (m *MonthlySum) AddListeners(amount int) {
	m.Listeners += amount
}

func (m *MonthlySum) AddStreams(amount int) {
	m.Streams += amount
}

func (m *MonthlySum) IsSameMonthAndYear(month, year int) (result bool) {
	result = m.Month == month && m.Year == year
	return
}

func Parse(filename string) (records []MonthlySum) {
	csv_file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open the csv file", err)
	}

	r := csv.NewReader(csv_file)

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if row[0] == "date" { continue }

		month, year := getMonthAndYear(row[0])
		listeners := convertString(row[1])
		streams := convertString(row[2])
		followers := convertString(row[3])
		var added bool

		for i := range records {
			if records[i].IsSameMonthAndYear(month, year) {
				sum := &records[i]
				sum.AddListeners(listeners)
				sum.AddStreams(streams)
				added = true
			}
		}

		if added { continue }

		sum := MonthlySum{month, year, listeners, streams, followers}
		records = append(records, sum)
	}

	return
}

func convertString(str string) (value int) {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	return
}

func getMonthAndYear(str string) (month, year int) {
	splitDate := strings.Split(str, "-")
	year = convertString(splitDate[0])
	month = convertString(splitDate[1])
	return
}