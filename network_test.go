package goist

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httptest"
    "time"
)

func ExampleGet() {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, Get")
    }))
    defer ts.Close()

    res, err := Get(ts.URL, nil, nil, 10*time.Second)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", string(res))

    // Output: Hello, Get
}
