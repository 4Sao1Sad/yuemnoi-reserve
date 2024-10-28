package util

import (
	"context"
	"fmt"
	"log"

	"github.com/KKhimmoon/yuemnoi-reserve/config"
	postpb "github.com/KKhimmoon/yuemnoi-reserve/proto/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func connectToPostService() (*grpc.ClientConn, error) {
	cfg := config.Load()
	conn, err := grpc.NewClient(cfg.PostUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to PostService: %v", err)
		return nil, err
	}
	return conn, nil
}

func CallPostService(method string, postID uint64, activeStatus bool) error {
	conn, err := connectToPostService()
	if err != nil {
		return err
	}
	defer conn.Close()

	switch method {
	case "BorrowingPost":
		return updateBorrowingPost(conn, postID, activeStatus)
	case "LendingPost":
		return updateLendingPost(conn, postID, activeStatus)
	default:
		return fmt.Errorf("unknown method: %s", method)
	}
}

func updateBorrowingPost(conn *grpc.ClientConn, postID uint64, activeStatus bool) error {
	client := postpb.NewBorrowingPostServiceClient(conn)
	req := &postpb.UpdateBorrowingPostRequest{
		Id:           postID,
		ActiveStatus: activeStatus,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"active_status"},
		},
	}
	_, err := client.UpdateBorrowingPost(context.Background(), req)
	if err != nil {
		return status.Errorf(codes.Internal, "error updating borrowing post: %v", err)
	}
	return nil
}

func updateLendingPost(conn *grpc.ClientConn, postID uint64, activeStatus bool) error {
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
		return status.Errorf(codes.Internal, "error updating lending post: %v", err)
	}
	return nil
}

func GetPost(method string, postID uint64) (interface{}, error) {
	conn, err := connectToPostService()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	switch method {
	case "BorrowingPost":
		return getBorrowingPost(conn, postID)
	case "LendingPost":
		return getLendingPost(conn, postID)
	default:
		return nil, fmt.Errorf("unknown method: %s", method)
	}
}

func getBorrowingPost(conn *grpc.ClientConn, postID uint64) (*postpb.BorrowingPost, error) {
	client := postpb.NewBorrowingPostServiceClient(conn)
	req := &postpb.GetBorrowingPostDetailRequest{Id: postID}
	res, err := client.GetBorrowingPostDetail(context.Background(), req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "borrowing post id not found: %v", err)
	}
	return res, nil
}

func getLendingPost(conn *grpc.ClientConn, postID uint64) (*postpb.LendingPost, error) {
	client := postpb.NewLendingPostServiceClient(conn)
	req := &postpb.GetLendingPostDetailRequest{Id: postID}
	res, err := client.GetLendingPostDetail(context.Background(), req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "lending post id not found: %v", err)
	}
	return res, nil
}

func ValidatePostExists(borrowingPostID *uint64, lendingPostID uint64) error {
	if borrowingPostID != nil {
		if _, err := GetPost("BorrowingPost", *borrowingPostID); err != nil {
			return status.Errorf(codes.NotFound, "Borrowing post with ID %d does not exist", *borrowingPostID)
		}
	}
	if _, err := GetPost("LendingPost", lendingPostID); err != nil {
		return status.Errorf(codes.NotFound, "Lending post with ID %d does not exist", lendingPostID)
	}
	return nil
}

func CheckPostIsReady(borrowingPostID *uint64, lendingPostID uint64) error {
	if borrowingPostID != nil {
		postData, err := GetPost("BorrowingPost", *borrowingPostID)
		if err != nil {
			return status.Errorf(codes.NotFound, "Borrowing post not found: %v", err)
		}
		borrowingPost, ok := postData.(*postpb.BorrowingPost)
		if !ok {
			return status.Errorf(codes.Internal, "Unexpected data format for BorrowingPost")
		}
		if !borrowingPost.ActiveStatus {
			return status.Errorf(codes.FailedPrecondition, "Borrowing post is not active")
		}
	}
	postData, err := GetPost("LendingPost", lendingPostID)
	if err != nil {
		return status.Errorf(codes.NotFound, "Lending post not found: %v", err)
	}
	lendingPost, ok := postData.(*postpb.LendingPost)
	if !ok {
		return status.Errorf(codes.Internal, "Unexpected data format for LendingPost")
	}
	if !lendingPost.ActiveStatus {
		return status.Errorf(codes.FailedPrecondition, "Lending post is not ready")
	}
	return nil
}
