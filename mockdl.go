package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"time"
)

type mockDownloader struct{}

func (f mockDownloader) Info(_ string) (Info, error) {
	return Info{
		ext:      "mp4",
		id:       "some_id",
		duration: time.Second * 100,
	}, nil
}

func (f mockDownloader) Download(_ string) (io.ReadCloser, error) {
	r := rand.Reader
	time.Sleep(time.Millisecond*200)
	return ioutil.NopCloser(io.LimitReader(r, 10*1024)), nil
}
