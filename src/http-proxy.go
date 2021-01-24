package main
import (
    "io"
    "log"
    "net/http"

)

func handleHTTP(w http.ResponseWriter, req *http.Request) {
    resp, err := http.DefaultTransport.RoundTrip(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusServiceUnavailable)
        return
    }
    defer resp.Body.Close()
    copyHeader(w.Header(), resp.Header)
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}
func copyHeader(dst, src http.Header) {
    for k, vv := range src {
        for _, v := range vv {
            dst.Add(k, v)
        }
    }
}

func main() {

    server := &http.Server{
        Addr: "0.0.0.0:9977",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            handleHTTP(w, r)
        }),
    }

    log.Fatal(server.ListenAndServe())
}