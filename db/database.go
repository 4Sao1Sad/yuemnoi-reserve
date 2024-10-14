package db

import (
	"fmt"
	"log"
	"net"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/handler"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	reserve "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
	"google.golang.org/grpc"
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
	if err := db.AutoMigrate(&model.BorrowingRequest{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	if err := db.AutoMigrate(&model.LendingRequest{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func ServerInit(cfg *config.Config, db *gorm.DB) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		listen.Close()
	}()

	fmt.Printf("Go gRPC server on port %v!\n", cfg.Port)
	grpcServer := grpc.NewServer()

	//repository
	itemRepository := repository.NewItemRepository(db)
	lendingRequestRepo := repository.NewLendingRequestRepository(db)
	BorrowingRepository := repository.NewBorrowingRepository(db)

	//gRPC handler
	BorrowingServer := handler.NewBorrowingGRPC(BorrowingRepository)
	itemServer := handler.NewItemGRPC(itemRepository)
	lendingRequestServer := handler.NewLendingRequestGRPC(lendingRequestRepo)

	// Register service with the gRPC server
	pb.RegisterItemServiceServer(grpcServer, itemServer)
	pb.RegisterReserveServiceServer(grpcServer, lendingRequestServer)
	reserve.RegisterBorrowingServiceServer(grpcServer, BorrowingServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err)
	}
	return nil
}
