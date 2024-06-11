package head

import "bytes"

func ReadFirstNLines(text []byte, n int) (nlines []string) {
	reader := bytes.NewBuffer(text)
	if n == 0 {
		for i := 0; i < 10; i++ {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			nlines = append(nlines, line)
		}
	} else {
		for i := 0; i < n; i++ {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			nlines = append(nlines, line)
		}
	}
	return nlines
}
