package main

import (
	"flag"
	"fmt"
	"head"
	"io"
	"log"
	"os"
)

func main() {
	n := flag.Int("n", 10, "print the first NUM lines instead of the first 10;")
	c := flag.Int("c", 0, "print the first NUM bytes of each file")
	flag.Parse()

	var input []byte

	if filename := flag.Arg(0); filename == "" {
		for i := 0; i < 10; i++ {
			fmt.Scanln(&input)
			fmt.Println(string(input))
		}
	} else {
		input = read_from_file(filename)
		if *c != 0 {
			output := head.ReadFirstCBytes(input, *c)
			fmt.Print(output)
		} else {
			output := head.ReadFirstNLines(input, *n)
			for _, line := range output {
				fmt.Print(string(line))
			}
		}
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
