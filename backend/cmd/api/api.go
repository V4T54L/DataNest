package main

import (
	"backend/internals/middlewares"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type application struct {
	db *sql.DB
}

func NewApplication(db *sql.DB) *application {
	return &application{db}
}

func (a *application) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", a.healthHandler)

		r.Post("/auth/signup", a.signupHandler)
		r.Post("/auth/login", a.loginHandler)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.AuthMiddleware)

			r.Get("/dashboards", a.getDashboardsHandler)
			r.Post("/dashboards", a.createDashboardHandler)
			r.Put("/dashboards/{dashboardId}", a.updateDashboardByID)

			r.Get("/dashboards/{dashboardId}", a.getDashboardByIDHandler)
			r.Post("/dashboards/{dashboardId}/charts", a.addChartHandler)
			r.Put("/dashboards/{dashboardId}/charts/{chartId}", a.updateChartHandler)
		})
	})

	return r
}
