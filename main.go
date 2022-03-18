package main

import (
	"github.com/gorilla/handlers"
	"log"
	"strings"

	"fmt"
	"net/http"
	"os"
)


func ShowUsage()  {
	myslice := strings.Split(os.Args[0], "/")
	fmt.Println("")
	fmt.Println("#########################################################")
	fmt.Println("  Usage:")
	fmt.Println("\t ", myslice[len(myslice)-1], "ServePort", "ServeFolder")
	fmt.Println("e.g.")
	fmt.Println("  ./oohttp")
	fmt.Println("  ./oohttp 8081")
	fmt.Println("  ./oohttp 12345 testfolder")
	fmt.Println("########################################################")
}


func main() {

	wwwFolder := "."
	wwwPort := "5678"

	argc := len(os.Args) -1

	if argc > 0 {
		if argc > 2 {
			ShowUsage()
			os.Exit(1)
		} else {
			wwwPort = os.Args[1]

			if argc == 2 {
				wwwFolder = os.Args[2]
			}
		}
	}

	// Only listen on localhost by default for security reason
	listenAddr := "0.0.0.0:" + wwwPort

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(wwwFolder))))

	log.Println("Hi master, I am listening on " , listenAddr)
	log.Printf("Open http://%s and enjoy yourself!", listenAddr)

	_ = http.ListenAndServe(listenAddr, nil)

}
