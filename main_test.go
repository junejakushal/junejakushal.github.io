package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandleHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handleHome(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	if !strings.Contains(string(body), "Audio Transcription Service") {
		t.Errorf("Expected page to contain title")
	}

	if !strings.Contains(string(body), "AssemblyAI") {
		t.Errorf("Expected page to mention AssemblyAI")
	}
}

func TestHandleUploadNoFiles(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("apiKey", "test-key")
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	handleUpload(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", resp.StatusCode)
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if !strings.Contains(result["error"], "No files uploaded") {
		t.Errorf("Expected 'No files uploaded' error, got %v", result["error"])
	}
}

func TestHandleUploadNoAPIKey(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// No API key provided
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	handleUpload(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %v", resp.StatusCode)
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if !strings.Contains(result["error"], "API key") {
		t.Errorf("Expected 'API key' error, got %v", result["error"])
	}
}

func TestHandleDownloadInvalidFilename(t *testing.T) {
	// Test path traversal attempt
	req := httptest.NewRequest(http.MethodGet, "/download/../../../etc/passwd", nil)
	w := httptest.NewRecorder()

	handleDownload(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest for path traversal, got %v", resp.StatusCode)
	}
}

func TestHandleDownloadInvalidFormat(t *testing.T) {
	// Test invalid filename format
	req := httptest.NewRequest(http.MethodGet, "/download/malicious.txt", nil)
	w := httptest.NewRecorder()

	handleDownload(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest for invalid format, got %v", resp.StatusCode)
	}
}

func TestHandleDownloadNonExistentFile(t *testing.T) {
	// Create output directory for test
	os.MkdirAll(outputDir, 0755)
	defer os.RemoveAll(outputDir)

	req := httptest.NewRequest(http.MethodGet, "/download/transcript_session_999999999.txt", nil)
	w := httptest.NewRecorder()

	handleDownload(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status NotFound, got %v", resp.StatusCode)
	}
}

func TestHandleDownloadSuccess(t *testing.T) {
	// Create output directory and test file
	os.MkdirAll(outputDir, 0755)
	defer os.RemoveAll(outputDir)

	testFile := filepath.Join(outputDir, "transcript_session_12345.txt")
	testContent := "Test transcription content"
	os.WriteFile(testFile, []byte(testContent), 0644)

	req := httptest.NewRequest(http.MethodGet, "/download/transcript_session_12345.txt", nil)
	w := httptest.NewRecorder()

	handleDownload(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	if string(body) != testContent {
		t.Errorf("Expected content '%s', got '%s'", testContent, string(body))
	}

	contentDisposition := resp.Header.Get("Content-Disposition")
	if !strings.Contains(contentDisposition, "transcript_session_12345.txt") {
		t.Errorf("Expected Content-Disposition to contain filename, got '%s'", contentDisposition)
	}
}

func TestCombineTranscriptions(t *testing.T) {
	transcriptions := []string{
		"This is the first transcript.",
		"This is the second transcript.",
	}
	filePaths := []string{
		"/path/to/audio1.mp3",
		"/path/to/audio2.mp3",
	}

	result := combineTranscriptions(transcriptions, filePaths)

	if !strings.Contains(result, "Combined Audio Transcription") {
		t.Errorf("Expected header in combined transcription")
	}

	if !strings.Contains(result, "audio1.mp3") {
		t.Errorf("Expected first filename in combined transcription")
	}

	if !strings.Contains(result, "audio2.mp3") {
		t.Errorf("Expected second filename in combined transcription")
	}

	if !strings.Contains(result, "This is the first transcript.") {
		t.Errorf("Expected first transcript content")
	}

	if !strings.Contains(result, "This is the second transcript.") {
		t.Errorf("Expected second transcript content")
	}
}

func TestRespondJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{
		"message": "test message",
		"status":  "success",
	}

	respondJSON(w, http.StatusOK, data)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Expected Content-Type to be application/json, got '%s'", contentType)
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if result["message"] != "test message" {
		t.Errorf("Expected message 'test message', got '%s'", result["message"])
	}
}
