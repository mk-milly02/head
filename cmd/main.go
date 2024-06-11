package main

import (
	"fmt"
	"head"
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	input := read_from_file(filename)
	output := head.ReadFirstNLines(input, 0)

	for _, line := range output {
		fmt.Print(string(line))
	}
}

func read_from_file(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s: no such file or directory", filename)
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
