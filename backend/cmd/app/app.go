package app

import (
	"idea-repository-backend/internal/models"
	"net/http"

	"gorm.io/driver/postgres"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"idea-repository-backend/config"
	"idea-repository-backend/internal/database"
	"idea-repository-backend/internal/handlers"
	"idea-repository-backend/internal/middleware"
	"idea-repository-backend/internal/services"
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

	app.routes()

	return app, nil
}

func NewAppWithDependencies(cfg *config.Config, db *database.Database) (*App, error) {
	// Initialize services
	svc := services.New(db)

	// Initialize handlers
	h := handlers.New(svc)

	// Set up router
	router := chi.NewRouter()

	app := &App{
		Config:   cfg,
		DB:       db.GormDB,
		Router:   router,
		Handlers: h,
	}

	app.routes()

	return app, nil
}

func (a *App) routes() {
	a.Router.Use(middleware.EnableCors)
	a.Router.Post("/ideas", a.Handlers.CreateIdeaHandler)
	// Add more routes here
}

func (a *App) Serve() error {
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
