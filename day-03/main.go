package main

import (
	"fmt"
	"os"
	"regexp"
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

func FirstStar(f []string) (int, error){
	res := 0
	for _, v := range f {
		muls := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`).FindAllString(v, -1)
		for _, vv := range muls{
			curmul := regexp.MustCompile(`[0-9]{1,3}`).FindAllString(vv, -1)
			res += strToInt(curmul[0]) * strToInt(curmul[1])
		}
	}
	return res, nil
}

func SecondStar(f []string) (int, error){
	res := 0
	do := true
	for _, v := range f {
		muls := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|don't\(\)|do\(\)`).FindAllString(v, -1)
		for _, vv := range muls{
			if vv == "do()" {
				do = true
				continue
			}
			if vv == "don't()" {
				do = false
				continue
			}
			if do {
				curmul := regexp.MustCompile(`[0-9]{1,3}`).FindAllString(vv, -1)
				res += strToInt(curmul[0]) * strToInt(curmul[1])
			}
		}
	}
	return res, nil
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error occurred converting string to int: %v", err)
		os.Exit(1)
	}
	return i
}