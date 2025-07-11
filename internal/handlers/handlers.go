package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"yp-go6/internal/service"
)

func GetHtml(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Expected get request", http.StatusBadRequest)
		return
	}

	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(html))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Expected post request", http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(10)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := service.Convert(string(data))

	err = os.Mkdir("uploads", 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	root, err := os.OpenRoot("uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer root.Close()

	filename := filepath.Join("uploads", time.Now().UTC().String())
	filename += filepath.Ext(handler.Filename)

	outfile, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer outfile.Close()

	_, err = fmt.Fprint(outfile, str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(str))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
