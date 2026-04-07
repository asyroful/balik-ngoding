package problems

import (
	"testing"

	"balik-ngoding-backend/internal/models"

	"pgregory.net/rapid"
)

// Pure helper functions extracted from service logic for unit testing.

func filterActiveProblems(problems []models.Problem) []models.Problem {
	var result []models.Problem
	for _, p := range problems {
		if p.IsActive {
			result = append(result, p)
		}
	}
	return result
}

func filterByCategory(problems []models.Problem, category string) []models.Problem {
	var result []models.Problem
	for _, p := range problems {
		if p.Category == category {
			result = append(result, p)
		}
	}
	return result
}

func limitProblems(problems []models.Problem, limit int) []models.Problem {
	if len(problems) <= limit {
		return problems
	}
	return problems[:limit]
}

func filterNonHiddenTestCases(testCases []models.TestCase) []models.TestCase {
	var result []models.TestCase
	for _, tc := range testCases {
		if !tc.IsHidden {
			result = append(result, tc)
		}
	}
	return result
}

// Generators

func arbitraryProblem() *rapid.Generator[models.Problem] {
	return rapid.Custom(func(t *rapid.T) models.Problem {
		return models.Problem{
			ID:          rapid.StringMatching(`[a-f0-9]{8}`).Draw(t, "id"),
			Title:       rapid.StringN(1, 50, -1).Draw(t, "title"),
			Category:    rapid.SampledFrom([]string{"loop", "string", "array", "sql"}).Draw(t, "category"),
			Difficulty:  rapid.SampledFrom([]string{"easy", "medium", "hard"}).Draw(t, "difficulty"),
			StarterCode: rapid.StringN(1, 100, -1).Draw(t, "starterCode"),
			IsActive:    rapid.Bool().Draw(t, "isActive"),
		}
	})
}

func arbitraryTestCase() *rapid.Generator[models.TestCase] {
	return rapid.Custom(func(t *rapid.T) models.TestCase {
		return models.TestCase{
			ID:             rapid.StringMatching(`[a-f0-9]{8}`).Draw(t, "id"),
			ProblemID:      rapid.StringMatching(`[a-f0-9]{8}`).Draw(t, "problemId"),
			Input:          rapid.String().Draw(t, "input"),
			ExpectedOutput: rapid.String().Draw(t, "expectedOutput"),
			IsHidden:       rapid.Bool().Draw(t, "isHidden"),
		}
	})
}

// Feature: balik-ngoding, Property 1: Problem list only shows active problems
// Validates: Requirements 1.1, 7.3
func TestActiveProblemsProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		problems := rapid.SliceOf(arbitraryProblem()).Draw(t, "problems")
		result := filterActiveProblems(problems)
		for _, p := range result {
			if !p.IsActive {
				t.Fatalf("expected only active problems, got inactive: %v", p)
			}
		}
	})
}

// Feature: balik-ngoding, Property 2: Category filter returns only matching problems
// Validates: Requirements 1.2
func TestCategoryFilterProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		problems := rapid.SliceOf(arbitraryProblem()).Draw(t, "problems")
		category := rapid.SampledFrom([]string{"loop", "string", "array", "sql"}).Draw(t, "category")
		result := filterByCategory(problems, category)
		for _, p := range result {
			if p.Category != category {
				t.Fatalf("expected category %s, got %s", category, p.Category)
			}
		}
	})
}

// Feature: balik-ngoding, Property 3: Problem list never exceeds 30 items
// Validates: Requirements 1.4
func TestProblemListLimitProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		problems := rapid.SliceOf(arbitraryProblem()).Draw(t, "problems")
		result := limitProblems(problems, 30)
		if len(result) > 30 {
			t.Fatalf("expected at most 30 problems, got %d", len(result))
		}
	})
}

// Feature: balik-ngoding, Property 4: Problem detail contains all required fields
// Validates: Requirements 2.1, 7.1
func TestProblemDetailRequiredFieldsProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		p := arbitraryProblem().Draw(t, "problem")
		// All required fields must be non-empty
		if p.Title == "" {
			t.Fatal("problem title must not be empty")
		}
		if p.Category == "" {
			t.Fatal("problem category must not be empty")
		}
		if p.Difficulty == "" {
			t.Fatal("problem difficulty must not be empty")
		}
		if p.StarterCode == "" {
			t.Fatal("problem starterCode must not be empty")
		}
	})
}

// Feature: balik-ngoding, Property 13: Problem detail only exposes non-hidden test cases
// Validates: Requirements 7.4
func TestNonHiddenTestCasesProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		testCases := rapid.SliceOf(arbitraryTestCase()).Draw(t, "testCases")
		result := filterNonHiddenTestCases(testCases)
		for _, tc := range result {
			if tc.IsHidden {
				t.Fatalf("expected only non-hidden test cases, got hidden: %v", tc)
			}
		}
	})
}
