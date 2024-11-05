package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
	handler "github.com/KKhimmoon/yuemnoi-reserve/internal/handler"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreSQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.Db.Host, cfg.Db.Username, cfg.Db.Password, cfg.Db.Database, cfg.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	Migration(db)
	return db
}

func Migration(db *gorm.DB) {
	if err := db.AutoMigrate(&model.BorrowingRequest{}, &model.LendingRequest{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func ServerInit(cfg *config.Config, db *gorm.DB) error {
	// Initialize Fiber app
	app := fiber.New()
	app.Use(requestid.New())

	// Initialize repositories
	borrowingRequestRepository := repository.NewBorrowingRequestRepository(db)
	lendingRequestRepository := repository.NewLendingRequestRepository(db)

	// Initialize handlers
	borrowingRequestRestHandler := handler.NewBorrowingRequestRestHandler(borrowingRequestRepository)
	lendingRequestRestHandler := handler.NewLendingRequestRestHandler(lendingRequestRepository)
	requestRestHandler := handler.NewRequestRestHandler(borrowingRequestRepository, lendingRequestRepository)

	// Register routes
	r := route.NewHandler(borrowingRequestRestHandler, lendingRequestRestHandler, requestRestHandler)
	r.RegisterRouter(app, cfg)

	// Start the server with error handling
	port := strconv.Itoa(int(cfg.Port))
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
		return err
	}

	return nil
}
