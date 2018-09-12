package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	a := makeStringSlice(10000, 20)
	b := makeStringSlice(10000, 20)
	timeIt(compare, a, b)
	timeIt(compare2, a, b)

	c := makeStringSlice(2000, 10)
	d := makeStringSlice(2000, 10)
	timeIt(compare, c, d)
	timeIt(compare2, c, d)

	e := makeStringSlice(50000, 20)
	timeIt(compare, e, e)
	timeIt(compare2, e, e)
}

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func makeStringSlice(sliceLength int, wordLength int) []string {
	s := make([]string, sliceLength)
	for i := range s {
		s[i] = randomString(wordLength)
	}
	return s
}

type cf func(a, b []string) bool

func timeIt(f cf, a []string, b []string) {
	start := time.Now()
	fmt.Printf("Result: %v\n", f(a, b))
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Timer: %.4f\n", elapsed.Seconds())
}

func compare(a, b []string) bool {
	//5.9646s to compare 50,000 20-char strings
	fmt.Println("compare()")
	if len(a) != len(b) {
		return false
	}
	for i := len(a) - 1; i >= 0; i-- {
		for _, vD := range b {
			//fmt.Println("compare")
			if a[i] == vD {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}

	}
	if len(a) == 0 {
		return true
	}
	return false
}

func compare2(a, b []string) bool {
	//insanely fast comparison
	//0.0278s to compare 50,000 20-char strings
	fmt.Println("compare2()")
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
