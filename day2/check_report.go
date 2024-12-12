package main

import (
	"math"
)

func CheckReport(report Report) Report {
	reportVector := reportVector{minDistance: 1}

	for i := 1; i < len(report.levels); i++ {
		distance := report.levels[i-1] - report.levels[i]

		reportVector.distances = append(reportVector.distances, distance)

		distanceInAbsolute := int(math.Abs(float64(distance)))

		if distanceInAbsolute > reportVector.maxDistance {
			reportVector.maxDistance = distanceInAbsolute
		}

		if distanceInAbsolute < reportVector.minDistance {
			reportVector.minDistance = distanceInAbsolute
		}

		reportVector.totalDistance += distance
		reportVector.totalAbsDistance += int(math.Abs(float64(distance)))
	}

	if reportVector.maxDistance > 3 || reportVector.minDistance == 0 {
		report.safe = false
	}

	if int(math.Abs(float64(reportVector.totalDistance))) != reportVector.totalAbsDistance {
		report.safe = false
	}

	return report
}

type reportVector = struct {
	distances             []int
	maxDistance           int
	minDistance           int
	allInTheSameDirection bool
	totalDistance         int
	totalAbsDistance      int
}
