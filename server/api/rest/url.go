package rest

import (
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/error2215/url_shortener/server/models"
	"github.com/error2215/url_shortener/server/models/url"
)

func (s *Server) CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		log.Error(err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return
	}
	urlStr := r.PostForm.Get("url")
	if urlStr == "" {
		err := errors.New("Missing required params")
		log.Error(err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return

	}
	shortened, err := url.FindShortenedUrl(r.Context(), urlStr)
	if err != nil {
		log.Error(err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return

	}
	if shortened == "" {
		shortened, err = url.CreateShortenedUrl(r.Context(), urlStr)
		if err != nil {
			log.Error(err)
			_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
			return
		}
	}
	_, _ = w.Write(models.Response(0, "", []byte(`{"alias": "`+shortened+`"}`)).ToString())
}

func (s *Server) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	shortened := strings.Split(r.URL.String(), "/")[2]
	realUrl, err := url.FindRealUrl(r.Context(), shortened)
	if err != nil {
		log.Error(err)
		_, _ = w.Write(models.Response(1, err.Error(), nil).ToString())
		return

	}
	if realUrl == "" {
		w.WriteHeader(404)
		_, _ = w.Write(models.Response(1, "not found", nil).ToString())
	}
	w.WriteHeader(302)
	http.Redirect(w, r, realUrl, 302)
}
