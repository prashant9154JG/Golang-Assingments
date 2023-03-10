package main

import "io"

type Downloader interface {
	Download(uri string) (r io.Reader, err error)
}

type Archiver interface {
	Archive(names []string, readers ...io.Reader) (outR io.Reader, err error)
}
