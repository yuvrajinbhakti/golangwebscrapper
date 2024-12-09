# Web Scraper using Go and Colly

This project demonstrates how to build a web scraper using the Go programming language and the Colly library. In this example, the scraper extracts recipe data from an Italian recipe website.

## Installation

To get started, follow these steps:

1. Clone the repository
2. Install Go if you haven't already. You can follow the installation guide from the official Go documentation.
3. Run `go mod init webscraper` to initialize the Go module
4. Install the Colly library using the following command:
   ```bash
   go get github.com/gocolly/colly


Example
go run scraper.go https://ricette.giallozafferano.it/Schiacciata-fiorentina.html
