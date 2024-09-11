package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"path"
	"time"
)

// Use generics to get a random file from the embed.FS

//go:embed images/*
var webpFiles embed.FS

func main() {
	http.HandleFunc("/", serveWebP)
	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func serveWebP(w http.ResponseWriter, r *http.Request) {
	fileName := path.Base(r.URL.Path)

	// Serve a list of available files
	files, err := webpFiles.ReadDir("images")
	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}
	log.Println("Serving file:", fileName)
	if fileName == "/" {
		fmt.Fprintf(w, "<html><body><h1>Available WebP Files:</h1><ul>")
		for _, file := range files {
			if !file.IsDir() {
				fmt.Fprintf(w, "<li><a href=\"/%s\">%s</a></li>", file.Name(), file.Name())
			}
		}
		fmt.Fprintf(w, "</ul></body></html>")
		return
	}
	if fileName == "random.webp" {
		index := rand.Intn(len(files))
		fileName = files[index].Name()
	}
	serveWebPFile(w, r, fileName)
}

// Look up a file by name and serve it
func serveWebPFile(w http.ResponseWriter, r *http.Request, fileName string) {

	file, err := webpFiles.Open(path.Join("images", fileName))
	if err != nil {
		if err == fs.ErrNotExist {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Unable to open file", http.StatusInternalServerError)
		}
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/webp")
	http.ServeContent(w, r, fileName, time.Time{}, file.(io.ReadSeeker))
}
