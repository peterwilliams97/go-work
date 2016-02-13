package main

import (
	"encoding/csv"
	"fmt"
	"index/suffixarray"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// Returns: Sum of elements in `arr`
func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

// T int, // size of text
// P int, // size of pattern.
func make_pattern_text(T int, // size of text
	P int, // size of pattern.
	R int, // 1 / R = fraction of text made up of pattern repetitions
) ([]byte, []byte) {
	// ([]byte, []byte)

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
	return pattern, text
}

const MAX_TIME = 10

func test_sa_speed(T int, // size of text
	P int, // size of pattern.
	R int, // 1 / R = fraction of text made up of pattern repetitions
) (float64, float64) {

	M := int(T / (P * R))
	fmt.Printf("Test suffix array speed: T=%v, P=%v, R=%v, M=%v\n",
		T, P, R, M)
	if M <= 0 {
		panic("M must be greater than zero")
	}

	// text:  String to be indexed and searched
	// pattern:  Substring to search for
	pattern, text := make_pattern_text(T, P, R)

	var (
		count int
		start time.Time
	)

	var text_index *suffixarray.Index
	count = 0
	for start = time.Now(); time.Since(start).Seconds() < MAX_TIME; {
		text_index = suffixarray.New(text[:])
		count++
	}
	if count < 1 {
		panic("Count index")
	}
	duration_index := time.Since(start).Seconds() / float64(count)
	fmt.Printf(" count=%d,dt_index=%.1f,duration_index=%g\n",
		count, time.Since(start).Seconds(), duration_index)

	matches := []int{} // number of pattern matches
	count = 0
	for start = time.Now(); time.Since(start).Seconds() < MAX_TIME; {
		offsets := text_index.Lookup(pattern, -1)
		if len(offsets) != M {
			var msg string
			fmt.Sprintf(msg, "%d matched, expected %d", len(offsets), M)
			panic(msg)
		}
		matches = append(matches, len(offsets))
		count++
	}
	if count < 1 {
		panic("Count lookup")
	}
	duration_lookup_total := time.Since(start).Seconds()
	duration_lookup := duration_lookup_total / float64(count)
	fmt.Printf(" count=%d,dt_lookup=%.1f, duration_lookup=%g\n",
		count, duration_lookup_total, duration_lookup)

	n_matches := sum(matches)
	if n_matches < 1 {
		panic("matches")
	}

	duration_match := duration_lookup_total / float64(n_matches) // Duration per pattern match
	duration_char := duration_match / float64(P)

	fmt.Printf(" %7d %e %e %e\n", n_matches, duration_lookup, duration_match, duration_char)

	return duration_index, duration_lookup
}

func main() {
	R := 2
	T := 10 << 7
	results := [][]float64{}
	for P := 10; P <= 10*1000; P += 10 {
		duration_index, duration_lookup := test_sa_speed(T, P, R)
		results = append(results,
			[]float64{float64(T), float64(P), duration_index, duration_lookup})
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range results {
		rec := []string{}
		for _, v := range record {
			rec = append(rec, strconv.FormatFloat(v, 'g', -1, 64))
		}
		if err := w.Write(rec); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
