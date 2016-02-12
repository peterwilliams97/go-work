package main

import (
	"fmt"
	"index/suffixarray"
	"math/rand"
	"time"
)

// Returns: Sum of elements in `arr`
func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

func test_sa_speed(P int) {
	const (
		T = 2e7 // Max size of text
		// P = 100  	// Size of repeated substrings in text
		L = 1e9 // (Size of text) x (number of test runs)
		R = 2   // 1 / R = fraction of text made up of pattern repetitions
	)

	fmt.Printf("Test suffix array speed: text=%e chars, pattern=%d chars, %e matches\n",
		T, P, T/R)

	text := make([]byte, T, T)    // String to be indexed and searched
	pattern := make([]byte, P, P) // Substring to search for

	for j := 0; j < P; j++ {
		pattern[j] = byte(j)
	}
	for i := 0; i < T; i++ {
		text[i] = pattern[i%P]
	}
	for i := P; i <= T-(R-1)*P; i += R * P {
		j := int(rand.Int31n(int32(P) - 1))
		k := 1 + byte(rand.Int31n(0xFE))
		text[i+j] += k
	}

	// di := -1
	// var dvp, dvt byte
	// for j := 0; j < P; j++ {
	// 	if pattern[j] != text[P + j] {
	// 		di = j
	// 		dvp = pattern[j]
	// 		dvt = text[P + j]
	// 		break
	// 	}
	// }

	// fmt.Printf("diff: j=%d,dvp=%d,dvt=%d\n", di, dvp, dvt)
	// fmt.Println(pattern[:10])
	// fmt.Println(text[P:P + 10])

	for n := P * R; n <= T; n *= 10 {
		fmt.Printf("%9d", n)
		//start0 := time.Now()
		text_index := suffixarray.New(text[:n])
		fmt.Printf(": ")
		n_repeats := L / n // number of test repeats
		if n_repeats < 1 {
			n_repeats = 1
		}

		matches := []int{} // number of pattern matches
		start := time.Now()
		for i := 0; i < n_repeats; i++ {
			offsets := text_index.Lookup(pattern, -1)
			matches = append(matches, len(offsets))
		}

		//duration_total := time.Since(start0).Seconds()

		duration := time.Since(start).Seconds() / float64(n_repeats)
		//duration := time.Since(start).Seconds() / float64(n_repeats) // Duration per repeat

		n_matches := sum(matches) / len(matches)
		if n_matches*R*P != n {
			msg := fmt.Sprintf("Unexpected number of matches: n_matches=%d, n / (R * P)=%d (%d / %d %d)",
				n_matches, n/(R*P), sum(matches), len(matches), n_repeats)
			panic(msg)
		}

		duration_match := duration / float64(n_matches) // Duration per pattern match
		duration_char := duration_match / float64(P)

		fmt.Printf("%7d %e %e %e %e\n", n_matches, duration, duration_match, duration_char,
			duration/float64(n))
		//duration_all, duration_total)
	}
}

func main() {
	for P := 10; P <= 10*1000; P += 10 {
		test_sa_speed(P)
	}
}
