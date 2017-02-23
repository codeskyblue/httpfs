package httpfs

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type HTTPFile struct {
	url     string
	size    int64
	name    string
	modTime time.Time
}

var (
	ErrEmptySize       = errors.New("empty size")
	ErrNotAcceptRanges = errors.New("not support Accept-Ranges")
)

// Open return a file like object
func Open(url string) (file *HTTPFile, err error) {
	res, err := http.Head(url)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP Error: " + res.Status)
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return nil, ErrNotAcceptRanges
	}
	if res.ContentLength == 0 {
		return nil, ErrEmptySize
	}
	dateStr := res.Header.Get("Last-Modified")
	if dateStr == "" {
		dateStr = res.Header.Get("Date")
	}
	modTime, _ := time.Parse(http.TimeFormat, dateStr)
	v := &HTTPFile{
		url:     url,
		size:    res.ContentLength,
		name:    filepath.Base(url),
		modTime: modTime,
	}
	return v, nil
}

func (f *HTTPFile) Name() string {
	return f.name
}

func (f *HTTPFile) Size() int64 {
	return f.size
}

func (f *HTTPFile) ModTime() time.Time {
	return f.modTime
}

func (f *HTTPFile) ReadAt(p []byte, off int64) (n int, err error) {
	log.Println("READ", off, len(p))
	req, err := http.NewRequest("GET", f.url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", off, off+int64(len(p))))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	lrd := io.LimitedReader{
		R: resp.Body,
		N: int64(len(p)),
	}
	return lrd.Read(p)
}
