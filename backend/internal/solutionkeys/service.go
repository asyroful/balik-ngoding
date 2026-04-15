package solutionkeys

import (
	"errors"
	"fmt"

	"balik-ngoding-backend/internal/models"

	"gorm.io/gorm"
)

// CreateSolutionKeyRequest is the request struct for creating a solution key
type CreateSolutionKeyRequest struct {
	ProblemID string `json:"problemId" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Language  string `json:"language" binding:"required"`
}

// UpdateSolutionKeyRequest is the request struct for updating a solution key
type UpdateSolutionKeyRequest struct {
	Code     string `json:"code" binding:"required"`
	Language string `json:"language"`
}

// SolutionKeyService handles business logic for solution keys
type SolutionKeyService struct {
	db *gorm.DB
}

// NewSolutionKeyService creates a new SolutionKeyService
func NewSolutionKeyService(db *gorm.DB) *SolutionKeyService {
	return &SolutionKeyService{db: db}
}

// CreateSolutionKey creates a new solution key with validation
func (s *SolutionKeyService) CreateSolutionKey(req CreateSolutionKeyRequest) (*models.SolutionKey, error) {
	// Validate request
	if err := s.ValidateSolutionKeyRequest(req); err != nil {
		return nil, err
	}

	// Check if problem exists
	var problem models.Problem
	if result := s.db.First(&problem, "id = ?", req.ProblemID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("problem not found")
		}
		return nil, result.Error
	}

	// Create solution key
	sk := &models.SolutionKey{
		ProblemID: req.ProblemID,
		Code:      req.Code,
		Language:  req.Language,
	}

	if result := s.db.Create(sk); result.Error != nil {
		// Check if it's a unique constraint violation
		if result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"idx_solution_keys_unique\" (SQLSTATE 23505)" {
			return nil, fmt.Errorf("solution key already exists for this problem and language")
		}
		return nil, result.Error
	}

	return sk, nil
}

// UpdateSolutionKey updates an existing solution key with validation
func (s *SolutionKeyService) UpdateSolutionKey(id string, req UpdateSolutionKeyRequest) (*models.SolutionKey, error) {
	// Validate request
	if err := s.ValidateSolutionKeyRequest(req); err != nil {
		return nil, err
	}

	// Get existing solution key
	sk := &models.SolutionKey{}
	if result := s.db.First(sk, "id = ?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("solution key not found")
		}
		return nil, result.Error
	}

	// Update fields
	sk.Code = req.Code
	if req.Language != "" {
		sk.Language = req.Language
	}

	if result := s.db.Save(sk); result.Error != nil {
		return nil, result.Error
	}

	return sk, nil
}

// GetSolutionKeyByProblemID retrieves a solution key by problem ID
func (s *SolutionKeyService) GetSolutionKeyByProblemID(problemID string) (*models.SolutionKey, error) {
	sk := &models.SolutionKey{}
	if result := s.db.First(sk, "problem_id = ?", problemID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("solution key not found")
		}
		return nil, result.Error
	}
	return sk, nil
}

// DeleteSolutionKey deletes a solution key by ID
func (s *SolutionKeyService) DeleteSolutionKey(id string) error {
	if result := s.db.Delete(&models.SolutionKey{}, "id = ?", id); result.Error != nil {
		return result.Error
	}
	return nil
}

// ValidateSolutionKeyRequest validates solution key request fields
func (s *SolutionKeyService) ValidateSolutionKeyRequest(req interface{}) error {
	switch r := req.(type) {
	case CreateSolutionKeyRequest:
		if r.ProblemID == "" {
			return fmt.Errorf("problemId is required")
		}
		if r.Code == "" {
			return fmt.Errorf("code is required")
		}
		if r.Language != "javascript" && r.Language != "sql" {
			return fmt.Errorf("language must be 'javascript' or 'sql'")
		}
	case UpdateSolutionKeyRequest:
		if r.Code == "" {
			return fmt.Errorf("code is required")
		}
		if r.Language != "" && r.Language != "javascript" && r.Language != "sql" {
			return fmt.Errorf("language must be 'javascript' or 'sql'")
		}
	default:
		return fmt.Errorf("invalid request type")
	}
	return nil
}
