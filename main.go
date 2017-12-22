package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func statusCodeServer(w http.ResponseWriter, r *http.Request) {
	s := strings.SplitN(r.URL.Path, "/", 3)
	c, err := strconv.Atoi(s[1])
	if err != nil {
		fmt.Fprintf(w, "Specify desired status code after /")
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
	h.HandleFunc("/", statusCodeServer)

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
