package main

import (
	"math"
)

func ProblemDampener(report Report, vector ReportVector) Report {
	if report.problemDampenerApplied {
		return report
	}

	problematicLevelsIndexes := []int{}
	directions := map[string]int{}

	for index, distance := range vector.distances {
		if math.Abs(float64(distance)) > 3 || distance == 0 {
			problematicLevelsIndexes = append(problematicLevelsIndexes, index, index+1)
		}

		directions[getDirection(distance)]++
	}

	if directions["up"] > 1 && directions["down"] == 1 {
		for index, distance := range vector.distances {
			if distance > 0 {
				problematicLevelsIndexes = append(problematicLevelsIndexes, index, index+1)
			}
		}
	}

	if directions["up"] == 1 && directions["down"] > 1 {
		for index, distance := range vector.distances {
			if distance < 0 {
				problematicLevelsIndexes = append(problematicLevelsIndexes, index, index+1)
			}
		}
	}

	uniqueProblematicLevelsIndexes := uniq(problematicLevelsIndexes)
	if len(uniqueProblematicLevelsIndexes) == 0 {
		return report
	}

	var newReport, reportWithoutProblematicLevel Report

	for _, index := range uniqueProblematicLevelsIndexes {
		reportWithoutProblematicLevel = removeProblematicLevel(report, index)

		newReport = CheckReport(reportWithoutProblematicLevel)

		if newReport.safe {
			newReport.problemDampenerApplied = true

			return newReport
		}
	}
	return CheckReport(newReport)
}

func removeProblematicLevel(report Report, index int) Report {
	newLevels := []int{}
	for i, level := range report.levels {
		if i != index {
			newLevels = append(newLevels, level)
		}
	}
	return Report{levels: newLevels, safe: true, problemDampenerApplied: true}
}

func getDirection(distance int) string {
	if distance < 0 {
		return "up"
	} else {
		return "down"
	}
}

func uniq(entryList []int) (list []int) {
	appended := map[int]bool{}
	for _, element := range entryList {
		if !appended[element] {
			list = append(list, element)
			appended[element] = true
		}
	}

	return list
}
