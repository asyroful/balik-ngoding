package database

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"

	"balik-ngoding-backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:embed seeds/*.json
var seedFiles embed.FS

var seedFileNames = []string{
	"seeds/loop.json",
	"seeds/string.json",
	"seeds/array.json",
	"seeds/sql.json",
}

// SeedProblem is the JSON representation of a problem in seed files.
// Different from models.Problem because hints in JSON are plain []string.
type SeedProblem struct {
	ID             string         `json:"id"`
	Title          string         `json:"title"`
	Category       string         `json:"category"`
	Difficulty     string         `json:"difficulty"`
	Description    string         `json:"description"`
	StarterCode    string         `json:"starterCode"`
	Schema         string         `json:"schema,omitempty"`
	ThinkingGuide  *string        `json:"thinkingGuide"`
	Hints          []string       `json:"hints"`
	PrerequisiteID *string        `json:"prerequisiteId"`
	TestCases      []SeedTestCase `json:"testCases"`
}

type SeedTestCase struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	IsHidden       bool   `json:"isHidden"`
}

// loadProblemsFromFile reads and validates problems from a JSON file in the embed.FS.
func loadProblemsFromFile(fs embed.FS, filename string) ([]SeedProblem, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", filename, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}

	var problems []SeedProblem
	if err := json.Unmarshal(data, &problems); err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", filename, err)
	}

	for _, p := range problems {
		// Validate required fields
		if p.ID == "" {
			return nil, fmt.Errorf("%s: problem has empty id", filename)
		}
		if p.Title == "" {
			return nil, fmt.Errorf("%s: problem %q has empty title", filename, p.ID)
		}
		if p.Category == "" {
			return nil, fmt.Errorf("%s: problem %q has empty category", filename, p.ID)
		}
		if p.Difficulty == "" {
			return nil, fmt.Errorf("%s: problem %q has empty difficulty", filename, p.ID)
		}
		if p.Description == "" {
			return nil, fmt.Errorf("%s: problem %q has empty description", filename, p.ID)
		}
		if p.StarterCode == "" && p.Category != "sql" {
			return nil, fmt.Errorf("%s: problem %q has empty starterCode", filename, p.ID)
		}

		// Validate hints
		if len(p.Hints) != 3 {
			return nil, fmt.Errorf("%s: problem %q must have exactly 3 hints, got %d", filename, p.ID, len(p.Hints))
		}

		// Validate thinkingGuide
		if p.ThinkingGuide == nil || *p.ThinkingGuide == "" {
			return nil, fmt.Errorf("%s: problem %q has null or empty thinkingGuide", filename, p.ID)
		}

		// Validate test cases
		visible := 0
		hidden := 0
		for _, tc := range p.TestCases {
			if tc.IsHidden {
				hidden++
			} else {
				visible++
			}
		}
		if visible < 3 {
			return nil, fmt.Errorf("%s: problem %q has only %d visible test cases (need >= 3)", filename, p.ID, visible)
		}
		if hidden < 2 {
			return nil, fmt.Errorf("%s: problem %q has only %d hidden test cases (need >= 2)", filename, p.ID, hidden)
		}
	}

	return problems, nil
}

// SeedFromJSON loads problems from embedded JSON files and upserts them into the database.
// It first deletes all existing problems (and cascades to test cases) to ensure a clean state.
func SeedFromJSON(db *gorm.DB) error {
	// Delete all test cases and problems to remove stale data from old seeds
	if err := db.Exec("DELETE FROM test_cases").Error; err != nil {
		return fmt.Errorf("failed to truncate test_cases: %w", err)
	}
	if err := db.Exec("DELETE FROM problems").Error; err != nil {
		return fmt.Errorf("failed to truncate problems: %w", err)
	}

	for _, filename := range seedFileNames {
		problems, err := loadProblemsFromFile(seedFiles, filename)
		if err != nil {
			return fmt.Errorf("seed error in %s: %w", filename, err)
		}

		for _, sp := range problems {
			hintsJSON, err := json.Marshal(sp.Hints)
			if err != nil {
				return fmt.Errorf("failed to marshal hints for problem %q: %w", sp.ID, err)
			}

			problem := models.Problem{
				ID:             sp.ID,
				Title:          sp.Title,
				Category:       sp.Category,
				Difficulty:     sp.Difficulty,
				Description:    sp.Description,
				StarterCode:    sp.StarterCode,
				Schema:         sp.Schema,
				ThinkingGuide:  sp.ThinkingGuide,
				Hints:          json.RawMessage(hintsJSON),
				PrerequisiteID: sp.PrerequisiteID,
				IsActive:       true,
			}

			result := db.Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "id"}},
				DoUpdates: clause.AssignmentColumns([]string{
					"title", "description", "category", "difficulty",
					"starter_code", "schema", "thinking_guide", "hints",
					"prerequisite_id", "is_active",
				}),
			}).Create(&problem)
			if result.Error != nil {
				return fmt.Errorf("failed to upsert problem %q: %w", sp.ID, result.Error)
			}

			// Delete existing test cases then re-insert
			if err := db.Where("problem_id = ?", sp.ID).Delete(&models.TestCase{}).Error; err != nil {
				return fmt.Errorf("failed to delete test cases for problem %q: %w", sp.ID, err)
			}

			for _, stc := range sp.TestCases {
				tc := models.TestCase{
					ProblemID:      sp.ID,
					Input:          stc.Input,
					ExpectedOutput: stc.ExpectedOutput,
					IsHidden:       stc.IsHidden,
				}
				if err := db.Create(&tc).Error; err != nil {
					return fmt.Errorf("failed to insert test case for problem %q: %w", sp.ID, err)
				}
			}
		}
	}

	return nil
}
