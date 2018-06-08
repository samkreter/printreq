package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	addrPtr := flag.String("addr", ":8081", "the addr  to server on.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	})

	log.Printf("Serving on addr %s", *addrPtr)

	err := http.ListenAndServe(*addrPtr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
