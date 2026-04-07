package problems

import (
	"balik-ngoding-backend/internal/database"
	"balik-ngoding-backend/internal/models"
	"errors"

	"gorm.io/gorm"
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

	result := query.Find(&problems)
	if result.Error != nil {
		return nil, result.Error
	}

	return problems, nil
}

// FindOne returns a problem by ID with only non-hidden test cases.
// Returns nil (no error) when the problem is not found.
func (s *ProblemsService) FindOne(id string) (*models.Problem, error) {
	var problem models.Problem

	result := database.DB.
		Preload("TestCases", "is_hidden = ?", false).
		First(&problem, "id = ?", id)

	if result.Error != nil {
		// record not found — return nil without error so handler can 404
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &problem, nil
}
