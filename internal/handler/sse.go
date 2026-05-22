package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"
	"sync"
)

var (
	sseBufferSize = 64 * 1024 // 64KB
	sseDataPrefix = "data:"

	pool = sync.Pool{
		New: func() any {
			buf := make([]byte, 0, 4096)
			return &buf
		},
	}
)

func getBuffer() *[]byte {
	return pool.Get().(*[]byte)
}

func putBuffer(buf *[]byte) {
	*buf = (*buf)[:0]
}

func sseHandler(body io.ReadCloser) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		defer body.Close()

		scanner := bufio.NewScanner(body)
		buf := getBuffer()
		defer putBuffer(buf)

		scanner.Buffer(*buf, sseBufferSize)
		for scanner.Scan() {
			lineBytes := scanner.Bytes()
			prefix := []byte(sseDataPrefix)
			if !bytes.HasPrefix(lineBytes, prefix) {
				continue
			}
			data := lineBytes[len(prefix):]
			if len(data) > 0 && data[0] == ' ' {
				data = data[1:]
			}
			if !yield(data, nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield(nil, fmt.Errorf("sseHandler scanning error: %v", err))
		}
	}
}
