package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strings"
	"os"
	"net"
	"time"
	"fmt"
)

var dataSource string

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func GetData(w http.ResponseWriter, req *http.Request) {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	defer timeTrack(time.Now(), fmt.Sprintf("%s %s from %s", req.Method, req.URL.String(), host))

	resp, err := http.Get(dataSource)
	if err == nil {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			w.Write(bodyBytes)
		} else {
			w.Write([]byte("error"))
		}
	} else {
		log.Println(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide an URL to serve")
	}

	if !strings.HasPrefix(os.Args[1], "http") {
		log.Fatal("Provided URL must start with http or https")
	}

	dataSource = os.Args[1]
	router := mux.NewRouter()
	router.HandleFunc("/", GetData).Methods("GET")
	log.Println("Starting server on port 12345")
	log.Fatal(http.ListenAndServe(":12345", router))
}
