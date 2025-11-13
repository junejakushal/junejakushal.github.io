# Audio Transcription Service

A Go web application that transcribes multiple audio files in parallel using AssemblyAI and provides a combined transcript for download.

## Features

- 🎙️ **Multiple File Upload**: Upload multiple audio files simultaneously
- ⚡ **Parallel Processing**: Transcribes all files in parallel for faster processing
- 📝 **Combined Transcript**: Automatically combines all transcriptions into a single downloadable text file
- 🎨 **Modern UI**: Beautiful, responsive web interface
- 🚀 **Cloud Ready**: Configured for deployment on Fly.io

## Prerequisites

- Go 1.24 or higher
- AssemblyAI API key ([Get one here](https://www.assemblyai.com/))
- (For deployment) Fly.io account

## Local Development

### 1. Clone the repository

```bash
git clone https://github.com/junejakushal/junejakushal.github.io.git
cd junejakushal.github.io
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Run the application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### 4. Usage

1. Open your browser and navigate to `http://localhost:8080`
2. Enter your AssemblyAI API key
3. Upload one or more audio files (MP3, WAV, M4A, FLAC, etc.)
4. Click "Transcribe Audio Files"
5. Wait for processing to complete
6. Download your combined transcript

## Deployment to Fly.io

### 1. Install Fly CLI

```bash
curl -L https://fly.io/install.sh | sh
```

### 2. Login to Fly.io

```bash
fly auth login
```

### 3. Create a new app (first time only)

```bash
fly launch
```

Follow the prompts to configure your app. The `fly.toml` file is already configured.

### 4. Deploy

```bash
fly deploy
```

### 5. Open your app

```bash
fly open
```

## Configuration

### Environment Variables

- `PORT`: Server port (default: 8080)

### fly.toml Configuration

The application is configured with:
- Automatic HTTPS
- Auto-start/stop machines to save costs
- Health checks
- 256MB memory allocation
- Shared CPU

You can modify these settings in `fly.toml` as needed.

## API Endpoints

- `GET /`: Main upload page (HTML interface)
- `POST /upload`: Handle file uploads and transcription
  - Form fields:
    - `apiKey`: AssemblyAI API key
    - `files`: Multiple audio files
- `GET /download/{filename}`: Download transcript file

## Supported Audio Formats

The application supports all audio formats supported by AssemblyAI, including:
- MP3
- WAV
- M4A
- FLAC
- OGG
- And more

## Technical Details

### Architecture

1. **File Upload**: Users upload multiple audio files via the web interface
2. **Storage**: Files are temporarily stored on the server
3. **Parallel Processing**: Each file is:
   - Uploaded to AssemblyAI
   - Submitted for transcription
   - Polled until completion
4. **Combination**: All transcripts are combined with metadata
5. **Download**: Combined transcript is made available for download
6. **Cleanup**: Files are automatically deleted after 10 minutes

### Performance

- Parallel processing reduces total transcription time
- Maximum file size: 100MB per file
- Automatic timeout handling
- Connection pooling for API requests

## Security Notes

- API keys are never stored on the server
- Temporary files are automatically cleaned up
- HTTPS enforced in production
- No persistent storage of user data

## Troubleshooting

### "File too large" error
- Reduce file size or split into smaller files
- Maximum size per file is 100MB

### "Transcription failed" error
- Check your AssemblyAI API key is valid
- Ensure you have sufficient credits in your AssemblyAI account
- Verify the audio file format is supported

### Deployment issues
- Ensure Fly CLI is properly installed
- Check you're logged in: `fly auth whoami`
- Verify your app name is unique: modify `app` in `fly.toml`

## License

This project is open source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues and questions:
- Open an issue on GitHub
- Check AssemblyAI documentation: https://www.assemblyai.com/docs
- Fly.io documentation: https://fly.io/docs
