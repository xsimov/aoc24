package main

import "math"

func CheckReport(report Report) Report {
	var initialDirection string
	reportLevels := report.levels

	if len(reportLevels) < 2 {
		report.safe = false
		return report
	}

	if reportLevels[0] > reportLevels[1] {
		initialDirection = "down"
	} else {
		initialDirection = "up"
	}

	for i := 1; i < len(reportLevels); i++ {
		distance := math.Abs(float64(reportLevels[i-1] - reportLevels[i]))

		if distance > 3 || distance == 0 {
			report.safe = false
		}

		if initialDirection == "down" {
			if reportLevels[i-1] > reportLevels[i] {
				report.safe = report.safe && true
			}

			if reportLevels[i-1] < reportLevels[i] {
				report.errorLevelPositions = append(report.errorLevelPositions, i)
				report.safe = false
			}
		}

		if initialDirection == "up" {
			if reportLevels[i-1] > reportLevels[i] {
				report.errorLevelPositions = append(report.errorLevelPositions, i)
				report.safe = false
			}

			if reportLevels[i-1] < reportLevels[i] {
				report.safe = report.safe && true
			}
		}
	}
	return report
}
