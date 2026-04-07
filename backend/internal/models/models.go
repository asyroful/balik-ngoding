package models

import (
	"encoding/json"
	"time"
)

type Problem struct {
	ID          string     `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`   // 'loop' | 'string' | 'array' | 'sql'
	Difficulty  string     `json:"difficulty"` // 'easy' | 'medium' | 'hard'
	StarterCode string     `json:"starterCode"`
	Schema      string     `json:"schema,omitempty" gorm:"type:text"`
	IsActive    bool       `json:"isActive" gorm:"default:true"`
	TestCases   []TestCase `json:"testCases,omitempty" gorm:"foreignKey:ProblemID"`
	CreatedAt   time.Time  `json:"createdAt"`
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
	Code         string          `json:"code"`
	Language     string          `json:"language" gorm:"default:'javascript'"`
	Status       string          `json:"status"` // 'accepted' | 'wrong_answer' | 'error'
	Score        int             `json:"score"`
	Total        int             `json:"total"`
	ResultDetail json.RawMessage `json:"resultDetail" gorm:"type:jsonb"`
	CreatedAt    time.Time       `json:"createdAt"`
}
