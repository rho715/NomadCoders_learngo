// url checker & go routines
package challenges3

import (
	"fmt"
	"net/http"
)

// // 3.0 ~ 3.1
// var errRequestFailed = errors.New("Request Failure")

// func main() {
// 	results := map[string]string{}
// 	urls := []string{
// 		"https://www.naver.com",
// 		"https://www.google.com",
// 		"https://www.amazon.com",
// 		"https://www.instagram.com",
// 		"https://academy.nomadcoders.co",
// 	}
// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAIL"
// 		}
// 		results[url] = result
// 	}
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}

// }

// func hitURL(url string) error {
// 	fmt.Println("Checking:", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		return errRequestFailed
// 	}
// 	return nil
// }

// // 3.2
// func main() {
// 	go sexyCount("rho")
// 	go sexyCount("yoon")
// 	time.Sleep(time.Second * 3) // go routine is alive only when main function is alive
// }
// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

// // 3.3
// func main() {
// 	c := make(chan bool)
// 	people := [2]string{"rho", "yoon"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	//result := <-c
// 	//fmt.Println(result)

// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }

// func isSexy(person string, c chan bool) {
// 	time.Sleep(time.Second * 5)
// 	fmt.Println(person)
// 	c <- true
// }

// // 3.4
// func main() {
// 	c := make(chan string)
// 	people := [2]string{"rho", "yoon"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	//result := <-c
// 	//fmt.Println(result)
// 	fmt.Println("waiting for msg")
// 	resultOne := <-c
// 	fmt.Println("This is a first message : ", resultOne)
// 	resultTwo := <-c
// 	fmt.Println("This is a second message : ", resultTwo)
// 	// resultThree := <-c
// 	// fmt.Println("This is a second message : ", resultThree)

// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + " is sexy"
// }

// // 3.4 - loop
// func main() {
// 	c := make(chan string)
// 	people := [4]string{"rho", "yoon", "eun", "young"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println("waiting for", i)
// 		fmt.Println(<-c) //receiving msg : blocking operation
// 	}

// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- person + " is sexy"
// }

// 3.6 ~ 3.7

type requstResult struct {
	url    string
	status string
}

func challenges3() {
	results := make(map[string]string)
	c := make(chan requstResult)
	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.instagram.com",
		"https://academy.nomadcoders.co",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requstResult) { // chan<- : this channel Send Only

	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAIL"
	}
	c <- requstResult{url: url, status: status}
}
