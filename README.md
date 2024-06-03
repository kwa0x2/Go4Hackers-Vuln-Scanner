# Go4Hackers-Vuln-Scanner

[![Go 1.21x](https://img.shields.io/badge/Go-1.21.x-blue.svg)](https://go.dev/) [![License](https://img.shields.io/badge/License-MIT%20License-red.svg)](https://raw.githubusercontent.com/kwa0x2/Go4Hackers-Vuln-Scanner/master/LICENSE) [![Website](https://img.shields.io/badge/Website-nettasec.com-red.svg)](https://nettasec.com/)

Go4Hackers is a simple vulnerability scanner written in Go (Golang) that focuses on identifying directory listing vulnerabilities in URLs based on a specified wordlist. It performs a directory listing attack scan to detect whether the content of the URL directories exposes directory listings.

## Features

- **Directory Listing Scan:** Performs a directory listing attack on URLs.
- **Wordlist-based Scanning:** Utilizes a provided wordlist to check directory paths.
- **X-Frame-Options Header Checker:** Checks for the presence of the X-Frame-Options header to prevent Clickjacking attacks.
- **HTTP Trace Method Checker:** Verifies if the HTTP TRACE method is enabled, potentially exposing security vulnerabilities.
- **More is coming soon.**

## Installation
```
go install -v github.com/kwa0x2/go4hackers-vuln-scanner@latest
```

## Screenshot
![Image](https://i.hizliresim.com/m8b5zyz.png)

![Image](https://i.hizliresim.com/pzo58q3.png)
