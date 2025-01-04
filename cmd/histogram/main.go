package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"strings"
)

// Write a program to generate N random integers between 0 and 99 inclusive. Then display a histogram of those numbers on the command line, with each line corresponding to a bucket of size 20.

// So if I asked for 10 numbers, and behind the scenes the program generated:
// 52, 15, 62, 26, 57, 15, 7, 81, 95, 49

// Then the command and output would look like this:

// $ histogram –-num 10
// ***
// *
// ***
// *
// **

func main() {
	max_num_flag := flag.Int("max", 99, "the largest random number to generate")
	bucket_count_flag := flag.Int("buckets", 5, "how many buckets to sort the numbers into")
	count_flag := flag.Int("count", 10, "how many numbers to generate")
	prefix_flag := flag.Bool("print-interval", true, "whether or not to incloude a bucket interval in the output")

	flag.Parse()

	max_number := *max_num_flag
	bucket_count := *bucket_count_flag
	count := *count_flag

	buckets := make([]int, bucket_count)

	interval_size := max_number/bucket_count + 1

	for i := 0; i < bucket_count; i++ {

	}

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(max_number + 1)
	}
	slices.Sort(numbers)
	for i := 0; i < count; i++ {
		for b := 0; b < bucket_count; b++ {
			if numbers[i] < (b+1)*interval_size {
				buckets[b]++
				break
			}
		}
	}

	fmt.Printf("Sorting %d numbers from 0 to %d into %d buckets\n", count, max_number, bucket_count)

	prefix_length := (len(fmt.Sprintf("%d", max_number)) * 2) + 5
	construct_prefix := func(x int) string {
		if !*prefix_flag {
			return ""
		}
		min := x * interval_size
		max := int(math.Min(float64(((x+1)*interval_size)-1), float64(max_number)))
		prefix := fmt.Sprintf("[%d-%d]: ", min, int(max))
		return fmt.Sprintf("%s%s", strings.Repeat(" ", prefix_length-len(prefix)), prefix)
	}

	for x, b := range buckets {
		fmt.Printf("%s%s\n", construct_prefix(x), strings.Repeat("*", b))
	}
}
