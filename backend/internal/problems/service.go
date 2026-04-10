package problems

import (
	"balik-ngoding-backend/internal/database"
	"balik-ngoding-backend/internal/models"
)

// ProblemsService handles business logic for problems.
type ProblemsService struct{}

// FindAll returns active problems, optionally filtered by category and/or difficulty.
func (s *ProblemsService) FindAll(category, difficulty string) ([]models.Problem, error) {
	var problems []models.Problem

	query := database.DB.Where("is_active = ?", true)

	if category != "" {
		query = query.Where("category = ?", category)
	}
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}

	query = query.Order("CASE difficulty WHEN 'easy' THEN 1 WHEN 'medium' THEN 2 WHEN 'hard' THEN 3 ELSE 4 END, created_at ASC")

	result := query.Find(&problems)
	if result.Error != nil {
		return nil, result.Error
	}

	return problems, nil
}

// FindSummary returns the count of active problems per category.
// Always returns entries for all 4 categories: loop, string, array, sql.
func (s *ProblemsService) FindSummary() ([]models.CategorySummary, error) {
	var rows []models.CategorySummary
	result := database.DB.Raw(
		"SELECT category, COUNT(*) as total FROM problems WHERE is_active = true GROUP BY category",
	).Scan(&rows)
	if result.Error != nil {
		return nil, result.Error
	}

	counts := map[string]int{}
	for _, row := range rows {
		counts[row.Category] = row.Total
	}

	categories := []string{"loop", "string", "array", "sql"}
	summary := make([]models.CategorySummary, len(categories))
	for i, cat := range categories {
		summary[i] = models.CategorySummary{Category: cat, Total: counts[cat]}
	}
	return summary, nil
}

// FindOne returns a problem by ID with only non-hidden test cases.
// Returns nil (no error) when the problem is not found.
// Uses a self-join to populate PrerequisiteTitle from the related problem.
func (s *ProblemsService) FindOne(id string) (*models.Problem, error) {
	var problem models.Problem

	result := database.DB.Raw(
		`SELECT p.*, prereq.title AS prerequisite_title
		 FROM problems p
		 LEFT JOIN problems prereq ON prereq.id = p.prerequisite_id
		 WHERE p.id = ?`,
		id,
	).Scan(&problem)

	if result.Error != nil {
		return nil, result.Error
	}

	// RowsAffected == 0 means no record found — return nil so handler can 404
	if result.RowsAffected == 0 {
		return nil, nil
	}

	// Preload non-hidden test cases separately
	if err := database.DB.
		Where("problem_id = ? AND is_hidden = ?", problem.ID, false).
		Find(&problem.TestCases).Error; err != nil {
		return nil, err
	}

	return &problem, nil
}
