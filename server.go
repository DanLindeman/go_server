package main

import (
    "fmt"
    "net/http"
    "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
    var to_validate string = r.URL.Path[1:]
    fmt.Fprintf(w, "Validating: %s \n\n", to_validate)
    validate(w, to_validate)
}

func validate(w http.ResponseWriter, to_validate string) {
    input := strings.Split(to_validate, "/")
    var kind, value = input[0], input[1]
    if strings.Contains("first", kind) {
        validate_first_kind(w, value)
    } else {
        validate_second_kind(w, value)
    }
}

func validate_first_kind(w http.ResponseWriter, value string){
    if strings.Contains("cats", value) {
        fmt.Fprintf(w, "Super Valid!")
    } else {
        fmt.Fprintf(w, "Not Valid!")
    }
}

func validate_second_kind(w http.ResponseWriter, value string){
    if strings.Contains("dogs", value) {
        fmt.Fprintf(w, "Super Valid!")
    } else {
        fmt.Fprintf(w, "Not Valid!")
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}