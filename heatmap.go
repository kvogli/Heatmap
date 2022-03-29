package main

import (
	"encoding/csv"
	"fmt"
	"net/http"

	// "github.com/experiencedHuman/heatmap/LRZscraper"
	"github.com/experiencedHuman/heatmap/RoomFinder"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not retrieve csv data from URL!")
		return nil, err
	} else {
		fmt.Println("Successfully retrieved data from URL!")
	}

	defer resp.Body.Close()
	csvReader := csv.NewReader(resp.Body)
	csvReader.Comma = ','
	var data []string
	for i := 0; i < 5; i++ {
		data, err = csvReader.Read()
		if err != nil {
			return nil, err
		} else {
			fmt.Printf("%+v\n", data)
		}
	}
	return make([][]string, 0), nil
}

func getGraphiteDataFromLRZ() {
	url := "http://graphite-kom.srv.lrz.de/render/?xFormat=%d.%m.%20%H:%M&tz=CET&from=-5days&target=cactiStyle(group(alias(ap.gesamt.ssid.eduroam,%22eduroam%22),alias(ap.gesamt.ssid.lrz,%22lrz%22),alias(ap.gesamt.ssid.mwn-events,%22mwn-events%22),alias(ap.gesamt.ssid.@BayernWLAN,%22@BayernWLAN%22),alias(ap.gesamt.ssid.other,%22other%22)))&format=csv"

	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		if idx == 6 {
			break
		}

		fmt.Println(row[2])
	}
}

func main() {
	// LRZscraper.ScrapeData()
	fmt.Println("hi")
	RoomFinder.Crawl("https://golang.org/", 4, fetcher)
	RoomFinder.Scrape()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
