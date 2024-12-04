package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	f := ParseFileToStringByNewLine("input.txt")

	first, err := FirstStar(f)

	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
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
	valid := 0
	for _, v := range f {
		row := regexp.MustCompile("[0-9]{1,}").FindAllString(v, -1)
		direction := ""
		for i, rv := range row {
			if i == len(row)-1 {
				valid++
				direction = ""
				break
			}

			cur := strToInt(rv)
			next := strToInt(row[i+1])

			dif := cur - next
			if int(math.Abs(float64(dif))) > 3 {
				break
			}
			if dif == 0 {
				direction = ""
				break
			}
			if i == 0 {
				if dif < 0 {
					direction = "UP"
				} else {
					direction = "DOWN"
				}
			} else {
				if dif < 0 && direction == "DOWN" {
					direction = ""
					break
				}
				if dif > 0 && direction == "UP" {
					direction = ""
					break
				}
			}
		}
	}
	return valid, nil
}

func SecondStar(f []string) (int, error) {
	valid := 0
	for _, v := range f {
		row := regexp.MustCompile("[0-9]{1,}").FindAllString(v, -1)
		direction := ""
		original := []int{}
		sorted := []int{}
		

		for _, rv := range row {
			value := strToInt(rv)
			original = append(original, value)
			if len(sorted) == 0 {
				sorted = append(sorted, value)
				continue
			}
			if sorted[len(sorted)-1] != value {
				sorted = append(sorted, value)
			}
		}

		if len(original)-1 > len(sorted) {
			continue
		}

		if original[0] < original[1] {
			direction = "UP"
		} else if original[0] > original[1] {
			direction = "DOWN"
		} else {
			if original[0] < original[2] {
				direction = "UP"
			} else if original[0] > original[2] {
				direction = "DOWN"
			} else {
				break
			}
		}
		
		sortSlice(sorted, direction)
		diffCount := 0
		if len(original) > len(sorted) {
			diffCount = len(original) - len(sorted)
		} else {
			for i, v := range sorted {
				if v != original[i] {
					diffCount++
				}
			}
		}
		if diffCount > 2 {
			continue
		}
		
		for i, rv := range sorted {
			if i == len(sorted)-1 {
				valid++
				direction = ""
				break
			}

			cur := rv
			next := sorted[i+1]

			isValid := isValidChange(&direction, cur, next)
			if isValid {
				continue
			}
			break
		}
	}
	return valid, nil
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error occurred converting string to int: %v", err)
		os.Exit(1)
	}
	return i
}

func isValidChange(d *string, cur int, next int) bool {
	dif := cur - next

	if int(math.Abs(float64(dif))) > 3 {
		return false
	}

	if dif == 0 {
		return false
	}

	if dif < 0 && *d == "DOWN" {
		return false
	}
	if dif > 0 && *d == "UP" {
		return false
	}
	
	return true
}

func sortSlice(s []int, d string) []int {
	if d == "UP" {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	} else {
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
	}
	return s
}
