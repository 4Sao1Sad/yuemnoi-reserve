package util

import (
	"context"
	"fmt"
	"log"

	"github.com/KKhimmoon/yuemnoi-reserve/config"

	postpb "github.com/KKhimmoon/yuemnoi-reserve/proto/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func CallPostService(method string, postID uint64, activeStatus bool) error {
	cfg := config.Load()
	addr := fmt.Sprintf("localhost:%d", cfg.PostPort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to PostService: %v", err)
		return err
	}
	defer conn.Close()
	switch method {
	case "BorrowingPost":
		req := &postpb.UpdateBorrowingPostRequest{
			Id:           postID,
			ActiveStatus: activeStatus,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"active_status"},
			},
		}
		client := postpb.NewBorrowingPostServiceClient(conn)
		_, err := client.UpdateBorrowingPost(context.Background(), req)
		if err != nil {
			log.Printf("Error calling update borrowing post: %v", err)
			return err
		}

	case "LendingPost":
		client := postpb.NewLendingPostServiceClient(conn)
		req := &postpb.UpdateLendingPostRequest{
			Id:           postID,
			ActiveStatus: activeStatus,
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{"active_status"},
			},
		}
		_, err := client.UpdateLendingPost(context.Background(), req)
		if err != nil {
			log.Printf("Error calling update lending post: %v", err)
			return err
		}
	default:
		log.Printf("Unknown method: %s", method)
		return nil
	}
	return nil
}
