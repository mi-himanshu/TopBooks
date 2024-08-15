package main

import (
    "log"
    "net"
    "net/http"
    "encoding/json"
    "topbooks/models"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from BE!"))
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
    log.Println("HandleBooks - Request received.")
    json.NewEncoder(w).Encode(models.GetAllBooks())
}


func main() {
    // set a HTTP request handle function for path /greeting and registrate it
    http.HandleFunc("/books", handleBooks)
    http.HandleFunc("/", handleHome)

    // print out the server is going to start and show the time
    log.Println("Starting server....")

    // create server at localhost:8080 and using tcp as the network
    listener, err := net.Listen("tcp", ":8080")

    // if recieve error, record it and exit the program
    if err != nil {
        log.Fatal(err)
    }

    // setup HTTP connection for the listener of the server
    http.Serve(listener, nil)
}