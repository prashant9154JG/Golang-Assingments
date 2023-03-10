package api

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

type Zip struct{}

// Archive archive the given multiple files into a single zip file
func (z *Zip) Archive(names []string, readers ...io.Reader) (outR io.Reader, err error) {
	fmt.Println("creating zip archive...")

	res, err := os.Create("archive.zip")

	zipWriter := zip.NewWriter(res)
	defer zipWriter.Close()
	defer zipWriter.Flush()

	for i := 0; i < len(names); i++ {
		fmt.Printf("opening %dth file...\n", i+1)

		fmt.Printf("writing %dth file to archive...\n", i+1)
		w, err := zipWriter.Create(names[i])
		if err != nil {
			return nil, fmt.Errorf("could not add file to zip file : %w", err)
		}
		if _, err := io.Copy(w, readers[i]); err != nil {
			return nil, fmt.Errorf("could copy reader to zip writer : %w", err)
		}
	}

	fmt.Println("closing zip archive...")
	// res.Seek(0, io.SeekStart)

	return res, nil
}
