package main

import (
	"fmt"
	"math"
)

func CheckReport(report Report) Report {
	vector := calculateVector(report)

	if vector.maxDistance > 3 || vector.minDistance == 0 {
		report.safe = false
	}

	if int(math.Abs(float64(vector.totalDistance))) != vector.totalAbsDistance {
		report.safe = false
	}

	if report.safe {
		return report
	}

	afterProblemDampener := ProblemDampener(report, vector)

	return afterProblemDampener
}

func calculateVector(report Report) ReportVector {
	vector := ReportVector{minDistance: 1}

	for i := 1; i < len(report.levels); i++ {
		distance := report.levels[i-1] - report.levels[i]

		vector.distances = append(vector.distances, distance)

		distanceInAbsolute := int(math.Abs(float64(distance)))

		if distanceInAbsolute > vector.maxDistance {
			vector.maxDistance = distanceInAbsolute
		}

		if distanceInAbsolute < vector.minDistance {
			vector.minDistance = distanceInAbsolute
		}

		vector.totalDistance += distance
		vector.totalAbsDistance += int(math.Abs(float64(distance)))
	}

	return vector
}

type ReportVector struct {
	distances             []int
	maxDistance           int
	minDistance           int
	totalDistance         int
	totalAbsDistance      int
	allInTheSameDirection bool
}

func (r ReportVector) String() string {
	return fmt.Sprintln("{\ndistances: ", r.distances,
		",\nmaxDistance: ", r.maxDistance,
		",\nminDistance: ", r.minDistance,
		",\ntotalDistance: ", r.totalDistance,
		",\ntotalAbsDistance: ", r.totalAbsDistance,
		",\nallInTheSameDirection: ", r.allInTheSameDirection, "\n}")
}
