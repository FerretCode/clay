package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/forward", func(w http.ResponseWriter, r *http.Request) {
		target := r.Header.Get("target")

		url, err := url.Parse(target)

		if err != nil {
			http.Error(w, "The target head did not contain a valid URL!", http.StatusBadRequest)

			return
		}

		proxy := httputil.NewSingleHostReverseProxy(url)

		proxy.Director = func(request *http.Request) {
			request.Host = url.Host
			request.URL.Scheme = url.Scheme
			request.URL.Host = url.Host
			request.URL.Path = url.Path
		}

		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":3000", r)
}
