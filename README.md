# JoeyScan4Me

A easy-to-use recon tool kit for subdomain enumeration, HTTP probing, web crawling, and screenshot capturing.
<img width="692" height="520" alt="image" src="https://github.com/user-attachments/assets/f68e172e-d69b-4d8d-b0ef-2c4304d13eea" />

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

## Output Files
The output files will be stored in the specified working directory (or current directory by default) with the following structure:
```
/output
├── example.com/
│   ├── subdomains.txt              # List of discovered subdomains
│   ├── up_subdomains.txt           # List of live HTTP services
│   ├── up_subdomains_with_tech.txt # Live services with technology detection
│   ├── crawling_results.txt        # Web crawling results
│   └── screenshots/
│       └── gowitness.sqlite3       # Screenshot database
│       └── screenshots...       # Screenshot images
```
