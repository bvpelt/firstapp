package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
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

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(1)                               // set maximum number threads
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // get all available threads
}

func test19() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 1000; i++ { // Locks in single context - routine
		wg.Add(2)
		m.RLock()
		go sayHello01()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello01() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()

}

func test18() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	wg.Wait()
}

func test17() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye" // nameless go routine uses this value, race condition
	time.Sleep(100 * time.Millisecond)

	msg = "Hello" // passing value to nameless go routine
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}

func test16() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func test15() {
	fmt.Println("Hello")

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

// interface is struct
// interface has methods, not data
type Writer interface {
	Write([]byte) (int, error)
}

// Implicit implement Writer interface
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func test14() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is: ", g.name)

	g.greets()
	fmt.Println("The new name is: ", g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greets() { // passing context g greeter by reference
	fmt.Println(g.greeting, g.name)
	g.name = "set by greets"
}
func (g greeter) greet() { // passing context g greeter by value
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

func test13() {
	var f func(msg string) = func(msg string) {
		fmt.Println(msg)
	}
	f("Hello Go!")

	var d func(a, b float64) (float64, error) = divide
	x, err := d(5.0, 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x)
}

func test12() { // error handling
	d, err := divide(5.0, 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func test11() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is: ", s)
}

func sum(values ...int) int {
	fmt.Println("values: ", values)
	result := 0

	for _, v := range values {
		result += v
	}

	fmt.Println("The sum is: ", result)

	return result
}

func test10() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", "yes", i)
	}
}

func sayMessage(msg, msg1 string, idx int) {
	fmt.Println(msg, msg1)
	fmt.Println("The value of the index is: ", idx)
}

func test09() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a // b points to address of map a!
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

func test08() {
	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
	ms = new(myStruct)
	(*ms).foo = 48
	fmt.Println(ms)
	ms.foo = 52
	fmt.Println(ms, ms.foo)

	a := []int{1, 2, 3} // if array has size b is copy if array has no size (is a slice) b has address
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}

type myStruct struct {
	foo int
}

func test07() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)
}

func test06() { // Pointers
	var a int = 42
	var b *int = &a
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 31
	fmt.Println(a, *b)
}

func test05() {
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
