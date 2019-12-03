package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func add(x int, y int) int {
	return x + y
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Sqrt(x float64) float64 {
	for z := float64(1); z*z <= x; z += 1 {
		if z*z == x {
			return z
		}
	}

	return -1
}

func show_os() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func check_day() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func check_day2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func test_defer() {
	defer fmt.Println("world") // same as process.nextTick() in NodeJS

	fmt.Println("hello")
}

func test_pointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

func test_struc() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func test_struc_pointer() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func test_array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func test_array_slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func test_array_reference() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func test_array_slice2() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func test_array_slice3() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func test_array_slice4() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func test_array_slice5() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func test_array_slice6() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

var pow2 = []int{1, 2, 4, 8, 16, 32, 64, 128}

func test_range() {
	for i, v := range pow2 {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func test_range2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// func Pic(dx, dy int) [][]uint8 {
// }

// func test_array_slice7() {
// 	pic.Show(Pic)
// }

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func test_map() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

var m2 = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

func test_map2() {
	fmt.Println(m2)
}

var m3 = map[string]Vertex2{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func test_map3() {
	fmt.Println(m3)
}

func test_map4() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func test_function() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func test_function2() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func test_function3() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
// func fibonacci() func() int {
// }

// func test_fibonacci() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }

type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func test_method() {
	v := Vertex4{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func test_method2() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Vector struct {
	X, Y int
	z    interface{}
}

func test_type() {
	do(21)
	do("hello")
	do(true)

	var v interface{}

	v = "hello"

	v, ok := v.(string)

	vv := Vector{}

	fmt.Println(ok)

	fmt.Println("vv", vv)

	t := vv.z

	t, ok = t.(string)

	fmt.Println("t", t.(string), "ok", ok, t == nil)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %0.2f", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return 0, nil
}

func test_error() {
	fmt.Println("------------------------------------------------------ test_error")

	fmt.Println(Sqrt2(2))

	res, err := Sqrt2(-2)

	fmt.Println("res:", res, "err", err)
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

func test_error2() {
	fmt.Println("------------------------------------------------------ test_error2")

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	fmt.Println()
}

func test_map_types() {
	hosts := map[string]interface{}{
		"loopback":  []int{127, 0, 0, 1},
		"googleDNS": []string{"8"},
		"bla":       "bla",
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println([]int{0, 1, 2, 3})
}

func test_io() {
	fmt.Println("------------------------------------------------------ test_io")

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Println("b: ", string(b))
		fmt.Printf("n = %v err = %v b = %v len: %d\n", n, err, b, len(b))
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func test_goroutine() {
	fmt.Println("------------------------------------------------------ test_goroutine")

	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func test_channel() {
	fmt.Println("------------------------------------------------------ test_channel")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func test_switch_default() {
	fmt.Println("------------------------------------------------------ test_switch_default")

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	var i = time.Now()

	for {
		select {
		case i = <-tick:
			fmt.Println("tick. i:", i)
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    . i:", i)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func test_select() {
	fmt.Println("------------------------------------------------------ test_select")

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func walkImpl(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walkImpl(t.Left, ch)
	ch <- t.Value
	walkImpl(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// Need to close the channel here
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// NOTE: The implementation leaks goroutines when trees are different.
// See binarytrees_quit.go for a better solution.
func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func test_binary_tree() {
	fmt.Println("------------------------------------------------------ test_binary_tree")

	fmt.Print("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(1) != tree.New(2): ")
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func test_mutex() {
	fmt.Println("------------------------------------------------------ test_mutex")

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fetched tracks URLs that have been (or are being) fetched.
// The lock must be held while reading from or writing to the map.
// See https://golang.org/ref/spec#Struct_types section on embedded types.
var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("url load in progress") // sentinel value

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	}

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("<- Done with %v, already fetched.\n", url)
		return
	}
	// We mark the url to be loading to avoid others reloading it at the same time.
	fetched.m[url] = loading
	fetched.Unlock()

	// We load it concurrently.
	body, urls, err := fetcher.Fetch(url)

	// And update the status in a synced zone.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Printf("<- Error on %v: %v\n", url, err)
		return
	}
	fmt.Printf("Found: %s %q\n", url, body)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}
	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done
	}
	fmt.Printf("<- Done with %v\n", url)
}

func test_web_crawler() {
	fmt.Println("------------------------------------------------------ test_web_crawler")

	Crawl("https://golang.org/", 4, fetcher)

	fmt.Println("Fetching stats\n--------------")
	for url, err := range fetched.m {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
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

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now(), rand.Intn(10))

	fmt.Println(add(42, rand.Intn(10)))

	show_os()

	check_day()

	check_day2()

	test_defer()

	fmt.Println("Was defered?")

	test_pointer()

	test_struc()

	test_struc_pointer()

	fmt.Println("fmt.Println(v1, p, v2, v3) ⌵ ⌵ ⌵ ⌵ ⌵ ⌵")
	fmt.Println(v1, p, v2, v3)

	test_array()

	test_array_slice()

	test_array_reference()

	test_array_slice2()

	test_array_slice3()

	test_array_slice4()

	test_array_slice5()

	test_array_slice6()

	test_range()

	test_range2()

	// test_array_slice7()

	test_map()

	test_map2()

	test_map3()

	test_map4()

	test_function()

	test_function2()

	test_function3()

	// test_fibonacci()

	test_method()

	test_method2()

	test_type()

	test_map_types()

	test_error()

	test_error2()

	test_io()

	test_goroutine()

	test_channel()

	test_switch_default()

	test_select()

	test_binary_tree()

	test_mutex()

	test_web_crawler()

	vcm := &VCMessenger{}

	fmt.Println(vcm)
}

}
