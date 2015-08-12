package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const TICK_MS = 2000

var transport = &http.Transport{
	DisableCompression: true,
	Dial: (&net.Dialer{
		Timeout:   10000 * time.Millisecond,
		KeepAlive: 30000 * time.Millisecond,
	}).Dial,
	TLSHandshakeTimeout: 10000 * time.Millisecond,
}

var httpClient = &http.Client{Transport: transport}
var errorMsg = regexp.MustCompilePOSIX(`[^:]*$`)

func humot(url string, responses chan<- string) {
	// fmt.Print("Go!")
	for {
		res, err := httpClient.Get(url)
		if err == nil {
			responses <- strconv.Itoa(res.StatusCode)
			res.Body.Close()
		} else {
			errStr := strings.TrimSpace(errorMsg.FindString(err.Error()))
			if errStr == "" {
				errStr = err.Error()
			}
			responses <- "Error(" + errStr + ")"
			// fmt.Print(err)
		}
	}
}

func main() {
	URL := ""
	if len(os.Args) > 1 {
		URL = os.Args[1]
	} else {
		fmt.Println("Usage: humot <url> [<concurrency>]")
		os.Exit(0)
	}

	CONCURRENCY := 250
	if len(os.Args) > 2 {
		arg2, err := strconv.Atoi(os.Args[2])
		if err == nil {
			CONCURRENCY = arg2
		}
	}

	fmt.Printf("Hitting %s with %d concurrent requests ... (^C to stop)\n", URL, CONCURRENCY)
	responses := make(chan string, CONCURRENCY)
	for i := 1; i < CONCURRENCY; i++ {
		go humot(URL, responses)
	}
	statusCount := make(map[string]int)
	total := 0
	ticker := time.NewTicker(time.Millisecond * TICK_MS).C
	startTime := time.Now()
	for {
		select {
		case s := <-responses:
			statusCount[s]++
			total++
		case t := <-ticker:
			// fmt.Println(t.Sub(startTime), "Requests:", total, "Responses:", statusCount)
			fmt.Print(t.Sub(startTime))
			for k, v := range statusCount {
				fmt.Printf(" | %v: %v", k, v)
			}
			fmt.Println(" | Total: ", total)
		}
	}
}
