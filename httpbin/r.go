package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

var (
	flagRootDir     = flag.String("dir", ".", "root dir")
	flagHttpAddr    = flag.String("http", "localhost:8081", "http service address")
	flagNoCache     = flag.Bool("no-cache", true, "disable browser cache")
	flagOpenBrowser = flag.Bool("openbroswer", true, "open browser automatically")
)

func main() {
	flag.Parse()

	host, port, err := net.SplitHostPort(*flagHttpAddr)
	if err != nil {
		log.Fatal(err)
	}

	if host == "" {
		host = getLocalIp()
	}
	httpAddr := fmt.Sprintf("%s:%s", host, port)
	url := fmt.Sprintf("http://%s", httpAddr)

	go func() {
		if waitServer(url) && *flagOpenBrowser && startBrowser(url) {
			log.Printf("A browser window should open. If not please visit %s", url)
		} else {
			log.Printf("Please open your web browser and visit %s", url)
		}
		log.Printf("Hit CTRL-C to stop the server\n")
	}()

	mainHandler := http.FileServer(http.Dir(*flagRootDir))
	if *flagNoCache {
		mainHandler = NoCache(mainHandler)
	}

	log.Fatal(http.ListenAndServe(httpAddr, mainHandler))
}

var epoch = time.Unix(0, 0).Format(time.RFC1123)
var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

func NoCache(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for _, v := range etagHeaders {
			if r.Header.Get(v) != "" {
				r.Header.Del(v)
			}
		}
		for k, v := range noCacheHeaders {
			r.Header.Set(k, v)
		}
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			_ = resp.Body.Close()
			return true
		}
		time.Sleep(1000 * time.Microsecond)
		tries = tries - 1
	}

	return false
}

func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func getLocalIp() string {
	adders, err := net.InterfaceAddrs()
	d := "127.0.0.1"
	if err != nil {
		return d
	}
	for _, address := range adders {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return d
}
