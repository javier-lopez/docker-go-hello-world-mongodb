package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Declare a new router
    r := mux.NewRouter()

    // This is where the router is useful, it allows us to declare methods that
    // this path will be valid for
    r.HandleFunc("/", root).Methods("GET")
    r.HandleFunc("/hello",  hello).Methods("GET")
    r.HandleFunc("/hello/", hello).Methods("GET")
    r.HandleFunc("/hello/{name}", hello).Methods("GET")
    r.HandleFunc("/bye",  bye).Methods("GET")

    // We can then pass our router (after declaring all our routes) to this method
    // (where previously, we were leaving the secodn argument as nil)
    http.ListenAndServe(":3001", r)
}

func root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, welcome to your app!</h1>")
    fmt.Fprintf(w, "<h2>Some other interesting endpoints include:</h2>")
    fmt.Fprintf(w, "<code>/hello/&lt;name&gt;</code><br>")
    fmt.Fprintf(w, "<code>/bye</code><br>")
    fmt.Fprintf(w, "<br>")
    fmt.Fprintf(w, "Enjoy!")
}

func hello(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := "anonymous"
    if val, ok := vars["name"]; ok {
        name = val
    }
    fmt.Fprintf(w, "Hello %s!, nice to meet you!", name)
}

func bye(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bye world!")
}
