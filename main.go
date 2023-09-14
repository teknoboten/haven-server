package main

import (
	"fmt"
	"net/http"
)

func sup(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "gimme some food!\n")
}

func main() {
	http.HandleFunc("/sup", sup)
	http.Handle("/", http.FileServer(http.Dir("./haven-test-minimal/src/")))		
	http.ListenAndServe(":8080", nil)
}