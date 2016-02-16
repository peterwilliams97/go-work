package main

import (
	"encoding/csv"
	"fmt"
	"index/suffixarray"
	"log"
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// T int, // size of text
// P int, // size of pattern.
func make_pattern_text(T int, // size of text
	P int, // size of pattern.
	R int, // 1 / R = fraction of text made up of pattern repetitions
) ([]byte, []byte) {
	// ([]byte, []byte)

	M := int(T / (P * R))

	text := make([]byte, T, T)    // String to be indexed and searched
	pattern := make([]byte, P, P) // Substring to search for

	for j := 0; j < P; j++ {
		pattern[j] = byte(j%0xFF + 1)
	}
	// for j := P - 10; j < P; j++ {
	// 	fmt.Printf("%5d: %3d\n", j, pattern[j])
	// }

	for m := 0; m < M; m++ {
		t0 := m * P * R
		for j := 0; j < P; j++ {
			text[t0+j] = pattern[j]
		}
		for r := 1; r < R; r++ {
			t1 := t0 + P*r
			//text[t1+P-1] = pattern[P-1] + 1
			for j := 0; j < P-1; j++ {
				text[t1+j] = pattern[j]
			}
			//text[t1+P-1] = pattern[P-1] + 1
			// 	for j := P - 10; j < P; j++ {
			// 		fmt.Printf("%5d: %3d %3d\n", j, pattern[j], text[t1+j])
			// 	}
			// 	fmt.Printf("t0=%v, t1=%v\n", t0, t1)
			// 	panic("!!!!")
		}

		// for r := 1; r < R; r++ {
		// 	t1 := t0 + P*r
		// 	text[t1] = text[t1] + 1
		// 	// t2 := min(T, t0+P*r)
		// 	// if t2 < t1 {
		// 	// 	panic("t2 < t1")
		// 	// }
		// 	// if t2 <= t1 {
		// 	// 	break
		// 	// }
		// 	// var j int
		// 	// if t2-t1 == 1 {
		// 	// 	j = 0
		// 	// } else {
		// 	// 	j = int(rand.Int31n(int32(t2-t1) - 1))
		// 	// }

		// 	// //k := 1 + byte(rand.Int31n(0xFE))
		// 	// text[t1+j]++ // += k
		// }
	}

	return pattern, text
}

const MAX_TIME = 10

func test_sa_speed(T int, // size of text
	P int, // size of pattern.
	R int, // 1 / R = fraction of text made up of pattern repetitions
) (float64, float64) {

	M := int(T / (P * R))
	fmt.Printf("Test suffix array speed: T=%v, P=%v, R=%v, M=%v, P*R*M=%v, Δ=%v\n",
		T, P, R, M, P*R*M, T-P*R*M)
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
	// fmt.Printf(" count=%d,dt_index=%.1f,duration_index=%g\n",
	// 	count, time.Since(start).Seconds(), duration_index)

	// matches := []int{} // number of pattern matches
	count = 0
	for start = time.Now(); time.Since(start).Seconds() < MAX_TIME; {
		offsets := text_index.Lookup(pattern, -1)
		if len(offsets) != M {
			panic(fmt.Sprintf("%d matched, expected %d, count=%d", len(offsets), M, count))
		}
		// matches = append(matches, len(offsets))
		count++
	}
	if count < 1 {
		panic("Count lookup")
	}
	duration_lookup_total := time.Since(start).Seconds()
	duration_lookup := duration_lookup_total / float64(count)
	// fmt.Printf(" count=%d,dt_lookup=%.1f, duration_lookup=%g\n",
	// 	count, duration_lookup_total, duration_lookup)

	// n_matches := sum(matches)
	// if n_matches < 1 {
	// 	panic("matches")
	// }

	// duration_match := duration_lookup_total / float64(n_matches) // Duration per pattern match
	// duration_char := duration_match / float64(P)

	// fmt.Printf(" %7d %e %e : %e %e\n", n_matches, duration_index, duration_lookup,
	// 	duration_match, duration_char)

	return duration_index, duration_lookup
}

func write_results(results [][]float64) {
	w := csv.NewWriter(os.Stdout)
	if err := w.Write([]string{"T", "P", "index", "lookup"}); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

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

func main() {
	R := 2
	T_MAX := 10 * 1000 * 1000
	// T_MAX := 1000

	results_P := [][]float64{}
	for P := 1; P*R <= T_MAX; P *= 2 {
		T := T_MAX
		duration_index, duration_lookup := test_sa_speed(T, P, R)
		results_P = append(results_P,
			[]float64{float64(T), float64(P), duration_index, duration_lookup})
	}

	P := 100
	results_T := [][]float64{}
	for T := P * R; T <= T_MAX; T *= 2 {
		duration_index, duration_lookup := test_sa_speed(T, P, R)
		results_T = append(results_T,
			[]float64{float64(T), float64(P), duration_index, duration_lookup})
	}

	write_results(results_P)
	write_results(results_T)
}
