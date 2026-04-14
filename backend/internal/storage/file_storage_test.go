package storage

import (
	"os"
	"path/filepath"
	"testing"
)

// TestFileStorageServiceSaveAndRead verifies that code can be saved and read correctly
func TestFileStorageServiceSaveAndRead(t *testing.T) {
	// Create temporary directory for testing
	tmpDir := t.TempDir()
	storage, err := NewFileStorageService(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage service: %v", err)
	}

	tests := []struct {
		name         string
		submissionID string
		language     string
		code         string
	}{
		{
			name:         "simple javascript code",
			submissionID: "sub-1",
			language:     "javascript",
			code:         "function add(a, b) { return a + b; }",
		},
		{
			name:         "sql query",
			submissionID: "sub-2",
			language:     "sql",
			code:         "SELECT * FROM users WHERE id = 1",
		},
		{
			name:         "code with newlines",
			submissionID: "sub-3",
			language:     "javascript",
			code:         "function test() {\n  console.log('hello');\n  return 42;\n}",
		},
		{
			name:         "empty code",
			submissionID: "sub-4",
			language:     "javascript",
			code:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save code
			filePath, err := storage.SaveCode(tt.submissionID, tt.language, tt.code)
			if err != nil {
				t.Fatalf("failed to save code: %v", err)
			}

			// Verify file path format
			expectedPath := tt.submissionID + "." + tt.language
			if filePath != expectedPath {
				t.Errorf("expected file path %q, got %q", expectedPath, filePath)
			}

			// Read code back
			readCode, err := storage.ReadCode(filePath)
			if err != nil {
				t.Fatalf("failed to read code: %v", err)
			}

			// Verify code matches
			if readCode != tt.code {
				t.Errorf("code mismatch:\n  expected: %q\n  got: %q", tt.code, readCode)
			}
		})
	}
}

// TestFileStorageServiceDelete verifies that code files can be deleted
func TestFileStorageServiceDelete(t *testing.T) {
	tmpDir := t.TempDir()
	storage, err := NewFileStorageService(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage service: %v", err)
	}

	// Save code
	filePath, err := storage.SaveCode("sub-1", "javascript", "console.log('test');")
	if err != nil {
		t.Fatalf("failed to save code: %v", err)
	}

	// Verify file exists
	exists, err := storage.Exists(filePath)
	if err != nil {
		t.Fatalf("failed to check existence: %v", err)
	}
	if !exists {
		t.Fatal("file should exist after save")
	}

	// Delete file
	err = storage.DeleteCode(filePath)
	if err != nil {
		t.Fatalf("failed to delete code: %v", err)
	}

	// Verify file no longer exists
	exists, err = storage.Exists(filePath)
	if err != nil {
		t.Fatalf("failed to check existence: %v", err)
	}
	if exists {
		t.Fatal("file should not exist after delete")
	}

	// Deleting non-existent file should not error
	err = storage.DeleteCode(filePath)
	if err != nil {
		t.Fatalf("deleting non-existent file should not error: %v", err)
	}
}

// TestFileStorageServiceExists verifies that file existence check works
func TestFileStorageServiceExists(t *testing.T) {
	tmpDir := t.TempDir()
	storage, err := NewFileStorageService(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage service: %v", err)
	}

	filePath := "sub-1.javascript"

	// File should not exist initially
	exists, err := storage.Exists(filePath)
	if err != nil {
		t.Fatalf("failed to check existence: %v", err)
	}
	if exists {
		t.Fatal("file should not exist initially")
	}

	// Save code
	_, err = storage.SaveCode("sub-1", "javascript", "console.log('test');")
	if err != nil {
		t.Fatalf("failed to save code: %v", err)
	}

	// File should exist now
	exists, err = storage.Exists(filePath)
	if err != nil {
		t.Fatalf("failed to check existence: %v", err)
	}
	if !exists {
		t.Fatal("file should exist after save")
	}
}

// TestFileStorageServiceDirectoryTraversal verifies that directory traversal attacks are prevented
func TestFileStorageServiceDirectoryTraversal(t *testing.T) {
	tmpDir := t.TempDir()
	storage, err := NewFileStorageService(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage service: %v", err)
	}

	tests := []struct {
		name     string
		filePath string
	}{
		{
			name:     "parent directory traversal",
			filePath: "../../../etc/passwd",
		},
		{
			name:     "backslash traversal",
			filePath: "..\\..\\..\\windows\\system32",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ReadCode should fail
			_, err := storage.ReadCode(tt.filePath)
			if err == nil {
				t.Fatal("expected error for directory traversal attempt")
			}

			// DeleteCode should fail
			err = storage.DeleteCode(tt.filePath)
			if err == nil {
				t.Fatal("expected error for directory traversal attempt")
			}

			// Exists should fail
			_, err = storage.Exists(tt.filePath)
			if err == nil {
				t.Fatal("expected error for directory traversal attempt")
			}
		})
	}
}

// TestFileStorageServiceConcurrentAccess verifies that concurrent file operations work correctly
func TestFileStorageServiceConcurrentAccess(t *testing.T) {
	tmpDir := t.TempDir()
	storage, err := NewFileStorageService(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage service: %v", err)
	}

	numGoroutines := 10
	results := make(chan bool, numGoroutines)
	errors := make(chan string, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			submissionID := "sub-" + string(rune('0'+id))
			code := "function test() { return " + string(rune('0'+id)) + "; }"

			// Save code
			filePath, err := storage.SaveCode(submissionID, "javascript", code)
			if err != nil {
				errors <- err.Error()
				return
			}

			// Read code back
			readCode, err := storage.ReadCode(filePath)
			if err != nil {
				errors <- err.Error()
				return
			}

			// Verify code matches
			if readCode != code {
				errors <- "code mismatch"
				return
			}

			results <- true
		}(i)
	}

	// Collect results
	successCount := 0
	for i := 0; i < numGoroutines; i++ {
		select {
		case errMsg := <-errors:
			t.Errorf("concurrent operation failed: %s", errMsg)
		case <-results:
			successCount++
		}
	}

	if successCount != numGoroutines {
		t.Errorf("expected all %d concurrent operations to succeed, got %d", numGoroutines, successCount)
	}
}

// TestFileStorageServiceEmptyInputs verifies that empty inputs are rejected
func TestFileStorageServiceEmptyInputs(t *testing.T) {
	tmpDir := t.TempDir()
	storage, err := NewFileStorageService(tmpDir)
	if err != nil {
		t.Fatalf("failed to create storage service: %v", err)
	}

	tests := []struct {
		name         string
		submissionID string
		language     string
		code         string
		shouldFail   bool
	}{
		{
			name:         "empty submission ID",
			submissionID: "",
			language:     "javascript",
			code:         "console.log('test');",
			shouldFail:   true,
		},
		{
			name:         "empty language",
			submissionID: "sub-1",
			language:     "",
			code:         "console.log('test');",
			shouldFail:   true,
		},
		{
			name:         "valid with empty code",
			submissionID: "sub-1",
			language:     "javascript",
			code:         "",
			shouldFail:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := storage.SaveCode(tt.submissionID, tt.language, tt.code)
			if tt.shouldFail && err == nil {
				t.Fatal("expected error for invalid input")
			}
			if !tt.shouldFail && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

// TestFileStorageServiceDirectoryCreation verifies that base directory is created if it doesn't exist
func TestFileStorageServiceDirectoryCreation(t *testing.T) {
	tmpDir := t.TempDir()
	storagePath := filepath.Join(tmpDir, "submissions", "nested", "path")

	// Directory should not exist yet
	if _, err := os.Stat(storagePath); err == nil {
		t.Fatal("directory should not exist before creating storage service")
	}

	// Create storage service
	storage, err := NewFileStorageService(storagePath)
	if err != nil {
		t.Fatalf("failed to create storage service: %v", err)
	}

	// Directory should exist now
	if _, err := os.Stat(storagePath); err != nil {
		t.Fatalf("directory should exist after creating storage service: %v", err)
	}

	// Should be able to save files
	_, err = storage.SaveCode("sub-1", "javascript", "console.log('test');")
	if err != nil {
		t.Fatalf("failed to save code: %v", err)
	}
}
