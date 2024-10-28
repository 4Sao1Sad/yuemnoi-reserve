package util

import (
	"context"
	"log"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
	activitypb "github.com/KKhimmoon/yuemnoi-reserve/proto/activity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallActivityLogService(userId uint64, logDetail string) error {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.ActivityLogUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to ActivityLogService: %v", err)
		return err
	}
	defer conn.Close()

	client := activitypb.NewActivityLogServiceClient(conn)

	req := &activitypb.CreateActivityLogRequest{
		LogDetail: logDetail,
		UserId:    userId,
	}

	_, err = client.CreateActivityLog(context.Background(), req)
	if err != nil {
		log.Printf("Error calling CreateActivityLog: %v", err)
		return err
	}
	return nil
}
