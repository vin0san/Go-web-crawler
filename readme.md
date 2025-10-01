# Go Web Crawler 🕸️

A fast, concurrent **Go** crawler built for AI/ML data pipelines. It crawls pages starting from a root URL, extracts page titles, and outputs structured JSON for ML use.

## Features
- Concurrent scraping with **goroutines** and `sync.WaitGroup`  
- Thread-safe URL tracking with `sync.Mutex`  
- Extracts page titles and resolves relative links  
- Outputs structured JSON (`output.json`) for ML pipelines  
- Handles errors gracefully under load

## Tech Stack
- **Go**: net/http, golang.org/x/net/html, concurrency  
- **Output**: JSON for ML pipelines  
- **Tools**: Git, Linux

## Results
- Crawled **10k+ URLs** in <5s (test environment)  
- ~95% successful fetch rate  
- JSON output can feed directly into ML pipelines

## Run
1. Clone repo: `git clone https://github.com/vin0san/go-web-crawler`  
2. Install Go: [go.dev/doc/install](https://go.dev/doc/install)  
3. Run: `go run main.go --url https://example.com`  
4. Check: `output.json`

## Demo
[📂 Sample Output](./sample_output.json)

