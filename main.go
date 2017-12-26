package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const RepositoryUrl = "https://github.com/zunda/status-code-app/"

func writeUsage(w http.ResponseWriter) {
	fmt.Fprintf(w, `
<html><title>status-code-app</title><body><h1>status-code-app</h1>
<p>Hi, I'll respond to requests with specified status code or trigger errors. Try something like</p>
<ul>
<li><a href="/418">/418</a> - I'm a teapot
<li><a href="/H12">/H12</a> - Request timeout
<li><a href="/H13">/H13</a> - Connection closed without response
</ul>
<p>Source code is available at <a href="%s">%s</a></p>
</body></html>
`, RepositoryUrl, RepositoryUrl)
}

func h12Server(w http.ResponseWriter, r *http.Request) {
	log.Printf("Waiting for 35 sec to trigger an H12 - Request timeout")
	time.Sleep(35 * time.Second)
	fmt.Fprintf(w, "Hi. did you see an H12 from Heroku router in the app logs?")
	log.Printf("Done")
}

func h13Server(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if ok {
		conn, _, err := hj.Hijack()
		if err == nil {
			log.Printf("Closing writer to trigger an H13 - Connection closed without response")
			conn.Close()
			return
		}
	}

	log.Printf("Couldn't close the connection to trigegr an H13")
	fmt.Fprintf(w, "Hi. I'm unavailable to trigger an H13.")
}

func statusCodeServer(w http.ResponseWriter, r *http.Request) {
	s := strings.SplitN(r.URL.Path, "/", 3)
	c, err := strconv.Atoi(s[1])
	if err != nil {
		writeUsage(w)
		return
	}

	t := http.StatusText(c)
	if t == "" {
		t = "unknown status code"
	}
	log.Printf("Responding with the status code %d - %s", c, t)
	w.WriteHeader(c)
	fmt.Fprintf(w, "Hi. I'm responding with the status code %d - %s.\n", c, t)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	h := http.NewServeMux()
	h.HandleFunc("/favicon.ico", http.NotFound)
	h.HandleFunc("/H12", h12Server)
	h.HandleFunc("/H13", h13Server)
	h.HandleFunc("/", statusCodeServer)

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
