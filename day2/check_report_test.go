package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("expect upwards direction distance >3 to fail", func(t *testing.T) {
		report := Report{levels: []int{1, 2, 6, 7, 12}, safe: true}

		result := CheckReport(report)

		if result.safe {
			t.Fatalf("Expected false, got %v", result.safe)
		}
	})

	t.Run("expect correct Report to be safe in upwards direction", func(t *testing.T) {
		report := Report{levels: []int{1, 2, 3, 6, 8}, safe: true}

		result := CheckReport(report)

		if !result.safe {
			t.Fatalf("Expected report to be safe, got `safe: %v`", result.safe)
		}
	})

	t.Run("uniq works as expected", func(t *testing.T) {
		duplicates := []int{1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 9}
		expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		result := uniq(duplicates)

		if !reflect.DeepEqual(result, expectedResult) {
			t.Fatalf("Expected %v, got %v", expectedResult, result)
		}
	})
}
