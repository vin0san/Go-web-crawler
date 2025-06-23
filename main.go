package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"golang.org/x/net/html"
)

func main() {
	startURL := "https://4chan.org"
	maxDepth := 20

	// Create a file to store results
	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Println("Starting web crawler on:", startURL, "with max depth:", maxDepth)

	// Create a wait group to manage goroutines
	var wg sync.WaitGroup
	visited := make(map[string]bool)
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		crawl(startURL, maxDepth, &wg, visited, &mu, startURL, file)
	}()

	wg.Wait()
	fmt.Println("Crawling completed.")
}

func crawl(currentURL string, depth int, wg *sync.WaitGroup, visited map[string]bool, mu *sync.Mutex, base string, file *os.File) {
	if depth <= 0 {
		return
	}

	// Check and mark visited
	mu.Lock()
	if visited[currentURL] {
		mu.Unlock()
		return
	}
	visited[currentURL] = true
	mu.Unlock()

	fmt.Println("Crawling:", currentURL)

	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Println("Failed to fetch:", currentURL, "Error:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Failed to parse HTML from:", currentURL, "Error:", err)
		return
	}

	// Extract title
	title := getTitle(doc)
	fmt.Println("Title of", currentURL, "is:", title)

	// Write the title to the file
	mu.Lock()
	fmt.Fprint(file, "Title of ", currentURL, " is: ", title, "\n")
	mu.Unlock()

	links := getLinks(doc)
	fmt.Println("Found", len(links), "links on", currentURL)

	baseParsed, err := url.Parse(currentURL)
	if err != nil {
		fmt.Println("Invalid base URL:", currentURL)
		return
	}

	for _, href := range links {
		linkParsed, err := url.Parse(href)
		if err != nil {
			continue
		}
		absURL := baseParsed.ResolveReference(linkParsed).String()

		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			crawl(link, depth-1, wg, visited, mu, base, file)
		}(absURL)
	}
}

//function to extract links from HTML nodes

func getLinks(n *html.Node) []string {
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return links
}

// function to extract title from HTML nodes

func getTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title := getTitle(c)
		if title != "" {
			return title
		}
	}
	return ""
}
