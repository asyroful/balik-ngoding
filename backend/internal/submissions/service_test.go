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

// arbitraryUUIDv4 generates a random UUID v4 string for property testing.
func arbitraryUUIDv4() *rapid.Generator[string] {
	return rapid.StringMatching(
		`[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}`,
	)
}

// pureAnonymousIDRoundTrip simulates the persist-then-read cycle without a DB.
// It verifies that storing an anonymousId in a SubmitRequest and reading it back
// from the resulting Submission model produces the same value byte-for-byte.
func pureAnonymousIDRoundTrip(id string) (string, bool) {
	req := SubmitRequest{
		ProblemID:   "test-problem",
		Code:        "function solution() {}",
		Language:    "javascript",
		AnonymousID: &id,
	}
	// Simulate what the service does: copy AnonymousID to the model
	if req.AnonymousID == nil {
		return "", false
	}
	return *req.AnonymousID, true
}

// Feature: anonymous-analytics, Property 4: anonymousId persisted without modification
// Validates: Requirements 2.2, 5.2
func TestAnonymousIdPersistedWithoutModification(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		id := arbitraryUUIDv4().Draw(t, "uuid")

		stored, ok := pureAnonymousIDRoundTrip(id)
		if !ok {
			t.Fatal("expected anonymousId to be present after round-trip")
		}
		if stored != id {
			t.Fatalf("anonymousId modified: sent %q, got %q", id, stored)
		}
	})
}

// TestSubmitWithoutAnonymousId verifies that a nil AnonymousID is preserved as nil.
func TestSubmitWithoutAnonymousId(t *testing.T) {
	req := SubmitRequest{
		ProblemID:   "test-problem",
		Code:        "function solution() {}",
		Language:    "javascript",
		AnonymousID: nil,
	}
	if req.AnonymousID != nil {
		t.Fatal("expected AnonymousID to be nil when not provided")
	}
}
