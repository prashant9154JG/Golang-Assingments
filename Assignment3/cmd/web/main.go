package main

import (
	api "assignment3/cmd/api"
	"fmt"
	"log"
)

// NewDownloader1 initialise the downloader for file system
func NewDownloader1() *api.FSUrl {
	x := new(api.FSUrl)
	return x
}

// NewDownloader2 initialise the downloader for web
func NewDownloader2() *api.WebUrl {
	x := new(api.WebUrl)
	return x
}

// New initialise Zip
func New() *api.Zip {
	x := new(api.Zip)
	return x
}

const (
	filePath = "./public/mytext.txt"
	urlPath  = "https://static.remove.bg/sample-gallery/graphics/bird-thumbnail.jpg"
)

func main() {
	downloader1 := NewDownloader1()
	downloader2 := NewDownloader2()
	zipper := New()

	// calling download function for file system path
	r1, err := downloader1.Download(filePath)

	if err != nil {
		log.Println(fmt.Errorf("cannot download file from file system: %w", err))
	}

	//calling download function for web url
	r2, err := downloader2.Download(urlPath)
	if err != nil {
		log.Println(fmt.Errorf("cannot download file from web: %w", err))
	}

	// archiving two downloaded file into archive.zip
	_, err = zipper.Archive([]string{"mytext.txt", "logo.jpg"}, r1, r2)

	if err != nil {
		log.Println(fmt.Errorf("cannot archive file: %w", err))
	}
}
