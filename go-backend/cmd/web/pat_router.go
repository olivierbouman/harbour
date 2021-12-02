package main

import (
	"go-backend/pkg/config"
	"go-backend/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

// Low level but very functional router
func pat_router(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
