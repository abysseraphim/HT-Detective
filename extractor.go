package main

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

func title_extractor(body []byte) string {

	reader := bytes.NewReader(body)
	document, err := html.Parse(reader)
	if err != nil {
		return ""
	}

	return findTitleDFS(document)
}

func findTitleDFS(node *html.Node) string {
	if node == nil {
		return ""
	}

	if node.Type == html.ElementNode && node.Data == "title" {
		if node.FirstChild != nil {
			return strings.TrimSpace(node.FirstChild.Data)
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		title := findTitleDFS(child)

		if title != "" {
			return title
		}
	}

	return ""
}
