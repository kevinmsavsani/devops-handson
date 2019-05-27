package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    mt.Fprintf(w, "Hello DevOps Version 3")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8081", nil)
}
