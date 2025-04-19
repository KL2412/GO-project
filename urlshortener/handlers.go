package urlshortener

import (
	"html/template"
	"net/http"
	"github.com/KL2412/GO-project/config"
)

type PageData struct {
	Name         string
	ShortenedURL string
	OriginalURL  string
}

type Handlers struct {
	shortener *URLShortener
	tmpl      *template.Template
	config    *config.Config
}

func NewHandlers(shortener *URLShortener, cfg *config.Config) *Handlers {
	return &Handlers{
		shortener: shortener,
		tmpl:      template.Must(template.New("").ParseGlob("./templates/*")),
		config:    cfg,
	}
}

func (h *Handlers) HandleIndex(w http.ResponseWriter, r *http.Request) {
	h.tmpl.ExecuteTemplate(w, "index.html", PageData{
		Name: "URL Shortener",
	})
}

func (h *Handlers) HandleShortenWithTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// TODO: Check if the URL is already shortened

	// TODO: Check if the URL is valid

	// TODO: Check if the URL already exists in the database (RN all data is stored in memory)

	shortKey := generateShortKey()
	h.shortener.urls[shortKey] = originalURL

	shortenedURL := h.config.BaseURL + "/short/" + shortKey

	h.tmpl.ExecuteTemplate(w, "index.html", PageData{
		Name:         "URL Shortener",
		ShortenedURL: shortenedURL,
		OriginalURL:  originalURL,
	})
} 