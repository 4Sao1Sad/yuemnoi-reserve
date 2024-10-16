package handler

import (
	"context"

	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/util"
	pb "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BorrowingGRPC struct {
	pb.UnimplementedBorrowingServiceServer
	repository repository.BorrowingRepository
}

func NewBorrowingGRPC(repo repository.BorrowingRepository) *BorrowingGRPC {
	return &BorrowingGRPC{
		repository: repo,
	}
}

func (h *BorrowingGRPC) CreateRequestFromBorrowingPost(ctx context.Context, input *pb.CreateRequestFromBorrowingPostInput) (*pb.CreateRequestFromBorrowingPostResponse, error) {
	data := model.BorrowingRequest{
		LendingUserID:   uint(input.LendingUserId),
		BorrowingUserID: uint(input.BorrowingUserId),
		LendingPostID:   uint(input.LendingPostId),
		BorrowingPostID: uint(input.BorrowingPostId),
		Status:          model.Pending,
		ActiveStatus:    true,
	}
	_, err := h.repository.CreateRequestFromBorrowingPost(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create borrowing request: %v", err)
	}
	response := pb.CreateRequestFromBorrowingPostResponse{
		Message: "created successfully",
	}

	return &response, nil
}

func (h *BorrowingGRPC) GetBorrowingRequestById(ctx context.Context, input *pb.GetBorrowingRequestInput) (*pb.BorrowingRequest, error) {
	res, err := h.repository.GetRequestById(uint(input.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Borrowing request not found: %v", err)
	}
	response := pb.BorrowingRequest{
		Id:              uint64(res.ID),
		LendingUserId:   uint64(res.LendingUserID),
		BorrowingUserId: uint64(res.BorrowingUserID),
		LendingPostId:   uint64(res.LendingPostID),
		BorrowingPostId: uint64(res.BorrowingPostID),
		Status:          util.MapModelToProtoStatus(res.Status),
		ActiveStatus:    res.ActiveStatus,
	}

	return &response, nil
}

func (h *BorrowingGRPC) ConfirmBorrowingRequest(ctx context.Context, input *pb.ConfirmBorrowingRequestInput) (*pb.BorrowingRequest, error) {
	res, err := h.repository.GetRequestById(uint(input.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Borrowing request not found: %v", err)
	}

	res, err = h.repository.ConfirmBorrowingRequest(res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to confirm borrowing request: %v", err)
	}
	response := pb.BorrowingRequest{
		Id:              uint64(res.ID),
		LendingUserId:   uint64(res.LendingUserID),
		BorrowingUserId: uint64(res.BorrowingUserID),
		LendingPostId:   uint64(res.LendingPostID),
		BorrowingPostId: uint64(res.BorrowingPostID),
		Status:          util.MapModelToProtoStatus(res.Status),
		ActiveStatus:    res.ActiveStatus,
	}

	return &response, nil
}

func (h *BorrowingGRPC) RejectBorrowingRequest(ctx context.Context, input *pb.RejectBorrowingRequestInput) (*pb.BorrowingRequest, error) {
	res, err := h.repository.GetRequestById(uint(input.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Borrowing request not found: %v", err)
	}

	res, err = h.repository.RejectBorrowingRequest(res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to reject borrowing request: %v", err)
	}
	response := pb.BorrowingRequest{
		Id:              uint64(res.ID),
		LendingUserId:   uint64(res.LendingUserID),
		BorrowingUserId: uint64(res.BorrowingUserID),
		LendingPostId:   uint64(res.LendingPostID),
		BorrowingPostId: uint64(res.BorrowingPostID),
		Status:          util.MapModelToProtoStatus(res.Status),
		ActiveStatus:    res.ActiveStatus,
	}

	return &response, nil
}

func (h *BorrowingGRPC) ReturnItemRequest(ctx context.Context, input *pb.ReturnItemRequestInput) (*pb.BorrowingRequest, error) {
	res, err := h.repository.GetRequestById(uint(input.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Borrowing request not found: %v", err)
	}

	res, err = h.repository.ReturnItemRequest(res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to return item from borrowing request: %v", err)
	}
	response := pb.BorrowingRequest{
		Id:              uint64(res.ID),
		LendingUserId:   uint64(res.LendingUserID),
		BorrowingUserId: uint64(res.BorrowingUserID),
		LendingPostId:   uint64(res.LendingPostID),
		BorrowingPostId: uint64(res.BorrowingPostID),
		Status:          util.MapModelToProtoStatus(res.Status),
		ActiveStatus:    res.ActiveStatus,
	}

	return &response, nil
}
