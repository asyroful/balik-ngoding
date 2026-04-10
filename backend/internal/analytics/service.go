package analytics

import (
	"balik-ngoding-backend/internal/database"
)

// TopProblem is a DTO for a problem with its submission count.
type TopProblem struct {
	ProblemID       string `json:"problemId"`
	Title           string `json:"title"`
	SubmissionCount int    `json:"submissionCount"`
}

// StatsResult is the response DTO for GET /analytics/stats.
type StatsResult struct {
	UniqueDevices       int          `json:"uniqueDevices"`
	TotalSubmissions    int          `json:"totalSubmissions"`
	TotalAccepted       int          `json:"totalAccepted"`
	DevicesWithAccepted int          `json:"devicesWithAccepted"`
	TopProblems         []TopProblem `json:"topProblems"`
}

// AnalyticsService handles analytics aggregation logic.
type AnalyticsService struct{}

// NewAnalyticsService creates a new AnalyticsService.
func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{}
}

// GetStats runs 5 aggregation queries and returns the combined result.
func (s *AnalyticsService) GetStats() (*StatsResult, error) {
	db := database.DB

	var uniqueDevices int
	if err := db.Raw(
		"SELECT COUNT(DISTINCT anonymous_id) FROM submissions WHERE anonymous_id IS NOT NULL",
	).Scan(&uniqueDevices).Error; err != nil {
		return nil, err
	}

	var totalSubmissions int
	if err := db.Raw("SELECT COUNT(*) FROM submissions").Scan(&totalSubmissions).Error; err != nil {
		return nil, err
	}

	var totalAccepted int
	if err := db.Raw(
		"SELECT COUNT(*) FROM submissions WHERE status = 'accepted'",
	).Scan(&totalAccepted).Error; err != nil {
		return nil, err
	}

	var devicesWithAccepted int
	if err := db.Raw(
		"SELECT COUNT(DISTINCT anonymous_id) FROM submissions WHERE status = 'accepted' AND anonymous_id IS NOT NULL",
	).Scan(&devicesWithAccepted).Error; err != nil {
		return nil, err
	}

	var topProblems []TopProblem
	if err := db.Raw(`
		SELECT s.problem_id AS problem_id, p.title AS title, COUNT(*) AS submission_count
		FROM submissions s
		JOIN problems p ON p.id = s.problem_id
		GROUP BY s.problem_id, p.title
		ORDER BY submission_count DESC
		LIMIT 5
	`).Scan(&topProblems).Error; err != nil {
		return nil, err
	}

	if topProblems == nil {
		topProblems = []TopProblem{}
	}

	return &StatsResult{
		UniqueDevices:       uniqueDevices,
		TotalSubmissions:    totalSubmissions,
		TotalAccepted:       totalAccepted,
		DevicesWithAccepted: devicesWithAccepted,
		TopProblems:         topProblems,
	}, nil
}
