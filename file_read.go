package main

import (
	"bufio"
	"fmt"
	"os"
)

var newLine = []byte{'\n'}

func file_read(malDir string) {
	filePath := malDir + "impls/cs/step3_env.cs"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	collector := newBytesAndLines(8 * 1024)

	var lineCount int
	scanner := bufio.NewScanner(file)
	scanner.Split(collector.ScanLines)
	// scanner.Buffer(buf, cap(buf))
	for scanner.Scan() {
		lineCount++
		line := scanner.Bytes()
		fmt.Printf("LINE %d:[%s]\n", lineCount, string(line))
	}
	fmt.Printf("BUFFER: [%s]\n", string(collector.buffer[:collector.bufferSize]))

	fmt.Printf("byte count: %d\n", collector.byteCount)

	// fmt.Printf("ERROR: %s\n", scanner.Err())
	// fmt.Println("")
	// fmt.Println("cap(buf):", cap(buf))
	// fmt.Println("len(buf):", cap(buf))
	// fmt.Println("")
	// fmt.Println("INITIAL BUFFER:", string(buf[:copied]))
}

func newBytesAndLines(bufferSize int) *BytesAndLines {
	return &BytesAndLines{buffer: make([]byte, bufferSize)}
}

// create a data structure to hold the byte size of the file (really, the stream)
// along with the reading of the lines
type BytesAndLines struct {
	byteCount  uint64
	buffer     []byte
	bufferSize int
}

// entrypoint that gathers byte count of what's read so far
// requires buffer to be initialized
func (x *BytesAndLines) ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	a, t, e := bufio.ScanLines(data, atEOF)
	x.byteCount += uint64(a)
	if a > 0 && x.bufferSize+a <= cap(x.buffer) {
		copy(x.buffer[x.bufferSize:], data[0:a])
		x.bufferSize += a
	}
	return a, t, e
}

// // clone ScanLines, but without the `dropCR` so that byte size will be reliable
// func (x *BytesAndLines) scanNewLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
// 	if atEOF && len(data) == 0 {
// 		return 0, nil, nil
// 	}
// 	if i := bytes.IndexByte(data, '\n'); i >= 0 {
// 		// We have a full newline-terminated line.
// 		return i + 1, data[0:i], nil
// 	}
// 	// If we're at EOF, we have a final, non-terminated line. Return it.
// 	if atEOF {
// 		return len(data), data, nil
// 	}
// 	// Request more data.
// 	return 0, nil, nil
// }
