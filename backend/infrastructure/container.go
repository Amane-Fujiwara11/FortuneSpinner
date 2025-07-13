package infrastructure

import (
	"database/sql"

	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/domain/repository"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/infrastructure/mysql"
	infraRepo "github.com/Amane-Fujiwara11/FortuneSpinner/backend/infrastructure/repository"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/interface/handler"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/usecase/gacha"
	"github.com/Amane-Fujiwara11/FortuneSpinner/backend/usecase/point"
)

// Container holds all dependencies
type Container struct {
	// Database
	DB *sql.DB

	// Repositories
	UserRepository  repository.UserRepository
	GachaRepository repository.GachaRepository
	PointRepository repository.PointRepository

	// Use Cases
	GachaUsecase gacha.GachaUsecase
	PointUsecase point.PointUsecase

	// Handlers
	UserHandler  *handler.UserHandler
	GachaHandler *handler.GachaHandler
	PointHandler *handler.PointHandler
}

// NewContainer creates and initializes all dependencies
func NewContainer(dbConfig mysql.Config) (*Container, error) {
	// Initialize database connection
	db, err := mysql.NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	userRepo := infraRepo.NewUserRepository(db)
	gachaRepo := infraRepo.NewGachaRepository(db)
	pointRepo := infraRepo.NewPointRepository(db)

	// Initialize use cases
	gachaUsecase := gacha.NewGachaUsecase(gachaRepo, pointRepo, userRepo)
	pointUsecase := point.NewPointUsecase(pointRepo, userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userRepo)
	gachaHandler := handler.NewGachaHandler(gachaUsecase)
	pointHandler := handler.NewPointHandler(pointUsecase)

	return &Container{
		DB:              db,
		UserRepository:  userRepo,
		GachaRepository: gachaRepo,
		PointRepository: pointRepo,
		GachaUsecase:    gachaUsecase,
		PointUsecase:    pointUsecase,
		UserHandler:     userHandler,
		GachaHandler:    gachaHandler,
		PointHandler:    pointHandler,
	}, nil
}

// Close closes the database connection
func (c *Container) Close() error {
	return c.DB.Close()
}