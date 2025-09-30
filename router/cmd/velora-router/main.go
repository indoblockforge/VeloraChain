package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("ok"))
    })
    log.Println("velora-router running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
