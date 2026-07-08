package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

func worker(jobs <-chan string, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		start := time.Now()

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("error occured:", err)
			continue
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0 Safari/537.36")
		req.Header.Set("Referer", url)

		resp, err := (*client).Do(req)
		if err != nil {
			fmt.Println("error occured:", err)
			continue
		}
		latency := time.Since(start)

		var body []byte
		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Println("error occured:", err)
			continue
		}

		contentType := resp.Header.Get("Content-Type")

		var title string
		if strings.Contains(contentType, "text/html") {
			title = title_extractor(body)
		}

		var contentLength int64
		contentLength = resp.ContentLength

		if contentLength == -1 {
			contentLength = int64(len(body))
		}

		// tech := detectTech(resp)

		result := Result{
			URL:           url,
			Alive:         true,
			StatusCode:    resp.StatusCode,
			Protocol:      resp.Proto,
			PageTitle:     title,
			ContentType:   contentType,
			ContentLength: contentLength,
			LatencyMs:     latency.Milliseconds(),

			Tech:              detectTech(resp),
			InterestingHeader: detectInterestingHeaders(resp),
			Headers:           resp.Header,
		}

		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Println("json error:", err)
			continue
		}

		fmt.Println(string(jsonData))
	}
}
