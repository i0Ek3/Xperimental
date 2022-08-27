package main

import (
    "fmt"
    "log"
    "net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `
<html>
    <head><title>Go Web</title></head>
    <body>
        <form method="post" action="/body">
            <label for="usename"> Username: </label>
            <input type="text" id="username" name="username">
            <label for="email"> Email: </label>
            <input type="text" id="email" name="email">
            <button type="submit">Go</button>
        </form>
    </body>
</html>`)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, r.URL.RawQuery)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Fprintln(w, r.Form)
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", loginHandler)
    mux.HandleFunc("/query", queryHandler)
    mux.HandleFunc("/form", formHandler)

    server := &http.Server {
        Addr:    ":8080",
        Handler: mux,
    }

    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
