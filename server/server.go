package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	strs := strings.Split(r.URL.Path, "/")
	if len(strs) <= 1 {
		http.NotFound(w, r)
		return
	}
	log.Printf("url: %s", r.URL.Path)
	// r.URL.Path == "/" -> ["",""]
	// r.URL.Path == "/js/wasm_exec.js" -> ["","js","wasm_exec.js"]
	switch strs[1] {
	case "":
		http.ServeFile(w, r, "index.html")
		return
	case "wasm":
		retWasmFile(w, r)
		return
	case "js":
		http.ServeFile(w, r, r.URL.Path[1:])
		return
	}
	http.NotFound(w, r)
}

func retWasmFile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("main.wasm")
	if err != nil {
		log.Printf("read file error: %s", err.Error())
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=main.wasm")
	w.Header().Set("Content-Type", "application/wasm")
	w.Header().Set("Content-Length", string(len(data)))
	w.Write(data)
}
