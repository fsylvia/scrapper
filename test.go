package main

import (
	"fmt"
	"scrapper/colly"
	"scrapper/util"
	"scrapper/vanilla"
	"strings"
	"time"
)

const url = "https://www.justdial.com/LOCATION/KEYWORD"

func main() {
	var location = "Coimbatore"
	var keyword = "schools"
	fUrl := strings.Replace(url, "LOCATION", location, -1)
	fUrl = strings.Replace(fUrl, "KEYWORD", keyword, -1)
	//testVanillaCrawl(fUrl)
	testCollyCrawl(fUrl, location+"_"+keyword)
}

func testVanillaCrawl (fUrl string){
	response:= vanilla.Crawl(fUrl)
	dtlHrefs := vanilla.FetchDetailUrl(*response)
	fmt.Println(len(dtlHrefs))
	fmt.Println(dtlHrefs)
}

func testCollyCrawl (fUrl string , sName string) {
	defer util.TimeTrack(time.Now(), "testCollyCrawl")
	/*if strings.Index(url, page1) + len(page1) == len(url) {
		r.Abort()
	}else {
		r.Headers.Set("User-Agent", util.RandomString())
	}*/
	items := colly.Crawl(fUrl)
	util.WriteToFile(items, sName)
}

