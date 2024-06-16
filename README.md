# head

> This is a lite version of the Unix tool `head`.
>
> Head is a command line tool that displays the first n lines or bytes of a file, where the user can provide the value for n. If no file or value for n is provided then it displays the first 10 lines from the standard input.

`go run cmd/main.go`

`go run cmd/main.go test.txt`

`go run cmd/main.go -n 1 test.txt`

`go run cmd/main.go -c 10 test2.txt`

`go run cmd/main.go -n 10 test.txt test2.txt`
