package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	listenAddr := flag.String("http.addr", ":3000", "http listen address")
	flag.Parse()

	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func handler(numbers []int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		for k, v := range values {
			fmt.Print(k, v)
			for i := 0; i < len(v); i++ {

			}
		}
	}
}
