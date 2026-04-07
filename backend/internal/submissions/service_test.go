package submissions

import (
	"testing"

	"pgregory.net/rapid"
)

// arbitraryTestCaseResult generates a random TestCaseResult for property testing.
func arbitraryTestCaseResult() *rapid.Generator[TestCaseResult] {
	return rapid.Custom(func(t *rapid.T) TestCaseResult {
		passed := rapid.Bool().Draw(t, "passed")
		errMsg := ""
		if rapid.Bool().Draw(t, "hasError") {
			errMsg = rapid.StringN(1, 50, -1).Draw(t, "error")
			passed = false
		}
		return TestCaseResult{
			Passed:   passed,
			Input:    rapid.String().Draw(t, "input"),
			Expected: rapid.String().Draw(t, "expected"),
			Actual:   rapid.String().Draw(t, "actual"),
			Error:    errMsg,
		}
	})
}

// Feature: balik-ngoding, Property 14: Submission is persisted with all required fields
// Validates: Requirements 7.6
//
// This test verifies the pure determineStatus function which drives the Status field
// stored on every persisted Submission. It checks that:
//   - status is "error"        when any result has a non-empty Error
//   - status is "accepted"     when score == total (all passed, no errors)
//   - status is "wrong_answer" otherwise
func TestDetermineStatusProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		results := rapid.SliceOf(arbitraryTestCaseResult()).Draw(t, "results")

		hasError := false
		score := 0
		for _, r := range results {
			if r.Error != "" {
				hasError = true
			}
			if r.Passed {
				score++
			}
		}
		total := len(results)

		status := determineStatus(results, hasError, score, total)

		switch {
		case hasError:
			if status != "error" {
				t.Fatalf("expected 'error' when hasError=true, got %q", status)
			}
		case score == total:
			if status != "accepted" {
				t.Fatalf("expected 'accepted' when score==total (%d==%d), got %q", score, total, status)
			}
		default:
			if status != "wrong_answer" {
				t.Fatalf("expected 'wrong_answer' when score(%d) < total(%d), got %q", score, total, status)
			}
		}
	})
}
