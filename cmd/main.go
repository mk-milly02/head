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

	if filenames := flag.Args(); len(filenames) == 0 {
		for i := 0; i < 10; i++ {
			fmt.Scanln(&input)
			fmt.Println(string(input))
		}
	} else {
		if len(filenames) > 1 {
			for _, fn := range filenames {
				fmt.Printf("==> %s <==\n", fn)
				input = read_from_file(fn)
				if *c != 0 {
					output := head.ReadFirstCBytes(input, *c)
					fmt.Println(output)
				} else {
					output := head.ReadFirstNLines(input, *n)
					for _, line := range output {
						fmt.Print(string(line))
					}
				}
			}
		} else {
			input = read_from_file(filenames[0])
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
