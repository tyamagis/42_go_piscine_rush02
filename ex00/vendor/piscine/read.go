package piscine

import (
	"os"
)

const readBufSize = 1024

// pathの中身をすべて1つのstringとして読み出す
func Read(path string) (string, bool) {
	f, oErr := os.Open(path)
	if oErr != nil {
		return "", false
	}
	defer f.Close()

	var buffer [readBufSize]byte
	data := []byte{}
	for {
		n, rErr := f.Read(buffer[:])
		if rErr != nil {
			if rErr.Error() == "EOF" {
				break
			}
			return "", false
		}
		data = append(data, buffer[:n]...)
	}
	return string(data), true
}
