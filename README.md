# JoeyScan4Me

A easy-to-use tool for subdomain enumeration, HTTP probing, web crawling, and screenshot capturing all in one.

## Features

- **Subdomain Enumeration**: Uses Subfinder to discover subdomains
- **HTTP Probing**: Uses HTTPX to identify live web services
- **Web Crawling**: Uses Katana to crawl discovered websites
- **Screenshot Capture and Dashboard**: Uses Gowitness to capture screenshots with database storage and a web dashboard for easy viewing

## Execution flow

1. Subdomain Enumeration with Subfinder
2. HTTP Probing with HTTPX
3. Web Crawling with Katana
4. Screenshot Capture and Dashboard with Gowitness

## Installation

```bash
go build -o joesyscan4me cmd/JoeyScan4Me/main.go
```