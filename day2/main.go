package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Report struct {
	levels                 []int
	safe                   bool
	problemDampenerApplied bool
}

func (r Report) String() string {
	return fmt.Sprintf("{\nlevels: %v,\nsafe: %v,\nproblemDampenerApplied: %v\n}", r.levels, r.safe, r.problemDampenerApplied)
}

type CheckResult struct {
	safeReports int
}

func main() {
	rawContent := parseFile("./data.txt")
	rows := strings.Split(rawContent, "\n")
	rows = rows[:len(rows)-1]
	reports := make([]Report, len(rows))

	for report_position, row := range rows {
		stringLevels := strings.Fields(row)
		var levels []int

		for i, stringLevel := range stringLevels {
			level, err := strconv.ParseInt(stringLevel, 10, 64)
			if err != nil {
				log.Printf("stringValue: %q, %d", stringLevel, i)
				log.Printf("stringValue: %q", stringLevels)
				log.Fatal(err)
			}
			levels = append(levels, int(level))
		}

		reports[report_position] = Report{levels: levels, safe: true}
	}

	safeReports := 0

	for reportIndex := 0; reportIndex < len(reports); reportIndex++ {
		report := reports[reportIndex]

		newReport := CheckReport(report)

		if newReport.safe {
			safeReports += 1
		}
	}

	fmt.Println("All reports:", len(reports))
	fmt.Printf("Safe Reports: %d", safeReports)
}
