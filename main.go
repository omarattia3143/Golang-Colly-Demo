package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Fact struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func main() {

	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(colly.AllowedDomains("www.factretriever.com"))

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("couldnt get id")
		}
		factDesc := element.Text

		fact := Fact{
			Id:          factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/honey-bee-facts")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allFacts)

	writeToJsonFile(&allFacts)
}

func writeToJsonFile(data *[]Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("can not write to json file")
	}
	ioutil.WriteFile("bees.json", file, 0644)
}
