package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
)

type operation struct {
	instruction string
	x           int
	y           int
	position    int
}

func String(o operation) string {
	if o.instruction == "mul" {
		return fmt.Sprintf("instruction: %s, x: %d, y: %d, position: %d", o.instruction, o.x, o.y, o.position)
	} else {
		return fmt.Sprintf("instruction: %s, position: %d", o.instruction, o.position)
	}
}

var operations []operation

func main() {
	fileContent := parseFile("data.txt")
	mulReg := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	mulMatches := mulReg.FindAllStringSubmatchIndex(fileContent, -1)

	for _, match := range mulMatches {
		xStr := fileContent[match[2]:match[3]]
		x, err := strconv.ParseInt(xStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		yStr := fileContent[match[4]:match[5]]
		y, err := strconv.ParseInt(yStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		o := operation{x: int(x), y: int(y), position: match[0], instruction: "mul"}
		operations = append(operations, o)
	}

	dosDontsReg := regexp.MustCompile(`do(n't)?\(\)`)
	dosDontsMatches := dosDontsReg.FindAllStringSubmatchIndex(fileContent, -1)

	var instruction string
	for _, match := range dosDontsMatches {
		if match[2] == -1 {
			instruction = "do"
		} else {
			instruction = "don't"
		}
		o := operation{position: match[0], instruction: instruction}

		operations = append(operations, o)
	}

	slices.SortStableFunc[[]operation, operation](operations, func(a, b operation) int {
		return a.position - b.position
	})

	accum := 0
	enabled := true

	for _, op := range operations {
		switch op.instruction {
		case "mul":
			if enabled {
				accum += op.x * op.y
			}
		case "do":
			enabled = true
		case "don't":
			enabled = false
		}
	}
	fmt.Println(accum)
}
