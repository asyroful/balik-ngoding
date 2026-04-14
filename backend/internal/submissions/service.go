package submissions

import (
	"encoding/json"
	"errors"
	"fmt"

	"balik-ngoding-backend/internal/database"
	"balik-ngoding-backend/internal/evaluator"
	"balik-ngoding-backend/internal/models"
	"balik-ngoding-backend/internal/storage"

	"gorm.io/gorm"
)

// SubmitRequest holds the input for a submission.
type SubmitRequest struct {
	ProblemID   string  `json:"problemId"`
	Code        string  `json:"code"`
	Language    string  `json:"language"`
	AnonymousID *string `json:"anonymousId"` // nullable — not required
}

// TestCaseResult is the per-test-case result DTO.
type TestCaseResult struct {
	Passed   bool   `json:"passed"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	Error    string `json:"error,omitempty"`
}

// SubmissionResult is the response DTO returned to the client.
type SubmissionResult struct {
	Status       string           `json:"status"`
	Score        int              `json:"score"`
	Total        int              `json:"total"`
	Results      []TestCaseResult `json:"results"`
	ErrorMessage string           `json:"errorMessage,omitempty"`
}

// SubmissionsService handles submission business logic.
type SubmissionsService struct {
	evaluator   *evaluator.EvaluatorService
	fileStorage *storage.FileStorageService
}

// NewSubmissionsService creates a new SubmissionsService.
func NewSubmissionsService(fileStorage *storage.FileStorageService) *SubmissionsService {
	return &SubmissionsService{
		evaluator:   evaluator.NewEvaluatorService(),
		fileStorage: fileStorage,
	}
}

// Submit evaluates the code against all test cases for the given problem,
// persists the submission, and returns the result.
func (s *SubmissionsService) Submit(req SubmitRequest) (*SubmissionResult, error) {
	// Fetch all test cases (including hidden) for the problem
	var testCases []models.TestCase
	result := database.DB.Where("problem_id = ?", req.ProblemID).Find(&testCases)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch test cases: %w", result.Error)
	}

	// Verify the problem exists
	var problem models.Problem
	res := database.DB.First(&problem, "id = ?", req.ProblemID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil // caller will 404
		}
		return nil, fmt.Errorf("failed to fetch problem: %w", res.Error)
	}

	// Validate SQL problems have a schema
	if problem.Category == "sql" && problem.Schema == "" {
		return nil, fmt.Errorf("Schema soal SQL tidak ditemukan")
	}

	// Run evaluator against each test case
	tcResults := make([]TestCaseResult, 0, len(testCases))
	score := 0
	hasError := false

	for _, tc := range testCases {
		evalResult := s.evaluator.EvaluateWithLanguage(req.Language, problem.Schema, req.Code, tc.Input, tc.ExpectedOutput)

		tcr := TestCaseResult{
			Passed:   evalResult.Passed,
			Input:    tc.Input,
			Expected: tc.ExpectedOutput,
			Actual:   evalResult.Actual,
			Error:    evalResult.Error,
		}

		if evalResult.Error != "" {
			hasError = true
		}
		if evalResult.Passed {
			score++
		}

		tcResults = append(tcResults, tcr)
	}

	total := len(testCases)

	// Determine overall status
	status := determineStatus(tcResults, hasError, score, total)

	// Build error message if applicable
	errorMessage := ""
	if hasError {
		for _, r := range tcResults {
			if r.Error != "" {
				errorMessage = r.Error
				break
			}
		}
	}

	// Serialize result detail for persistence
	resultDetailJSON, err := json.Marshal(tcResults)
	if err != nil {
		return nil, err
	}

	// Save code to file storage
	codePath, err := s.fileStorage.SaveCode(req.ProblemID, req.Language, req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to save code: %w", err)
	}

	// Persist submission with code path reference
	submission := models.Submission{
		ProblemID:    req.ProblemID,
		AnonymousID:  req.AnonymousID,
		CodePath:     codePath,
		Language:     req.Language,
		Status:       status,
		Score:        score,
		Total:        total,
		ResultDetail: json.RawMessage(resultDetailJSON),
	}
	if err := database.DB.Create(&submission).Error; err != nil {
		return nil, fmt.Errorf("failed to create submission: %w", err)
	}

	return &SubmissionResult{
		Status:       status,
		Score:        score,
		Total:        total,
		Results:      tcResults,
		ErrorMessage: errorMessage,
	}, nil
}

// determineStatus returns the overall submission status based on results.
func determineStatus(_ []TestCaseResult, hasError bool, score, total int) string {
	if hasError {
		return "error"
	}
	if score == total {
		return "accepted"
	}
	return "wrong_answer"
}
