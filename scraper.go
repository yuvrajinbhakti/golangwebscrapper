package main

import (
    "fmt"
    "os"
    // import Colly
    "github.com/gocolly/colly"
)
func main() {
    args := os.Args
    url := args[1]
    collector := colly.NewCollector()

  // whenever the collector is about to make a new request
    collector.OnRequest(func(r *colly.Request) { 	
		/*
		The OnRequest is a method of the Collector object, and it registers a function to be called whenever a new request is about to be made by the collector.

We are passing it an anonymous function that takes a *colly.Request object as an argument (representing the request that is about to be made). Inside this function, we are using the fmt.Println function to print out the message "Visiting" along with the URL of the request that is currently being processed in the command line.
Essentially, whenever the collector is about to make a request, this function will be executed, and it will print a message indicating that the collector is visiting a particular URL. This can be useful for debugging and monitoring the progress of the web scraping process, allowing you to see which URLs are being visited by the collector.
		*/
        // print the url of that request
        fmt.Println("Visiting", r.URL)
    })
    collector.OnResponse(func(r *colly.Response) {
        fmt.Println("Got a response from", r.Request.URL)
    })
    collector.OnError(func(r *colly.Response, e error) {
        fmt.Println("Blimey, an error occurred!:", e)
    })
    collector.Visit(url)

}
// go run scraper.go https://ricette.giallozafferano.it/Schiacciata-fiorentina.html