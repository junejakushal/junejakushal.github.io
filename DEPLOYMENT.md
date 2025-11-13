# Quick Deployment Guide

This guide will help you deploy the Audio Transcription Service to Fly.io in minutes.

## Prerequisites

1. **AssemblyAI API Key**: Sign up at [AssemblyAI](https://www.assemblyai.com/) and get your API key
2. **Fly.io Account**: Create a free account at [Fly.io](https://fly.io/)

## Step-by-Step Deployment

### 1. Install Fly CLI

**macOS/Linux:**
```bash
curl -L https://fly.io/install.sh | sh
```

**Windows (PowerShell):**
```powershell
pwsh -Command "iwr https://fly.io/install.ps1 -useb | iex"
```

### 2. Login to Fly.io

```bash
fly auth login
```

This will open a browser window for authentication.

### 3. Clone and Navigate to Repository

```bash
git clone https://github.com/junejakushal/junejakushal.github.io.git
cd junejakushal.github.io
```

### 4. Create Your Fly App

```bash
fly launch
```

You'll be asked several questions:
- **App name**: Choose a unique name or press Enter for auto-generated
- **Organization**: Select your organization
- **Region**: Choose closest to your users (e.g., `iad` for US East)
- **Deploy now?**: Type `n` (we'll customize first)

### 5. Customize fly.toml (Optional)

Edit `fly.toml` to change:
- `app` name to your chosen name
- `primary_region` to your preferred region
- `memory` size if needed (default is 256MB)

### 6. Deploy Your App

```bash
fly deploy
```

This will:
- Build the Docker image
- Push to Fly.io registry
- Deploy your application

### 7. Open Your App

```bash
fly open
```

Your application is now live! 🎉

## Usage

1. Navigate to your deployed URL
2. Enter your AssemblyAI API key
3. Upload audio files (drag & drop or click to browse)
4. Click "Transcribe Audio Files"
5. Wait for processing (typically 1-2 minutes per minute of audio)
6. Download your combined transcript

## Managing Your App

### View Logs
```bash
fly logs
```

### Check Status
```bash
fly status
```

### Scale Resources
```bash
fly scale memory 512  # Increase to 512MB
fly scale count 2     # Run 2 instances
```

### Update After Code Changes
```bash
git pull
fly deploy
```

### Delete App
```bash
fly apps destroy your-app-name
```

## Troubleshooting

### App Won't Start
- Check logs: `fly logs`
- Verify your Dockerfile builds locally: `docker build -t test .`

### Out of Memory
- Increase memory: `fly scale memory 512`
- Check large file uploads aren't overwhelming the server

### Slow Transcription
- This depends on AssemblyAI's API speed
- Consider upgrading your AssemblyAI plan for faster processing

### Can't Access App
- Verify app is running: `fly status`
- Check if region is accessible from your location
- Try opening directly: `https://your-app-name.fly.dev`

## Cost Estimate

**Fly.io** (Free Tier Includes):
- 3 shared-cpu-1x VMs
- 3GB persistent storage
- 160GB bandwidth

With the default configuration (auto-stop enabled), the app will:
- Start when someone visits
- Stop after idle time
- Cost $0 if you stay within free tier limits

**AssemblyAI**:
- Free tier: 5 hours of transcription per month
- Pay-as-you-go: $0.00025 per second (~$0.015 per minute)

## Environment Variables

You can set secrets on Fly.io if needed:

```bash
fly secrets set MY_SECRET=value
```

However, this app doesn't require server-side secrets - users provide their own API keys.

## Security Best Practices

1. **Never commit API keys** to your repository
2. **Use HTTPS** (automatically enforced on Fly.io)
3. **Keep dependencies updated**: Run `go get -u` regularly
4. **Monitor logs** for unusual activity
5. **Rate limiting**: Consider adding if public

## Support

- **Fly.io Docs**: https://fly.io/docs
- **AssemblyAI Docs**: https://www.assemblyai.com/docs
- **Go Documentation**: https://golang.org/doc/

## Next Steps

Consider adding:
- User authentication
- Database for tracking transcriptions
- Email notifications when complete
- Support for more audio formats
- Custom vocabulary for better accuracy
- Multiple language support

Happy transcribing! 🎙️
