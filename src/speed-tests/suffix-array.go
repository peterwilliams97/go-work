package main

import (
	"fmt"
	"time"
	"math/rand"
	"index/suffixarray"
)


func sum(arr[] int) int {
	s := 0
	for i:= 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

func test_sa_speed() {
	const (
		N = 2e7 	// Max size of text
		M = 200  	// Size of repeated substrings in text
		T = 1e9		// Size of text x number of test runs
		R = 2		// 1 / R = fraction of text made up of pattern repetitions
	)

	fmt.Printf("Test suffix array speed: text=%e chars, pattern=%d chars\n", N, M)

	text := make([]byte, N, N) 			// String to be indexed and searched
	pattern := make([]byte, M, M) 		// Substring to search for

	for j := 0; j < M; j++ {
		pattern[j] = byte(j)
	}
	for i := 0; i < N; i++ {
		text[i] = pattern[i % M]
	}
	for i := M; i <= N - (R - 1) * M; i += R * M {
		j := int(rand.Int31n(M - 1))
		k := 1 + byte(rand.Int31n(0xFE))
		text[i + j] +=  k
	}

	di := -1
	var dvp, dvt byte
	for j := 0; j < M; j++ {
		if pattern[j] != text[M + j] {
			di = j
			dvp = pattern[j]
			dvt = text[M + j]
			break
		}
	}

	fmt.Printf("diff: j=%d,dvp=%d,dvt=%d\n", di, dvp, dvt)
	fmt.Println(pattern[:10])
	fmt.Println(text[M:M + 10])

	for q := 0; q < 3; q++ {
		for n := M; n <= N; n *= 10 {
			fmt.Printf("$ %9d", n)
			start0 := time.Now()
			text_index := suffixarray.New(text[:n])
			fmt.Printf(": ")
			t := T / n // number of test repeats
			if t < 1 {
				t = 1
			}

			matches := []int{} // number of pattern matches
			start := time.Now()
			for i := 0; i < t; i++ {
				offsets := text_index.Lookup(pattern, -1)
				matches = append(matches, len(offsets))
			}

			duration_total := time.Since(start0).Seconds()
			duration_all := time.Since(start).Seconds() // Duration of all t tests
			duration := duration_all                    // Duration per pattern match
			n_matches := sum(matches) / len(matches)
			if n_matches > 0.0 {
				duration /= float64(n_matches)
			}
			period_all := duration_all / float64(t) // period
			period := duration / float64(t)

			fmt.Printf("%e %e (%e %9d %7d %e %e)\n", period, period_all, duration, t, n_matches,
				duration_all, duration_total)
		}
		fmt.Println()
	}
}

func main() {
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
	for {
		test_sa_speed()
	}
}