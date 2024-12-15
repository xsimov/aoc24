package main

import (
	"fmt"
	"testing"
)

func TestProblemDampener(t *testing.T) {
	t.Run("Report with a distance of 0", func(t *testing.T) {
		reportWithOneProblem := Report{levels: []int{45, 48, 51, 52, 55, 58, 60, 60}, safe: false}
		vector := calculateVector(reportWithOneProblem)
		fmt.Print(vector)

		report := ProblemDampener(reportWithOneProblem, vector)

		if !report.safe {
			t.Errorf("Expected report to be safe, but it was not")
		}
	})

	t.Run("reports with only one level >3 in the end are safe", func(t *testing.T) {
		reportWithOneProblem := Report{levels: []int{1, 3, 4, 8}, safe: false}
		vector := calculateVector(reportWithOneProblem)

		report := ProblemDampener(reportWithOneProblem, vector)

		if !report.safe {
			t.Errorf("Expected report to be safe, but it was not")
		}
	})

	t.Run("reports with only one level >3 in the beginning are safe", func(t *testing.T) {
		reportWithOneProblem := Report{levels: []int{1, 6, 7, 8}, safe: false}
		vector := calculateVector(reportWithOneProblem)

		report := ProblemDampener(reportWithOneProblem, vector)

		if !report.safe {
			t.Errorf("Expected report to be safe, but it was not")
		}
	})

	t.Run("reports with only one level >3 as a jump are not safe", func(t *testing.T) {
		reportWithOneJumpProblem := Report{levels: []int{1, 3, 8, 9}, safe: false}
		vector := calculateVector(reportWithOneJumpProblem)

		report := ProblemDampener(reportWithOneJumpProblem, vector)

		if report.safe {
			t.Errorf("Expected report to be unsafe, but it was safe")
		}
	})

	t.Run("reports that change direction for only one element are safe", func(t *testing.T) {
		reportWithOneDirectionChange := Report{levels: []int{1, 3, 2, 1}, safe: false}
		vector := calculateVector(reportWithOneDirectionChange)

		report := ProblemDampener(reportWithOneDirectionChange, vector)

		if !report.safe {
			t.Errorf("Expected report to be safe, but it was not")
		}
	})

	t.Run("reports that change direction for only one element and have a big jump on that element are safe", func(t *testing.T) {
		reportWithOneDirectionChange := Report{levels: []int{77, 74, 71, 69, 64, 67}, safe: false}
		vector := calculateVector(reportWithOneDirectionChange)

		report := ProblemDampener(reportWithOneDirectionChange, vector)

		if !report.safe {
			t.Errorf("Expected report to be safe, but it was not")
		}
	})
}
