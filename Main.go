package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

var (
	counter int = 0
)

func main() {
	args := os.Args
	fmt.Printf("All arguments: %v\n", args)
	argsWithoutProgram := os.Args[1:]
	fmt.Printf("Arguments without program name: %v\n", argsWithoutProgram)

	fmt.Printf("counter: %v\n", counter)
	var counter int = 10

	fmt.Printf("counter: %v\n", counter)
	fmt.Printf("actor %v\n", actorName)

	var myURL = "https://nu.nl"
	if len(argsWithoutProgram) == 1 {
		myURL = argsWithoutProgram[0]
	}

	// get web page
	getWebPage(myURL)
}

func getWebPage(page string) {
	fmt.Printf("Show page %v\n", page)

	resp, err := http.Get(page)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Printf("\n=========== Headers  =========\n")
	for k, v := range resp.Header {

		fmt.Printf("%s %s\n", k, v)
	}
	fmt.Printf("\n=========== Headers  =========\n")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("=========== Response =========\n%v\n=========== Response =========\n", string(body)[0:80])
}
