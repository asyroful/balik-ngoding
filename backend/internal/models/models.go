package models

import (
	"encoding/json"
	"time"
)

type Problem struct {
	ID                string          `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	Category          string          `json:"category"`   // 'loop' | 'string' | 'array' | 'sql'
	Difficulty        string          `json:"difficulty"` // 'easy' | 'medium' | 'hard'
	StarterCode       string          `json:"starterCode"`
	Schema            string          `json:"schema,omitempty" gorm:"type:text"`
	IsActive          bool            `json:"isActive" gorm:"default:true"`
	ThinkingGuide     *string         `json:"thinkingGuide" gorm:"type:text"`
	Hints             json.RawMessage `json:"hints" gorm:"type:jsonb"`
	PrerequisiteID    *string         `json:"prerequisiteId" gorm:"type:uuid;index"`
	PrerequisiteTitle *string         `json:"prerequisiteTitle" gorm:"-"`
	TestCases         []TestCase      `json:"testCases,omitempty" gorm:"foreignKey:ProblemID"`
	CreatedAt         time.Time       `json:"createdAt"`
}

type TestCase struct {
	ID             string `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProblemID      string `json:"problemId" gorm:"type:uuid;not null;index:idx_test_cases_problem_id"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	IsHidden       bool   `json:"isHidden" gorm:"default:false"`
}

type Submission struct {
	ID           string          `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProblemID    string          `json:"problemId" gorm:"type:uuid;not null;index:idx_submissions_problem_id"`
	AnonymousID  *string         `json:"anonymousId" gorm:"type:varchar(36);index"`
	CodePath     *string         `json:"codePath" gorm:"type:varchar(255)"` // File path to stored code
	Language     string          `json:"language" gorm:"default:'javascript'"`
	Status       string          `json:"status"` // 'accepted' | 'wrong_answer' | 'error'
	Score        int             `json:"score"`
	Total        int             `json:"total"`
	ResultDetail json.RawMessage `json:"resultDetail" gorm:"type:jsonb"`
	CreatedAt    time.Time       `json:"createdAt"`
}

type SolutionKey struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProblemID string    `json:"problemId" gorm:"type:uuid;not null;index:idx_solution_keys_problem_id;uniqueIndex:idx_solution_keys_unique"`
	Code      string    `json:"code" gorm:"type:text;not null"`
	Language  string    `json:"language" gorm:"not null;uniqueIndex:idx_solution_keys_unique"` // 'javascript' | 'sql'
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Foreign key relationship
	Problem *Problem `json:"-" gorm:"foreignKey:ProblemID;constraint:OnDelete:CASCADE"`
}

// CategorySummary is the response struct for GET /problems/summary.
// Not persisted to database — used only as a DTO.
type CategorySummary struct {
	Category string `json:"category"`
	Total    int    `json:"total"`
}
