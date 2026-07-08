package main

import (
	"net/http"
	"strings"
)

func detectTech(resp *http.Response) []string {
	var detected map[string]bool = map[string]bool{}

	server := strings.ToLower(resp.Header.Get("Server"))

	if strings.Contains(server, "nginx") {
		detected["nginx"] = true
	}
	if strings.Contains(server, "apache") {
		detected["apache"] = true
	}
	if strings.Contains(server, "litespeed") {
		detected["litespeed"] = true
	}
	if strings.Contains(server, "caddy") {
		detected["caddy"] = true
	}
	if strings.Contains(server, "gunicorn") {
		detected["gunicorn"] = true
	}
	if strings.Contains(server, "uvicorn") {
		detected["uvicorn"] = true
	}
	if strings.Contains(server, "roadrunner") {
		detected["roadrunner"] = true
	}
	if strings.Contains(server, "iis") {
		detected["iis"] = true
	}
	if strings.Contains(server, "gws") {
		detected["gws"] = true
	}

	poweredBy := strings.ToLower(resp.Header.Get("X-Powered-By")) // Header.Get() only returns the first value as a string

	if strings.Contains(poweredBy, "php") {
		detected["php"] = true
	}
	if strings.Contains(poweredBy, "asp.net") {
		detected["asp.net"] = true
	}
	if strings.Contains(poweredBy, "express") {
		detected["express"] = true
	}

	for _, cookie := range resp.Header.Values("Set-Cookie") { // Header.Values() returns the entire values
		c := strings.ToLower(cookie)

		if strings.Contains(c, "phpsessid") {
			detected["php"] = true
		}
		if strings.Contains(c, "laravel_session") {
			detected["laravel"] = true
			detected["php"] = true
		}
		if strings.Contains(c, "jsessionid") {
			detected["java"] = true
		}
		if strings.Contains(c, "asp.net") || strings.Contains(c, "aspsession") || strings.Contains(c, "aspx") {
			detected["asp.net"] = true
		}
		if strings.Contains(c, "csrftoken") {
			detected["django"] = true
			detected["python"] = true
		}
		if strings.Contains(c, "connect.sid") {
			detected["express"] = true
			detected["nodejs"] = true
		}
	}

	if resp.Header.Get("CF-Ray") != "" {
		detected["cloudflare"] = true
	}
	if resp.Header.Get("X-Sucuri-ID") != "" {
		detected["sucuri"] = true
	}
	if resp.Header.Get("X-Akamai-Transformed") != "" {
		detected["akamai"] = true
	}

	var tech []string

	for t, _ := range detected {
		tech = append(tech, t)
	}

	return tech

}

func detectInterestingHeaders(resp *http.Response) []string {
	interesting := make(map[string]bool)

	for key, _ := range resp.Header {
		lkey := strings.ToLower(key)

		if strings.Contains(lkey, "cache") {
			interesting[key] = true
		}

		if strings.Contains(lkey, "server") {
			interesting[key] = true
		}

		if strings.Contains(lkey, "powered") {
			interesting[key] = true
		}

		if strings.Contains(lkey, "security") {
			interesting[key] = true
		}

		if strings.Contains(lkey, "cf-") || strings.Contains(lkey, "cloudflare") {
			interesting[key] = true
		}

		if strings.Contains(lkey, "access-control-allow") {
			interesting[key] = true
		}
	}

	var result []string

	for k, _ := range interesting {
		result = append(result, k)
	}

	return result
}
