package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerRoot(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	page, err := os.ReadFile("../index.html")

	if err != nil {
		http.Error(w, "Error while getting page", http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(page)

	if err != nil {
		http.Error(w, "Error while sending page", http.StatusInternalServerError)
	}
}

func HandlerUpload(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	err := req.ParseMultipartForm(10 << 20) // 10 Mb

	if err != nil {
		http.Error(w, "Parse Multipart Forme error", http.StatusInternalServerError)
	}

	file, fileHeader, err := req.FormFile("myFile")

	if err != nil {
		errorMessage := fmt.Sprintf("Loading file error: %s", err.Error())
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}

	defer file.Close()

	text := make([]byte, fileHeader.Size)
	_, err = file.Read(text)

	if err != nil {
		errorMessage := fmt.Sprintf("Reading file error: %s", err.Error())
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}

	translatedText := service.MorseTranslator(string(text))

	fileName := time.Now().UTC().Format("../02.01.2006 15.04.05.txt")

	err = os.WriteFile(fileName, []byte(translatedText), 0755)

	if err != nil {
		errorMessage := fmt.Sprintf("Writing file error: %s", err.Error())
		http.Error(w, errorMessage, http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, translatedText)
}
