# Go Web Crawler

A lightweight, concurrent web crawler written in Golang. This tool starts from a given URL and recursively fetches linked pages up to a specified depth, extracting page titles and writing them to a file.

---

## Features

-  Crawls web pages starting from a root URL
-  Extracts and resolves relative links
-  Retrieves and logs page titles
-  Uses goroutines and `sync.WaitGroup` for concurrent crawling
-  Thread-safe URL tracking using `sync.Mutex`
-  Outputs crawl results to `results.txt`

---

## Technologies Used

- **Golang**
- `net/http` – for HTTP requests  
- `golang.org/x/net/html` – for HTML parsing  
- `sync` – for goroutines, mutexes, and concurrency control  
- `os`, `url`, `fmt` – for I/O and utilities

---

---

## How to Run

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-web-crawler.git
cd go-web-crawler
```

### 2. Run the crawler
```bash
go run main.go
```
By default, it starts crawling from `https://4chan.org` up to a depth of 20.
you can change the URL to the one you want to crawl.

## sample Output
```bash
Title of https://4chan.org is: 4chan
Title of https://boards.4channel.org is: 4chan Boards
...
```
---
### Note: this project is meant for educational purpose or practice.
