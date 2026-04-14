package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileStorageService handles storing and retrieving code files from the file system.
type FileStorageService struct {
	basePath string
}

// NewFileStorageService creates a new FileStorageService with the given base path.
// If the base path doesn't exist, it will be created.
func NewFileStorageService(basePath string) (*FileStorageService, error) {
	// Create base directory if it doesn't exist
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	return &FileStorageService{
		basePath: basePath,
	}, nil
}

// SaveCode saves code to a file and returns the file path relative to basePath.
// The file path is in format: {submissionID}.{language}
func (s *FileStorageService) SaveCode(submissionID string, language string, code string) (string, error) {
	if submissionID == "" {
		return "", fmt.Errorf("submission ID cannot be empty")
	}
	if language == "" {
		return "", fmt.Errorf("language cannot be empty")
	}

	// Create file path
	fileName := fmt.Sprintf("%s.%s", submissionID, language)
	filePath := filepath.Join(s.basePath, fileName)

	// Write code to file
	if err := os.WriteFile(filePath, []byte(code), 0644); err != nil {
		return "", fmt.Errorf("failed to write code file: %w", err)
	}

	return fileName, nil
}

// ReadCode reads code from a file given the file path (relative to basePath).
func (s *FileStorageService) ReadCode(filePath string) (string, error) {
	if filePath == "" {
		return "", fmt.Errorf("file path cannot be empty")
	}

	// Prevent directory traversal attacks
	absPath := filepath.Join(s.basePath, filePath)
	absBasePath, err := filepath.Abs(s.basePath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute base path: %w", err)
	}

	absFilePath, err := filepath.Abs(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute file path: %w", err)
	}

	// Ensure the file is within the base path
	if !filepath.HasPrefix(absFilePath, absBasePath) {
		return "", fmt.Errorf("file path is outside base directory")
	}

	// Read code from file
	code, err := os.ReadFile(absFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read code file: %w", err)
	}

	return string(code), nil
}

// DeleteCode deletes a code file given the file path (relative to basePath).
func (s *FileStorageService) DeleteCode(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("file path cannot be empty")
	}

	// Prevent directory traversal attacks
	absPath := filepath.Join(s.basePath, filePath)
	absBasePath, err := filepath.Abs(s.basePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute base path: %w", err)
	}

	absFilePath, err := filepath.Abs(absPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute file path: %w", err)
	}

	// Ensure the file is within the base path
	if !filepath.HasPrefix(absFilePath, absBasePath) {
		return fmt.Errorf("file path is outside base directory")
	}

	// Delete file
	if err := os.Remove(absFilePath); err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist, consider it success
		}
		return fmt.Errorf("failed to delete code file: %w", err)
	}

	return nil
}

// Exists checks if a code file exists.
func (s *FileStorageService) Exists(filePath string) (bool, error) {
	if filePath == "" {
		return false, fmt.Errorf("file path cannot be empty")
	}

	// Prevent directory traversal attacks
	absPath := filepath.Join(s.basePath, filePath)
	absBasePath, err := filepath.Abs(s.basePath)
	if err != nil {
		return false, fmt.Errorf("failed to get absolute base path: %w", err)
	}

	absFilePath, err := filepath.Abs(absPath)
	if err != nil {
		return false, fmt.Errorf("failed to get absolute file path: %w", err)
	}

	// Ensure the file is within the base path
	if !filepath.HasPrefix(absFilePath, absBasePath) {
		return false, fmt.Errorf("file path is outside base directory")
	}

	_, err = os.Stat(absFilePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("failed to check file existence: %w", err)
}
