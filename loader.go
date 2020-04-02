package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	url          = flag.String("url", "", "Url to run.")
	totalReq     = flag.Int("total", 0, "Number of total requests to run.")
	concurentReq = flag.Int("concurrent", 0, "Number of concurrent requests to run.")
)

func cleanup() {
	fmt.Println("Cleanup...")
}

func running() {
	t1 := time.Now()
	makeRequest(*url, *concurentReq, *totalReq)
	t2 := time.Now()
	elapsed := t2.Sub(t1)
	fmt.Println("Total time to run:", elapsed)
}

func infinityRequests(url string, wg *sync.WaitGroup, s int) {
	defer wg.Done()
	tStart := time.Now()
	resp, err := http.Get(url)
	tFinish := time.Now()
	fmt.Println("HTTP Response Status:", s, resp.StatusCode, http.StatusText(resp.StatusCode), tStart.Second(), tStart.Nanosecond(), "|", tFinish.Second(), tFinish.Nanosecond())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	if len(*url) == 0 || *totalReq == 0 || *concurentReq == 0 {
		fmt.Println("getRekt: wrong number of arguments, use -h for help")
		os.Exit(1)
	}
	running()
}

func makeRequest(url string, concurentReq int, totalReq int) {

	var wg sync.WaitGroup
	alertMessage := `Number of concurrent requests, can't be bigger than number of total requests.`

	if concurentReq >= totalReq {
		fmt.Println(alertMessage)
		os.Exit(1)
	}

	wg.Add(totalReq)
	for i := 0; i < concurentReq; i++ {
		s := i
		// routines
		go func() {
			for {
				infinityRequests(url, &wg, s)
			}
		}()
	}
	wg.Wait()

	fmt.Println("Number of total requests to run:", totalReq)
	fmt.Println("Number of concurrent requests to run:", concurentReq)
}
