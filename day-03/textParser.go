package main

import (
	"bufio"
	"log"
	"os"
)

func ParseFileToBytes(path string) []byte {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal((err))
	}

	return f
}

func ParseFileToStringByNewLine(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal((err))
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return res
}
