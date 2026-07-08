package main

type Result struct {
	URL           string `json:"url"`
	Alive         bool   `json:"alive"`
	StatusCode    int    `json:"status_code"`
	Protocol      string `json:"protocol"`
	PageTitle     string `json:"page_title"`
	ContentType   string `json:"content_type"`
	ContentLength int64  `json:"content_length"`
	LatencyMs     int64  `json:"latency_ms"`

	Tech              []string            `json:"tech,omitempty"`
	InterestingHeader []string            `json:"interesting_headers,omitempty"`
	Headers           map[string][]string `json:"entire_headers"`
}
