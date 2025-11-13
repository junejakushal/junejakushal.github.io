package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	uploadDir     = "./uploads"
	outputDir     = "./outputs"
	maxUploadSize = 100 << 20 // 100 MB per file
	assemblyAIURL = "https://api.assemblyai.com"
)

type TranscriptionResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Text   string `json:"text"`
}

type UploadResponse struct {
	UploadURL string `json:"upload_url"`
}

type TranscriptRequest struct {
	AudioURL string `json:"audio_url"`
}

func main() {
	// Create necessary directories
	os.MkdirAll(uploadDir, 0755)
	os.MkdirAll(outputDir, 0755)

	// Routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/download/", handleDownload)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Audio Transcription Service</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            max-width: 600px;
            width: 100%;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 28px;
            text-align: center;
        }
        .subtitle {
            color: #666;
            text-align: center;
            margin-bottom: 30px;
            font-size: 14px;
        }
        .upload-area {
            border: 3px dashed #667eea;
            border-radius: 15px;
            padding: 40px 20px;
            text-align: center;
            background: #f8f9ff;
            transition: all 0.3s ease;
            cursor: pointer;
        }
        .upload-area:hover {
            border-color: #764ba2;
            background: #f0f1ff;
        }
        .upload-area.dragover {
            border-color: #764ba2;
            background: #e8e9ff;
            transform: scale(1.02);
        }
        .upload-icon {
            font-size: 48px;
            color: #667eea;
            margin-bottom: 15px;
        }
        input[type="file"] {
            display: none;
        }
        .file-label {
            color: #667eea;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
        }
        .file-info {
            margin-top: 20px;
            color: #666;
            font-size: 14px;
        }
        .file-list {
            margin-top: 15px;
            text-align: left;
            max-height: 200px;
            overflow-y: auto;
        }
        .file-item {
            background: #f5f5f5;
            padding: 10px;
            margin: 5px 0;
            border-radius: 8px;
            font-size: 14px;
            color: #333;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .remove-file {
            color: #dc3545;
            cursor: pointer;
            font-weight: bold;
            font-size: 18px;
            line-height: 1;
        }
        .remove-file:hover {
            color: #c82333;
        }
        button {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            padding: 15px 40px;
            border-radius: 25px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            width: 100%;
            margin-top: 20px;
            transition: all 0.3s ease;
            box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
        }
        button:hover:not(:disabled) {
            transform: translateY(-2px);
            box-shadow: 0 6px 20px rgba(102, 126, 234, 0.6);
        }
        button:disabled {
            background: #ccc;
            cursor: not-allowed;
            box-shadow: none;
        }
        .status {
            margin-top: 20px;
            padding: 15px;
            border-radius: 10px;
            text-align: center;
            font-weight: 500;
        }
        .status.success {
            background: #d4edda;
            color: #155724;
            border: 1px solid #c3e6cb;
        }
        .status.error {
            background: #f8d7da;
            color: #721c24;
            border: 1px solid #f5c6cb;
        }
        .status.processing {
            background: #d1ecf1;
            color: #0c5460;
            border: 1px solid #bee5eb;
        }
        .spinner {
            border: 3px solid #f3f3f3;
            border-top: 3px solid #667eea;
            border-radius: 50%;
            width: 30px;
            height: 30px;
            animation: spin 1s linear infinite;
            margin: 10px auto;
        }
        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
        .download-link {
            display: inline-block;
            margin-top: 10px;
            color: #667eea;
            text-decoration: none;
            font-weight: 600;
            padding: 10px 20px;
            border: 2px solid #667eea;
            border-radius: 20px;
            transition: all 0.3s ease;
        }
        .download-link:hover {
            background: #667eea;
            color: white;
        }
        .api-key-section {
            margin-bottom: 20px;
        }
        .api-key-section label {
            display: block;
            color: #333;
            font-weight: 600;
            margin-bottom: 8px;
        }
        .api-key-section input {
            width: 100%;
            padding: 12px;
            border: 2px solid #e0e0e0;
            border-radius: 10px;
            font-size: 14px;
            transition: border-color 0.3s ease;
        }
        .api-key-section input:focus {
            outline: none;
            border-color: #667eea;
        }
        .api-key-section small {
            display: block;
            color: #666;
            margin-top: 5px;
            font-size: 12px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎙️ Audio Transcription Service</h1>
        <p class="subtitle">Upload multiple audio files and get a combined transcript powered by AssemblyAI</p>
        
        <form id="uploadForm" enctype="multipart/form-data">
            <div class="api-key-section">
                <label for="apiKey">AssemblyAI API Key:</label>
                <input type="password" id="apiKey" name="apiKey" placeholder="Enter your AssemblyAI API key" required>
                <small>Get your API key from <a href="https://www.assemblyai.com/" target="_blank">AssemblyAI</a></small>
            </div>
            
            <div class="upload-area" id="uploadArea">
                <div class="upload-icon">📁</div>
                <div class="file-label">
                    Click to browse or drag and drop audio files here
                    <input type="file" id="fileInput" name="files" multiple accept="audio/*">
                </div>
                <p class="file-info">Supported formats: MP3, WAV, M4A, FLAC, etc. (Max 100MB per file)</p>
            </div>
            
            <div class="file-list" id="fileList"></div>
            
            <button type="submit" id="submitBtn" disabled>Transcribe Audio Files</button>
        </form>
        
        <div id="status"></div>
    </div>

    <script>
        let selectedFiles = [];
        const uploadArea = document.getElementById('uploadArea');
        const fileInput = document.getElementById('fileInput');
        const fileList = document.getElementById('fileList');
        const submitBtn = document.getElementById('submitBtn');
        const status = document.getElementById('status');
        const form = document.getElementById('uploadForm');
        const apiKeyInput = document.getElementById('apiKey');

        uploadArea.addEventListener('click', () => fileInput.click());

        uploadArea.addEventListener('dragover', (e) => {
            e.preventDefault();
            uploadArea.classList.add('dragover');
        });

        uploadArea.addEventListener('dragleave', () => {
            uploadArea.classList.remove('dragover');
        });

        uploadArea.addEventListener('drop', (e) => {
            e.preventDefault();
            uploadArea.classList.remove('dragover');
            handleFiles(e.dataTransfer.files);
        });

        fileInput.addEventListener('change', (e) => {
            handleFiles(e.target.files);
        });

        apiKeyInput.addEventListener('input', updateSubmitButton);

        function handleFiles(files) {
            selectedFiles = Array.from(files);
            displayFiles();
            updateSubmitButton();
        }

        function displayFiles() {
            fileList.innerHTML = '';
            if (selectedFiles.length === 0) {
                fileList.style.display = 'none';
                return;
            }
            
            fileList.style.display = 'block';
            selectedFiles.forEach((file, index) => {
                const fileItem = document.createElement('div');
                fileItem.className = 'file-item';
                const fileName = file.name;
                const fileSize = (file.size / 1024 / 1024).toFixed(2);
                fileItem.innerHTML = '<span>' + fileName + ' (' + fileSize + ' MB)</span>' +
                    '<span class="remove-file" onclick="removeFile(' + index + ')">×</span>';
                fileList.appendChild(fileItem);
            });
        }

        function removeFile(index) {
            selectedFiles.splice(index, 1);
            displayFiles();
            updateSubmitButton();
        }

        function updateSubmitButton() {
            submitBtn.disabled = selectedFiles.length === 0 || !apiKeyInput.value.trim();
        }

        form.addEventListener('submit', async (e) => {
            e.preventDefault();
            
            if (selectedFiles.length === 0) {
                showStatus('Please select at least one audio file', 'error');
                return;
            }

            if (!apiKeyInput.value.trim()) {
                showStatus('Please enter your AssemblyAI API key', 'error');
                return;
            }

            const formData = new FormData();
            formData.append('apiKey', apiKeyInput.value.trim());
            selectedFiles.forEach(file => {
                formData.append('files', file);
            });

            submitBtn.disabled = true;
            showStatus('Uploading and transcribing files... This may take a few minutes.', 'processing');

            try {
                const response = await fetch('/upload', {
                    method: 'POST',
                    body: formData
                });

                const result = await response.json();

                if (response.ok) {
                    showStatus(
                        'Success! Transcription complete.<br><a href="' + result.downloadURL + '" class="download-link" download>Download Transcript</a>',
                        'success'
                    );
                    selectedFiles = [];
                    displayFiles();
                    fileInput.value = '';
                } else {
                    showStatus('Error: ' + result.error, 'error');
                }
            } catch (error) {
                showStatus('Error: ' + error.message, 'error');
            } finally {
                submitBtn.disabled = false;
                updateSubmitButton();
            }
        });

        function showStatus(message, type) {
            status.className = 'status ' + type;
            status.innerHTML = message;
            if (type === 'processing') {
                status.innerHTML += '<div class="spinner"></div>';
            }
        }
    </script>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "File too large or invalid form data"})
		return
	}

	// Get API key
	apiKey := r.FormValue("apiKey")
	if apiKey == "" {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "AssemblyAI API key is required"})
		return
	}

	// Get uploaded files
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		respondJSON(w, http.StatusBadRequest, map[string]string{"error": "No files uploaded"})
		return
	}

	log.Printf("Processing %d files", len(files))

	// Save uploaded files and collect paths
	var filePaths []string
	sessionID := fmt.Sprintf("session_%d", time.Now().Unix())
	sessionDir := filepath.Join(uploadDir, sessionID)

	if err := os.MkdirAll(sessionDir, 0755); err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create upload directory"})
		return
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to read uploaded file"})
			return
		}
		defer file.Close()

		filePath := filepath.Join(sessionDir, fileHeader.Filename)
		dst, err := os.Create(filePath)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to save uploaded file"})
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to save uploaded file"})
			return
		}

		filePaths = append(filePaths, filePath)
		log.Printf("Saved file: %s", filePath)
	}

	// Transcribe all files in parallel
	transcriptions, err := transcribeFilesParallel(filePaths, apiKey)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Transcription failed: %v", err)})
		return
	}

	// Combine transcriptions
	combinedText := combineTranscriptions(transcriptions, filePaths)

	// Save to output file
	outputFilename := fmt.Sprintf("transcript_%s.txt", sessionID)
	outputPath := filepath.Join(outputDir, outputFilename)

	if err := os.WriteFile(outputPath, []byte(combinedText), 0644); err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to save transcript"})
		return
	}

	log.Printf("Transcript saved: %s", outputPath)

	// Clean up uploaded files
	go func() {
		time.Sleep(10 * time.Minute)
		os.RemoveAll(sessionDir)
		os.Remove(outputPath)
	}()

	respondJSON(w, http.StatusOK, map[string]string{
		"message":     "Transcription completed successfully",
		"downloadURL": "/download/" + outputFilename,
	})
}

func transcribeFilesParallel(filePaths []string, apiKey string) ([]string, error) {
	var wg sync.WaitGroup
	results := make([]string, len(filePaths))
	errors := make([]error, len(filePaths))

	for i, filePath := range filePaths {
		wg.Add(1)
		go func(index int, path string) {
			defer wg.Done()
			transcript, err := transcribeFile(path, apiKey)
			if err != nil {
				errors[index] = err
				log.Printf("Error transcribing %s: %v", path, err)
			} else {
				results[index] = transcript
				log.Printf("Successfully transcribed: %s", path)
			}
		}(i, filePath)
	}

	wg.Wait()

	// Check for errors
	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

func transcribeFile(filePath string, apiKey string) (string, error) {
	// Step 1: Upload file to AssemblyAI
	uploadURL, err := uploadToAssemblyAI(filePath, apiKey)
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}

	// Step 2: Create transcription job
	transcriptID, err := createTranscription(uploadURL, apiKey)
	if err != nil {
		return "", fmt.Errorf("transcription creation failed: %w", err)
	}

	// Step 3: Poll for completion
	transcript, err := pollTranscription(transcriptID, apiKey)
	if err != nil {
		return "", fmt.Errorf("transcription polling failed: %w", err)
	}

	return transcript, nil
}

func uploadToAssemblyAI(filePath string, apiKey string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	req, err := http.NewRequest("POST", assemblyAIURL+"/v2/upload", file)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{Timeout: 5 * time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	var uploadResp UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return "", err
	}

	return uploadResp.UploadURL, nil
}

func createTranscription(audioURL string, apiKey string) (string, error) {
	reqBody := TranscriptRequest{AudioURL: audioURL}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", assemblyAIURL+"/v2/transcript", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("transcription creation failed with status %d: %s", resp.StatusCode, string(body))
	}

	var transcriptResp TranscriptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&transcriptResp); err != nil {
		return "", err
	}

	return transcriptResp.ID, nil
}

func pollTranscription(transcriptID string, apiKey string) (string, error) {
	url := fmt.Sprintf("%s/v2/transcript/%s", assemblyAIURL, transcriptID)
	client := &http.Client{Timeout: 30 * time.Second}

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", err
		}

		req.Header.Set("Authorization", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		var transcriptResp TranscriptionResponse
		if err := json.NewDecoder(resp.Body).Decode(&transcriptResp); err != nil {
			resp.Body.Close()
			return "", err
		}
		resp.Body.Close()

		switch transcriptResp.Status {
		case "completed":
			return transcriptResp.Text, nil
		case "error":
			return "", fmt.Errorf("transcription failed")
		case "queued", "processing":
			time.Sleep(3 * time.Second)
		default:
			return "", fmt.Errorf("unknown status: %s", transcriptResp.Status)
		}
	}
}

func combineTranscriptions(transcriptions []string, filePaths []string) string {
	var sb strings.Builder
	sb.WriteString("Combined Audio Transcription\n")
	sb.WriteString("============================\n")
	sb.WriteString(fmt.Sprintf("Generated: %s\n", time.Now().Format(time.RFC1123)))
	sb.WriteString(fmt.Sprintf("Total Files: %d\n\n", len(transcriptions)))

	for i, transcript := range transcriptions {
		filename := filepath.Base(filePaths[i])
		sb.WriteString(fmt.Sprintf("\n--- File %d: %s ---\n\n", i+1, filename))
		sb.WriteString(transcript)
		sb.WriteString("\n\n")
	}

	return sb.String()
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	filename := strings.TrimPrefix(r.URL.Path, "/download/")

	// Sanitize filename to prevent path traversal attacks
	filename = filepath.Base(filename)

	// Validate filename format (should be transcript_session_<timestamp>.txt)
	if !strings.HasPrefix(filename, "transcript_session_") || !strings.HasSuffix(filename, ".txt") {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(outputDir, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, filePath)
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
