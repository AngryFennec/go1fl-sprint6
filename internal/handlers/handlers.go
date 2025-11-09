package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Parcing form error", http.StatusBadRequest)
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Converting form error", http.StatusBadRequest)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, "Converting form error", http.StatusBadRequest)
		return
	}

	input := string(content)

	result, err := service.ConvertInput(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().Format("20251212_121212") + ".txt"
	log.Println(fileName)

	err = os.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(result))
	if err != nil {
		log.Println(err.Error())
	}
}
