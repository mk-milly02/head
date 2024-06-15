# uniq-tool

> This is a lite version of the Unix tool `head`.
>
> The uniq utility reads the specified input_file comparing adjacent lines, writes a copy of each unique input line to the output_file. If input_file is a single dash ('-') or absent, the standard input is read. If output_file is absent, standard output is used for output. The second and succeeding copies of identical adjacent input lines are not written. Repeated lines in the input will not be detected if they are not adjacent, so it may be necessary to sort the files first.

`go run cmd/main.go`

`go run cmd/main.go test.txt`

`go run cmd/main.go -n 1 test.txt`

`go run cmd/main.go -c 10 test2.txt`

`go run cmd/main.go -n 10 test.txt test2.txt`
