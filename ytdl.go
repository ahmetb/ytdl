package main

import (
	"context"
	"errors"
	"github.com/wader/goutubedl"
	"io"
	"time"
)

type ytdl struct{
	res *goutubedl.Result
}

func (y *ytdl) Info(url string) (Info, error) {
	r, err := goutubedl.New(context.TODO(), url, goutubedl.Options{})
	if err != nil {
		return Info{}, err
	}
	y.res = &r
	return Info{
		ext:      r.Info.Ext,
		id:       r.Info.ID,
		duration: time.Duration(int64(r.Info.Duration)) * time.Second,
	}, nil
}

func (y *ytdl) Download(url string) (io.ReadCloser, error) {
	if y.res == nil {
		return nil, errors.New("call Info() first to get download details")
	}
	return y.res.Download(context.TODO(),"best")
}
