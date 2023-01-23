// Package scrapper url checker & go routines
package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	title    string
	link     string
	location string
}

// Scrape function for jobkorea
func Scrape(term string) {
	var baseURL string = "https://www.jobkorea.co.kr/Search/?stext=" + term + "&tabType=recruit" //"https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	// getPages()
	totalPages := getPages(baseURL)
	// totalPages := 1
	//fmt.Println(totalPages)
	for i := 1; i <= totalPages; i++ {
		go getPage(i, baseURL, c) //getPage() returns array of jobs
	}
	for i := 0; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}

	writeJobs(jobs)
	fmt.Println("DONE! ( ", len(jobs), " Jobs Extracted) ")
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() // write data to a file

	headers := []string{"Title", "Link", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.title, "https://www.jobkorea.co.kr" + job.link, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	//https://www.jobkorea.co.kr/Search/?stext=python&tabType=recruit&Page_No=11
	pageURL := url + "&Page_No=" + strconv.Itoa(page)
	fmt.Println("Requesting : ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".recruit-info .list-post")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	title, _ := card.Find(".post-list-info>a").Attr("title")
	link, _ := card.Find(".post-list-info>a").Attr("href")
	location := card.Find(".post-list-info>p .loc.long").Text()
	c <- extractedJob{
		title:    CleanString(title),
		link:     link,
		location: CleanString(location)}

}

// CleanString cleans a string 
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.FindMatcher(goquery.Single(".tplPagination")).Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Find("a").Length())
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}
