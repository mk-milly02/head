package head

import "bytes"

func ReadFirstNLines(text []byte, n int) (nlines []string) {
	reader := bytes.NewBuffer(text)
	for i := 0; i < n; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		nlines = append(nlines, line)
	}
	return nlines
}

func ReadFirstCBytes(text []byte, c int) (result string) {
	reader := bytes.NewBuffer(text)
	result = string(reader.Bytes()[0 : c])
	return result
}
