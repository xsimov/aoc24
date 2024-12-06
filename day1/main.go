package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Location struct {
	firstValue  int64
	secondValue int64
	distance    int64
}

type (
	Locations []Location
	Column    []int64
)

func main() {
	content := parseFile("./data.txt")
	rows := strings.Split(content, "\n")
	var (
		firstColumn  Column
		secondColumn Column
	)

	for _, row := range rows {
		values := strings.Fields(row)

		firstValue, err := strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		secondValue, err := strconv.ParseInt(values[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		firstColumn = append(firstColumn, firstValue)
		secondColumn = append(secondColumn, secondValue)
	}

	distance := calculateDistanceBetweenColumns(firstColumn, secondColumn)
	log.Printf("Total Distance: %d", distance)

	similarityScore := calculateSimilarityScore(firstColumn, secondColumn)
	log.Printf("Similarity Score: %d", similarityScore)
}

func calculateDistanceBetweenColumns(firstColumn Column, secondColumn Column) int64 {
	var locations Locations

	sort.SliceStable(firstColumn, func(i, j int) bool {
		return int(firstColumn[i]) < int(firstColumn[j])
	})
	sort.SliceStable(secondColumn, func(i, j int) bool {
		return int(secondColumn[i]) < int(secondColumn[j])
	})

	for i := 0; i < len(firstColumn); i++ {
		var distance int64
		firstValue := firstColumn[i]
		secondValue := secondColumn[i]

		if firstValue < secondValue {
			distance = secondValue - firstValue
		} else {
			distance = firstValue - secondValue
		}

		locations = append(locations, Location{
			firstValue:  firstValue,
			secondValue: secondValue,
			distance:    distance,
		})
	}

	var totalDistance int64 = 0

	for _, location := range locations {
		totalDistance += location.distance
	}

	return totalDistance
}

func parseFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(bytes)
}

func calculateSimilarityScore(firstColumn Column, secondColumn Column) int64 {
	scoresForSecondColumn := make(map[int64]int64, len(secondColumn))

	for i := 0; i < len(secondColumn); i++ {
		scoresForSecondColumn[secondColumn[i]]++
	}

	var similarityScore int64 = 0

	for i := 0; i < len(firstColumn); i++ {
		firstValue := firstColumn[i]

		similarityScore += (scoresForSecondColumn[firstValue] * firstValue)
	}

	return similarityScore
}
