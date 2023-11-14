# Go4Hackers-Vuln-Scanner

Go4Hackers is a simple vulnerability scanner written in Go (Golang) that focuses on identifying directory listing vulnerabilities in URLs based on a specified wordlist. It performs a directory listing attack scan to detect whether the content of the URL directories exposes directory listings.

## Features

- **Directory Listing Scan:** Performs a directory listing attack on URLs.
- **Wordlist-based Scanning:** Utilizes a provided wordlist to check directory paths.
- **More is coming soon.**


## Screenshot
![Image](https://cdn.discordapp.com/attachments/1166821043743236218/1174085940453576724/Screenshot_2023-11-14_233612.png?ex=65664fbd&is=6553dabd&hm=e987fbacd33b962314286a23af8b3c9fb1b405b2629a7d3ebc0cce5a2d0d4f01&)

## Usage For Linux
```
cd outputs/linux
chmod +x go4hackers-vuln-scanner
./go4hackers-vuln-scanner --help
```

## Usage For Windows
```
cd outputs/windows
icacls go4hackers-vuln-scanner.exe /grant everyone:F
./go4hackers-vuln-scanner.exe --help
```
