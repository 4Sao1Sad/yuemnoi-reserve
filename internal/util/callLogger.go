package util

import (
	"context"
	"fmt"
	"log"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
	activitypb "github.com/KKhimmoon/yuemnoi-reserve/proto/activity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallActivityLogService(userID uint64, logDetail string) error {
	cfg := config.Load()
	addr := fmt.Sprintf("localhost:%d", cfg.ActivityLogPort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to ActivityLogService: %v", err)
		return err
	}
	defer conn.Close()

	client := activitypb.NewActivityLogServiceClient(conn)

	req := &activitypb.CreateActivityLogRequest{
		LogDetail: logDetail,
		UserId:    userID,
	}

	_, err = client.CreateActivityLog(context.Background(), req)
	if err != nil {
		log.Printf("Error calling CreateActivityLog: %v", err)
		return err
	}
	return nil
}
