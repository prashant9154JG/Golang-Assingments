package api

import (
	"fmt"
	"io"
	"os"
)

type FSUrl struct{}

// Download download the file with given path of file in file system
func (d *FSUrl) Download(uri string) (r io.Reader, err error) {
	fmt.Println("Downloading file from File system....")
	r, w := io.Pipe()

	if err != nil {
		return nil, err
	}

	go func() {
		f, err := os.Open(uri)
		if err != nil {
			return

		}

		buf := make([]byte, 1024)
		for {
			chunk, err := f.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if chunk == 0 {
				break
			}

			if _, err := w.Write(buf[:chunk]); err != nil {
				return
			}

		}

		w.Close()
		f.Close()
	}()

	return r, nil
}
