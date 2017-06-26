// server start: go run chat_room.go
// access to 127.0.0.1:3000
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := flag.String("b", "127.0.0.1", "Binds to the specified IP.")
	port := flag.String("p", "3000", "Runs on the specified port.")
	flag.Parse()

	appRoot := fmt.Sprintf("%v:%v", *host, *port)

	http.Handle("/", &templateHandler{filename: "welcome.html"})

	log.Printf("application starting on %v\n", appRoot)

	if err := http.ListenAndServe(appRoot, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
