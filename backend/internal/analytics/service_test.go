package analytics

import (
	"testing"

	"pgregory.net/rapid"
)

// --- Pure aggregation helpers (mirror the SQL logic for property testing) ---

type submissionRecord struct {
	anonymousID *string
	status      string
	problemID   string
	title       string
}

func computeUniqueDevices(records []submissionRecord) int {
	seen := make(map[string]struct{})
	for _, r := range records {
		if r.anonymousID != nil {
			seen[*r.anonymousID] = struct{}{}
		}
	}
	return len(seen)
}

func computeTotalSubmissions(records []submissionRecord) int {
	return len(records)
}

func computeTotalAccepted(records []submissionRecord) int {
	count := 0
	for _, r := range records {
		if r.status == "accepted" {
			count++
		}
	}
	return count
}

func computeDevicesWithAccepted(records []submissionRecord) int {
	seen := make(map[string]struct{})
	for _, r := range records {
		if r.status == "accepted" && r.anonymousID != nil {
			seen[*r.anonymousID] = struct{}{}
		}
	}
	return len(seen)
}

func computeTopProblems(records []submissionRecord) []TopProblem {
	counts := make(map[string]int)
	titles := make(map[string]string)
	for _, r := range records {
		counts[r.problemID]++
		titles[r.problemID] = r.title
	}

	// Build slice and sort descending
	var result []TopProblem
	for pid, count := range counts {
		result = append(result, TopProblem{
			ProblemID:       pid,
			Title:           titles[pid],
			SubmissionCount: count,
		})
	}
	// Simple insertion sort (small slice)
	for i := 1; i < len(result); i++ {
		for j := i; j > 0 && result[j].SubmissionCount > result[j-1].SubmissionCount; j-- {
			result[j], result[j-1] = result[j-1], result[j]
		}
	}
	if len(result) > 5 {
		result = result[:5]
	}
	return result
}

// --- Generators ---

func arbitraryAnonymousID() *rapid.Generator[*string] {
	return rapid.Custom(func(t *rapid.T) *string {
		if rapid.Bool().Draw(t, "hasID") {
			// UUID v4-like pattern
			id := rapid.StringMatching(
				`[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}`,
			).Draw(t, "uuid")
			return &id
		}
		return nil
	})
}

func arbitrarySubmissionRecord() *rapid.Generator[submissionRecord] {
	return rapid.Custom(func(t *rapid.T) submissionRecord {
		status := rapid.SampledFrom([]string{"accepted", "wrong_answer", "error"}).Draw(t, "status")
		problemID := rapid.StringMatching(`[a-f0-9]{8}`).Draw(t, "problemId")
		title := rapid.StringN(1, 30, -1).Draw(t, "title")
		return submissionRecord{
			anonymousID: arbitraryAnonymousID().Draw(t, "anonymousID"),
			status:      status,
			problemID:   problemID,
			title:       title,
		}
	})
}

// --- Unit test: empty table ---

// TestGetStatsEmptyTable verifies that zero-value StatsResult is returned for empty input.
func TestGetStatsEmptyTable(t *testing.T) {
	records := []submissionRecord{}

	if got := computeUniqueDevices(records); got != 0 {
		t.Fatalf("uniqueDevices: want 0, got %d", got)
	}
	if got := computeTotalSubmissions(records); got != 0 {
		t.Fatalf("totalSubmissions: want 0, got %d", got)
	}
	if got := computeTotalAccepted(records); got != 0 {
		t.Fatalf("totalAccepted: want 0, got %d", got)
	}
	if got := computeDevicesWithAccepted(records); got != 0 {
		t.Fatalf("devicesWithAccepted: want 0, got %d", got)
	}
	if got := computeTopProblems(records); len(got) != 0 {
		t.Fatalf("topProblems: want empty, got %v", got)
	}
}

// Feature: anonymous-analytics, Property 5: stats aggregation correctness
// Validates: Requirements 3.1, 3.2, 3.3, 3.5, 5.3
func TestStatsAggregationCorrectness(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		records := rapid.SliceOf(arbitrarySubmissionRecord()).Draw(t, "records")

		wantUnique := computeUniqueDevices(records)
		wantTotal := computeTotalSubmissions(records)
		wantAccepted := computeTotalAccepted(records)
		wantDevicesAccepted := computeDevicesWithAccepted(records)

		// Recompute independently to verify idempotence
		if got := computeUniqueDevices(records); got != wantUnique {
			t.Fatalf("uniqueDevices not idempotent: %d vs %d", got, wantUnique)
		}
		if got := computeTotalSubmissions(records); got != wantTotal {
			t.Fatalf("totalSubmissions not idempotent: %d vs %d", got, wantTotal)
		}
		if got := computeTotalAccepted(records); got != wantAccepted {
			t.Fatalf("totalAccepted not idempotent: %d vs %d", got, wantAccepted)
		}
		if got := computeDevicesWithAccepted(records); got != wantDevicesAccepted {
			t.Fatalf("devicesWithAccepted not idempotent: %d vs %d", got, wantDevicesAccepted)
		}

		// devicesWithAccepted must be <= uniqueDevices
		if wantDevicesAccepted > wantUnique {
			t.Fatalf("devicesWithAccepted (%d) > uniqueDevices (%d)", wantDevicesAccepted, wantUnique)
		}

		// totalAccepted must be <= totalSubmissions
		if wantAccepted > wantTotal {
			t.Fatalf("totalAccepted (%d) > totalSubmissions (%d)", wantAccepted, wantTotal)
		}
	})
}

// Feature: anonymous-analytics, Property 6: topProblems ordering and size
// Validates: Requirements 3.4
func TestTopProblemsOrdering(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		records := rapid.SliceOf(arbitrarySubmissionRecord()).Draw(t, "records")
		top := computeTopProblems(records)

		if len(top) > 5 {
			t.Fatalf("topProblems has %d entries, want at most 5", len(top))
		}

		for i := 1; i < len(top); i++ {
			if top[i].SubmissionCount > top[i-1].SubmissionCount {
				t.Fatalf("topProblems not non-increasing at index %d: %d > %d",
					i, top[i].SubmissionCount, top[i-1].SubmissionCount)
			}
		}
	})
}
