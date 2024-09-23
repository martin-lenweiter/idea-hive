package app

import (
	"ideahive/backend/internal/models"
	"log"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"ideahive/backend/config"
	"ideahive/backend/internal/database"
	"ideahive/backend/internal/handlers"
	"ideahive/backend/internal/middleware"
	"ideahive/backend/internal/services"
)

type App struct {
	Config   *config.Config
	DB       *gorm.DB
	Router   *chi.Mux
	Handlers *handlers.Handlers
}

func NewApp() (*App, error) {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// Initialize database
	db, err := initDatabase(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// Initialize services
	svc, err := initServices(db)
	if err != nil {
		return nil, err
	}

	// Initialize handlers
	h := handlers.New(svc)

	// Set up router
	router := chi.NewRouter()

	app := &App{
		Config:   cfg,
		DB:       db,
		Router:   router,
		Handlers: h,
	}

	app.Routes()

	return app, nil
}

func (a *App) Routes() {
	// Middleware
	a.Router.Use(middleware.CorsMiddleware())
	a.Router.Use(middleware.ForceSSL)

	// API Routes
	a.Router.Route("/api", func(r chi.Router) {
		routeApi(r, a.Handlers)
	})

	// Serve static files
	a.Router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("/root/public/static"))))

	// Serve index.html for all other routes
	a.Router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/root/public/index.html")
	})
}

func routeApi(r chi.Router, handlers *handlers.Handlers) {
	r.Post("/ideas", handlers.CreateIdeaHandler)
	// Add more API Routes here
}

func (a *App) Serve() error {
	log.Printf("Starting server on %s", a.Config.ServerAddress)
	return http.ListenAndServe(a.Config.ServerAddress, a.Router)
}

func initDatabase(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate your models
	err = db.AutoMigrate(&models.Idea{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initServices(db *gorm.DB) (*services.Services, error) {
	// Create a database.Database instance from the gorm.DB
	dbInstance := &database.Database{GormDB: db}
	// Initialize services with Database
	return services.New(dbInstance), nil
}
