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
	// defer is called before panic
	fmt.Println("start")
	panicker()
	panic("something bad happened")
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
}

func test04() {
	// defer is called before panic
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}

func test03() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("err.Error()")
	}
}

func test02() {
	var myURL = "https://www.google.com/robots.txt"
	res, err := http.Get(myURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func test01() {
	args := os.Args
	fmt.Printf("All arguments: %v\n", args)
	argsWithoutProgram := os.Args[1:]
	fmt.Printf("Arguments without program name: %v\n", argsWithoutProgram)

	fmt.Printf("counter: %v\n", counter)
	var counter int = 10

	fmt.Printf("counter: %v\n", counter)
	fmt.Printf("actor %v\n", actorName)

	/*

		if len(argsWithoutProgram) == 1 {
			myURL = argsWithoutProgram[0]
		}

		// get web page
		getWebPage(myURL)
	*/
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
