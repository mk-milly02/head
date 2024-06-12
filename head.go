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
