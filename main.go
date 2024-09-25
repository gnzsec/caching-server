package main

import (
    "flag"
    "fmt" // Agregar esta línea
    "io/ioutil"
    "log"
    "net/http"
    "sync"
)

var (
    cache      = make(map[string][]byte)
    mu         sync.Mutex
    originURL  string
)

func main() {
    port := flag.Int("port", 8080, "Port to run the caching proxy server")
    origin := flag.String("origin", "http://dummyjson.com", "Origin server URL")
    clearCache := flag.Bool("clear-cache", false, "Clear the cache")
    flag.Parse()

    if *clearCache {
        mu.Lock()
        cache = make(map[string][]byte) // Clear the cache
        mu.Unlock()
        log.Println("Cache cleared")
        return
    }

    originURL = *origin
    http.HandleFunc("/", proxyHandler)
    log.Printf("Servidor proxy en ejecución en el puerto %d\n", *port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    cachedResponse, found := cache[r.URL.String()]
    mu.Unlock()

    if found {
        log.Println("Respuesta desde la caché")
        w.Header().Set("X-Cache", "HIT")
        w.Write(cachedResponse)
        return
    }

    resp, err := http.Get(originURL + r.URL.String())
    if err != nil {
        http.Error(w, "Error al obtener la respuesta", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Error al leer la respuesta", http.StatusInternalServerError)
        return
    }

    mu.Lock()
    cache[r.URL.String()] = body
    mu.Unlock()

    w.Header().Set("X-Cache", "MISS")
    w.Write(body)
}
