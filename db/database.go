package db

import (
	"fmt"
	"github.com/KKhimmoon/yuemnoi-reserve/config"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
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
	if err := db.AutoMigrate(&model.Item{}); err != nil {
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

	//gRPC handler
	itemServer := handler.NewItemGRPC(itemRepository)

	// Register service with the gRPC server
	pb.RegisterItemServiceServer(grpcServer, itemServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err)
	}
	return nil
}
