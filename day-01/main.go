package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	f := ParseFileToStringByNewLine("input.txt")

	first, err := FirstStar(f)

	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return
	}
	fmt.Printf("First star results: %v\n", first)

	second, err := SecondStar(f)

	if err != nil {
		fmt.Printf("An error occurred: %v", err)
		return
	}
	fmt.Printf("Second star results: %v", second)
}

func FirstStar(f []string) (int, error) {
	distance := 0
	lcol, rcol := parseCols(f)
	sort.Slice(lcol, func(i, j int) bool {
		return lcol[i] < lcol[j]
	})

	sort.Slice(rcol, func(i, j int) bool {
		return rcol[i] < rcol[j]
	})
	
	for i := range lcol {
		distance += int(math.Abs(float64(lcol[i] - rcol[i])))
	}

	return distance, nil
}

func SecondStar (f []string) (int, error) {
	distance := 0
	lcol, rcol := parseCols(f)

	rcmap := map[int]int{}

	for _, rv := range rcol {
		_, ok := rcmap[rv]
		if ok {
			rcmap[rv]+= 1
		} else {
			rcmap[rv] = 1
		}
	}
	
	for _, lv := range lcol {
		val, ok := rcmap[lv] 
		if ok {
			distance += lv * val
		} 
	}

	return distance, nil
}

func parseCols (f []string) ([]int, []int) {
	lcol, rcol := []int{}, []int{}
	for _, line := range f {
		colvals := regexp.MustCompile("[0-9]{1,}").FindAllString(line, -1)
		if l, err := strconv.Atoi(colvals[0]); err == nil {
			lcol = append(lcol, l)
		}
		if r, err := strconv.Atoi(colvals[1]); err == nil {
			rcol = append(rcol, r)
		}
	}
	return lcol, rcol
}