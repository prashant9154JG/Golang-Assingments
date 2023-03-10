package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type WebUrl struct{}

// Download downlaod the file with given url through web
func (d *WebUrl) Download(uri string) (r io.Reader, err error) {
	fmt.Println("Downloading file from Web....")
	// Get the data
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create("public/logo.jpg")
	if err != nil {
		return nil, err
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	_, pw := io.Pipe()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		buf := make([]byte, 1024)
		defer wg.Done()
		for {
			chunk, err := out.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if chunk == 0 {
				break
			}

			if _, err := pw.Write(buf[:chunk]); err != nil {
				return
			}
		}
		pw.Close()
	}()

	wg.Wait()

	out.Seek(0, io.SeekStart)

	return out, nil
}
