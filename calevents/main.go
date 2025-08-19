package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
)

var (
	monthMap = map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	calCSVHeaderMap = map[string]int{
		"Subject":       0,
		"Start Date":    1,
		"Start Time":    2,
		"End Date":      3,
		"End Time":      4,
		"All Day Event": 5,
	}

	idxSubject     = calCSVHeaderMap["Subject"]
	idxStartDate   = calCSVHeaderMap["Start Date"]
	idxStartTime   = calCSVHeaderMap["Start Time"]
	idxEndDate     = calCSVHeaderMap["End Date"]
	idxEndTime     = calCSVHeaderMap["End Time"]
	idxAllDayEvent = calCSVHeaderMap["All Day Event"]
)

// create an event on the calendarID
// POST https://www.googleapis.com/calendar/v3/calendars/calendarId/events

func calCSVHeaders() []string {
	headers := make([]string, 0, len(calCSVHeaderMap))
	for header := range calCSVHeaderMap {
		headers = append(headers, header)
	}

	return headers
}

func main() {
	f, err := os.Open("na-calendar.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	recs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	headers := calCSVHeaders()

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(headers); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for i, rec := range recs {
		line := make([]string, len(headers))

		if i == 0 {
			continue
		}

		month := rec[0]
		monthInt := monthMap[month]
		day := rec[1]
		event := rec[2]

		var year string
		if slices.Contains([]string{"September", "October", "November", "December"}, month) {
			year = "2025"
		} else {
			year = "2026"
		}

		line[0] = event
		line[1] = fmt.Sprintf("%d/%s/%s", monthInt, day, year)
		line[3] = line[1]
		line[5] = "True"

		if err := w.Write(line); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	slc := buf.Bytes()

	if err := os.WriteFile("formatted-cal.csv", slc, 0644); err != nil {
		log.Fatalln("error writing csv to file:", err)
	}
}
