package main

import (
	"fmt"
	// "github.com/mattn/go-sqlite3"
	"bufio"
	"index/suffixarray"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
)

type Coord struct {
	x int
	y int
}

func (c *Coord) sum() int {
	return c.x + c.y

}

func test_coord() {
	a := Coord{y: 2, x: 3}
	b := Coord{5, 6}
	a_s := a.sum()
	fmt.Printf("test_coord: a=%s,b=%s\n", a, b)
	fmt.Printf("sum=%d\n", a_s)
}

func test_array() {
	var arr [10]int
	arr[1] = 3

	sl0 := arr[4:7]
	sl0[0] = 111

	fmt.Printf("arr=%s\n", arr)
	for i, x := range arr {
		fmt.Printf("%d: %d\n", i, x)
	}

	var slc []int
	for i := 0; i <= 16; i++ {
		slc = append(slc, i)
		fmt.Printf("slc %2d: %2d %2d \n", i, len(slc), cap(slc))
	}

	sl3 := []int{}
	for i := 0; i < 5; i++ {
		sl3 = append(sl3, i)
		fmt.Printf("sl3 %d: %d %d \n", i, len(sl3), cap(sl3))
	}

	sl2 := make([]int, 1, 10)
	fmt.Printf("sl2: len=%d,cap=%d\n", len(sl2), cap(sl2))
	sl2[0] = 5
	// sl2[2] = 4
}

func test_copy() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000)) + 10
	}
	sort.Ints(scores)
	worst := make([]int, 10)
	copy(worst[2:], scores[:15])
	fmt.Println(worst)
}

func test_map(n int) {
	m := map[string]int{
		"x": 12,
		"y": 13,
	}
	m["a"] = 1
	m["B"] = 7
	if n == 3 {
		goto Done
	}
	fmt.Printf("map: m=%d %s\n", len(m), m)

Done:
	for k, v := range m {
		fmt.Printf("map: %s: %d\n", k, v)
	}
}

const TOL = 1e-12

func Sqrt(x float64) (float64, int) {
	z := x / 2

	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
		if math.Abs(z*z-x) <= TOL {
			// fmt.Printf("\t   Sqrt(%f) converged after %d iterations\n", x, i)
			return z, i
		}
	}
	panic("Sqrt did not converge")
}

func test_sqrt() {
	for i := 1; i <= 9; i++ {
		x := float64(i)
		z, iter := Sqrt(x)
		fmt.Printf("Sqrt(%f)=%f, delta=%e, iter=%d\n", x, z, z*z-x, iter)
	}
}

func test_io() {
	file, err := os.Open("/Users/pcadmin/go-work/src/hello/hello.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// read the file
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line += 1
		fmt.Printf("%3d: %s\n", line, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func test_types() {
	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)
	fmt.Println(stra)
	fmt.Println(strb)
	fmt.Println(byts)
	// fmt.Println(stra.(type))
	// fmt.Println(strb.(type))
	// fmt.Println(byts.(type))
}

func test_sa() {

	data := []byte("i am a test i am a test i am a test i am a test i am a test")
	s := data[2:4]

	// create index for some data
	index := suffixarray.New(data)

	// lookup byte slice s
	offsets1 := index.Lookup(s, -1) // the list of all indices where s occurs in data
	offsets2 := index.Lookup(s, 3)  // the list of at most 3 indices where s occurs in data
	fmt.Println("test_sa")
	fmt.Println(string(s))
	fmt.Println(offsets1)
	fmt.Println(offsets2)
	for _, i := range offsets1 {
		m := data[i : i+2]
		fmt.Println(string(m))
	}
}

func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

// func test_sa_speed() {
// 	const (
// 		N = 1e7 // Max size of text
// 		M = s   // Size of repeated substrings in text
// 		T = 1e9
// 		R = false
// 	)

// 	fmt.Printf("test_sa speed: N=%e, M=%d\n", N, M)
// 	text := make([]byte, N, N) // string to be indexed and searched
// 	for i := 0; i < N; i++ {
// 		text[i] = byte(i % M)
// 		if R && i >= M {
// 			text[i] = byte(rand.Int31n(M))
// 		}
// 	}
// 	for i := M; i <= N-M; i += 2 * M {
// 		for r := 0; r <= 5; r++ {
// 			j := int(rand.Int31n(M-1)) + 1
// 			k := byte(rand.Int31n(0xFE))
// 			text[i+j] = text[i+j] + 1 + k
// 		}
// 	}
// 	pattern := text[1:M]
// 	fmt.Println(pattern)
// 	fmt.Println(text[M : M*2])

// 	for q := 0; q < 3; q++ {
// 		for n := M; n <= N; n *= 10 {
// 			fmt.Printf("$ %9d", n)
// 			start0 := time.Now()
// 			text_index := suffixarray.New(text[:n])
// 			fmt.Printf(": ")
// 			t := T / n // number of test repeats
// 			if t < 1 {
// 				t = 1
// 			}

// 			matches := []int{} // number of pattern matches
// 			start := time.Now()
// 			for i := 0; i < t; i++ {
// 				offsets := text_index.Lookup(pattern, -1)
// 				matches = append(matches, len(offsets))
// 			}

// 			duration_total := time.Since(start0).Seconds()
// 			duration_all := time.Since(start).Seconds() // Duration of all t tests
// 			duration := duration_all                    // Duration per pattern match
// 			n_matches := sum(matches) / len(matches)
// 			if n_matches > 0.0 {
// 				duration /= float64(n_matches)
// 			}
// 			period_all := duration_all / float64(t) // period
// 			period := duration / float64(t)

// 			fmt.Printf("%e %e (%e %9d %7d %e %e)\n", period, period_all, duration, t, n_matches,
// 				duration_all, duration_total)
// 		}
// 		fmt.Println()
// 	}
// }

func echo() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("%2d: %s\n", i, os.Args[i])
	}
}

func main() {
	echo()
	// fmt.Println("hello, Mr. w0rld!")
	// a := 0
	// b := 1
	// if a+1 == b {
	// 	fmt.Printf("Hello again, a=%d\n", a)
	// }

	// fmt.Printf("program=%s\n", os.Args[0])

	// test_coord()
	// test_array()
	// test_copy()
	// test_map(2)
	// test_map(3)
	// test_sqrt()
	// //test_io()
	// test_types()
	// test_sa()
	// for {
	// 	test_sa_speed()
	// }
}
