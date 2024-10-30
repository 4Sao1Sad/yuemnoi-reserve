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

func CallPostService(method string, postId uint64, activeStatus bool) error {
	conn, err := connectToPostService()
	if err != nil {
		return err
	}
	defer conn.Close()

	switch method {
	case "BorrowingPost":
		return updateBorrowingPost(conn, postId, activeStatus)
	case "LendingPost":
		return updateLendingPost(conn, postId, activeStatus)
	default:
		return fmt.Errorf("unknown method: %s", method)
	}
}

func updateBorrowingPost(conn *grpc.ClientConn, postId uint64, activeStatus bool) error {
	client := postpb.NewBorrowingPostServiceClient(conn)
	req := &postpb.UpdateBorrowingPostRequest{
		Id:           postId,
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

func updateLendingPost(conn *grpc.ClientConn, postId uint64, activeStatus bool) error {
	client := postpb.NewLendingPostServiceClient(conn)
	req := &postpb.UpdateLendingPostRequest{
		Id:           postId,
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

func GetPost(method string, postId uint64) (interface{}, error) {
	conn, err := connectToPostService()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	switch method {
	case "BorrowingPost":
		return getBorrowingPost(conn, postId)
	case "LendingPost":
		return getLendingPost(conn, postId)
	default:
		return nil, fmt.Errorf("unknown method: %s", method)
	}
}

func getBorrowingPost(conn *grpc.ClientConn, postId uint64) (*postpb.BorrowingPost, error) {
	client := postpb.NewBorrowingPostServiceClient(conn)
	req := &postpb.GetBorrowingPostDetailRequest{Id: postId}
	res, err := client.GetBorrowingPostDetail(context.Background(), req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "borrowing post id not found: %v", err)
	}
	return res, nil
}

func getLendingPost(conn *grpc.ClientConn, postId uint64) (*postpb.LendingPost, error) {
	client := postpb.NewLendingPostServiceClient(conn)
	req := &postpb.GetLendingPostDetailRequest{Id: postId}
	res, err := client.GetLendingPostDetail(context.Background(), req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "lending post id not found: %v", err)
	}
	return res, nil
}

func ValidatePostExists(borrowingPostId *uint64, lendingPostId uint64) error {
	if borrowingPostId != nil {
		if _, err := GetPost("BorrowingPost", *borrowingPostId); err != nil {
			return status.Errorf(codes.NotFound, "Borrowing post with Id %d does not exist", *borrowingPostId)
		}
	}
	if _, err := GetPost("LendingPost", lendingPostId); err != nil {
		return status.Errorf(codes.NotFound, "Lending post with Id %d does not exist", lendingPostId)
	}
	return nil
}

func CheckPostIsReady(borrowingPostId *uint64, lendingPostId uint64) error {
	if borrowingPostId != nil {
		postData, err := GetPost("BorrowingPost", *borrowingPostId)
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
	postData, err := GetPost("LendingPost", lendingPostId)
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
