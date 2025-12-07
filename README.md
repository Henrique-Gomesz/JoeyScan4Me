# JoeyScan4Me

A easy-to-use tool kit for subdomain enumeration, HTTP probing, web crawling, and screenshot capturing all in one.

## Features

- **Subdomain Enumeration**: Uses [Subfinder](https://github.com/projectdiscovery/subfinder) to discover subdomains
- **HTTP Probing**: Uses [HTTPX](https://github.com/projectdiscovery/httpx) to identify live web services
- **Web Crawling**: Uses [Katana](https://github.com/projectdiscovery/katana) to crawl discovered websites
- **Screenshot Capture and Dashboard**: Uses [Gowitness](https://github.com/sensepost/gowitness) to capture screenshots with database storage and a web dashboard for easy viewing

## Installation
You can install JoeyScan4Me using the following methods:
Require Go 1.21 or higher.

## From source
```bash
go install github.com/Henrique-Gomesz/JoeyScan4Me/cmd/JoeyScan4Me@latest
```
## Manual build
```bash
git clone https://github.com/Henrique-Gomesz/JoeyScan4Me.git
cd JoeyScan4Me
go build -o joeyscan4me cmd/joeyscan4me/main.go
```

## Usage
```bash
$ joeyscan4me -h

JoeyScan4Me - Simple and helpful recon toolkit

    |\__/,|   ('\
  _.|o o  |_   ) )
-(((---(((--------
by: Henrique-Gomesz


Usage:
  joeyscan4me [flags]

Flags:
   -d string  domain to scan (e.g. example.com)
   -w string  working directory for output files, defaults to current directory (default "./")
   -server    start gowitness server at the end of scan to view screenshots
```

## Example
Running a scan on example.com and starting the gowitness server at the end:
```bash
joeyscan4me -d example.com -w /path/to/output -server
```