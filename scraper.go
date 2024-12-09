// package main

// import (
//     "fmt"
//     "os"
//     // import Colly
//     "github.com/gocolly/colly"
// )
// func main() {
//     args := os.Args
//     url := args[1]
//     collector := colly.NewCollector()

//   // whenever the collector is about to make a new request
//     collector.OnRequest(func(r *colly.Request) { 	
// 		/*
// 		The OnRequest is a method of the Collector object, and it registers a function to be called whenever a new request is about to be made by the collector.

// We are passing it an anonymous function that takes a *colly.Request object as an argument (representing the request that is about to be made). Inside this function, we are using the fmt.Println function to print out the message "Visiting" along with the URL of the request that is currently being processed in the command line.
// Essentially, whenever the collector is about to make a request, this function will be executed, and it will print a message indicating that the collector is visiting a particular URL. This can be useful for debugging and monitoring the progress of the web scraping process, allowing you to see which URLs are being visited by the collector.
// 		*/
//         // print the url of that request
//         fmt.Println("Visiting", r.URL)
//     })
//     collector.OnResponse(func(r *colly.Response) {
//         fmt.Println("Got a response from", r.Request.URL)
//     })
//     collector.OnError(func(r *colly.Response, e error) {
//         fmt.Println("Blimey, an error occurred!:", e)
//     })
//     collector.Visit(url)

// }
// // go run scraper.go https://ricette.giallozafferano.it/Schiacciata-fiorentina.html












package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/gocolly/colly"
)

// Define the Dictionary and RecipeSpecs structs to store the scraped data
type Dictionary map[string]string
type RecipeSpecs struct {
    difficulty, prepTime, cookingTime, servingSize, priceTier string
}

type Recipe struct {
    url, name      string
    ingredients    []Dictionary
    specifications RecipeSpecs
}

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Please provide a URL.")
        return
    }
    url := args[1]

    // Initialize the collector
    collector := colly.NewCollector()

    // Define the callbacks
    collector.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })
    
    collector.OnResponse(func(r *colly.Response) {
        fmt.Println("Got a response from", r.Request.URL)
    })
    
    collector.OnError(func(r *colly.Response, e error) {
        fmt.Println("Error occurred:", e)
    })

    // Define the HTML scraping logic
    collector.OnHTML("main", func(main *colly.HTMLElement) {
        recipe := Recipe{}
        ingredients_dictionary := Dictionary{}

        recipe.url = url
        recipe.name = main.ChildText(".gz-title-recipe")
        fmt.Println("Scraping recipe for:", recipe.name)

        main.ForEach(".gz-name-featured-data", func(i int, specListElement *colly.HTMLElement) {
            if strings.Contains(specListElement.Text, "Difficoltà: ") {
                recipe.specifications.difficulty = specListElement.ChildText("strong")
            }
            if strings.Contains(specListElement.Text, "Preparazione: ") {
                recipe.specifications.prepTime = specListElement.ChildText("strong")
            }
            if strings.Contains(specListElement.Text, "Cottura: ") {
                recipe.specifications.cookingTime = specListElement.ChildText("strong")
            }
            if strings.Contains(specListElement.Text, "Dosi per: ") {
                recipe.specifications.servingSize = specListElement.ChildText("strong")
            }
            if strings.Contains(specListElement.Text, "Costo: ") {
                recipe.specifications.priceTier = specListElement.ChildText("strong")
            }
        })

        main.ForEach(".gz-ingredient", func(i int, ingredient *colly.HTMLElement) {
            ingredients_dictionary[ingredient.ChildText("a")] = ingredient.ChildText("span")
        })
        recipe.ingredients = append(recipe.ingredients, ingredients_dictionary)
        fmt.Println("Ingredients: ", recipe.ingredients)

        // Optionally, you can store these in a database or file
    })

    // Start scraping the given URL
    collector.Visit(url)
}
