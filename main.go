package main

import (
	"fmt"
	"github.com/wader/goutubedl"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	maxDuration = time.Second * 60
)

type Info struct {
	ext      string
	id       string
	duration time.Duration
}

type Downloader interface {
	Info(url string) (Info, error)
	Download(url string) (io.ReadCloser, error)
}

func main() {
	addr := ""
	port := os.Getenv("PORT")
	if port == "" {
		addr, port = "localhost", "8080"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr+":"+port, nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	u := req.URL.Query().Get("url")
	if u == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "?url= not specified")
		return
	}
	start := time.Now()
	log.Printf("start url=%s", u)
	defer func() { log.Printf("done url=%s took=%v", u, time.Since(start)) }()

	var d Downloader

	if _, err := exec.LookPath(goutubedl.Path); err == nil {
		log.Printf("initializing real downloader")
		d = new(ytdl)
	} else {
		log.Printf("initializing mock downloader")
		d = new(mockDownloader)
	}

	i, err := d.Info(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to fetch video info: %v", err)
		return
	}
	log.Printf("found video %s format=%s (len=%v)", i.id, i.ext, i.duration)

	if i.duration > maxDuration {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "video too long (%v), max allowed=%v", i.duration, maxDuration)
		return
	}

	s, err := d.Download(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to download video: %v", err)
		return
	}

	fn := fmt.Sprintf("%s.%s", i.id, i.ext)
	w.Header().Set("Content-Disposition", "attachment; filename="+fn)
	w.Header().Set("Content-Type", "application/octet-stream")
	n, err := io.Copy(w, s)
	if err != nil {
		log.Printf("failed to copy resp body after %d bytes: %v", n, err)
		return
	}
	log.Printf("successfully copied %d bytes to response", n)
}
