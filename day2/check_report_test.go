package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("expect upwards direction distance >3 to fail", func(t *testing.T) {
		report := Report{levels: []int{1, 2, 6, 7, 9}, safe: true}

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

	t.Run("expect report with change in direction to be unsafe", func(t *testing.T) {
		report := Report{levels: []int{2, 3, 2, 1}, safe: true}

		result := CheckReport(report)

		if result.safe {
			t.Fatalf("Expected report not to be safe, got `safe: %v`", result.safe)
		}
	})
}
