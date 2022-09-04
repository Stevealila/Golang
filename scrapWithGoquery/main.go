package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	url := "https://profound-caramel-e3f140.netlify.app/"

	// make a request to the website
	response, err := http.Get(url)
	checkError(err)

	// Close the response body as soon as the scraping process is complete
	defer response.Body.Close()

	// parse the HTML
	remoteHTML, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	// Create a csv file to store the sieved data
	file, err := os.Create("Golang_scraped.csv")
	checkError(err)
	defer file.Close()
	csvWriter := csv.NewWriter(file)

	// dissect the target tags [h2, img src, figcaption] from the parsed HTML
	remoteHTML.Find("article").Each(func(index int, article *goquery.Selection) {

		articleTitle := strings.TrimSpace(article.Find("h2").Text())
		imgSource, _ := article.Find("figure").Find("img").Attr("src")
		imgDescription := strings.TrimSpace(article.Find("figure").Find("figcaption").Text())

		// save the data in the CSV file
		finalData := []string{articleTitle, imgDescription, imgSource}
		csvWriter.Write(finalData)

		// print the saved data on the console output
		fmt.Printf("%v: %v: %v", articleTitle, imgDescription, imgSource)
		fmt.Print("\n")
	})

	csvWriter.Flush()

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

