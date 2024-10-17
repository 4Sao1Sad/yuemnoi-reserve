package handler

import (
	"context"
	"log"

	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/util"
	pb "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LendingRequestGRPC struct {
	pb.UnimplementedReserveServiceServer
	repository repository.LendingRequestRepository
}

func NewLendingRequestGRPC(repo repository.LendingRequestRepository) *LendingRequestGRPC {
	return &LendingRequestGRPC{
		repository: repo,
	}
}

func (g *LendingRequestGRPC) CreateLendingRequest(ctx context.Context, input *pb.CreateLendingRequestRequest) (*pb.CreateLendingRequestResponse, error) {
	data := model.LendingRequest{
		LendingUserID:   uint64(input.LendingUserId),
		BorrowingUserID: uint64(input.BorrowingUserId),
		PostID:          uint64(input.PostId),
		Status:          model.Pending,
		ActiveStatus:    true,
	}
	log.Println(data)
	_, err := g.repository.CreateLendingRequest(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateLendingRequestResponse{
		Id:      uint64(data.ID),
		Message: "created successfully",
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) GetLendingRequestDetail(ctx context.Context, input *pb.GetLendingRequestDetailRequest) (*pb.LendingRequest, error) {
	request, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.LendingRequest{
		Id:              uint64(request.ID),
		LendingUserId:   uint64(request.LendingUserID),
		BorrowingUserId: uint64(request.BorrowingUserID),
		PostId:          uint64(request.PostID),
		Status:          util.MapModelToProtoStatus(request.Status),
		ActiveStatus:    request.ActiveStatus,
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) RejectLendingRequest(ctx context.Context, input *pb.RejectLendingRequestRequest) (*pb.LendingRequest, error) {
	res, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Lending request not found: %v", err)
	}

	res, err = g.repository.RejectLendingRequest(res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to reject lending request: %v", err)
	}
	resp := pb.LendingRequest{
		Id:              uint64(res.ID),
		LendingUserId:   uint64(res.LendingUserID),
		BorrowingUserId: uint64(res.BorrowingUserID),
		PostId:          uint64(res.PostID),
		Status:          util.MapModelToProtoStatus(res.Status),
		ActiveStatus:    res.ActiveStatus,
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) AcceptLendingRequest(ctx context.Context, input *pb.AcceptLendingRequestRequest) (*pb.LendingRequest, error) {
	res, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Lending request not found: %v", err)
	}

	res, err = g.repository.AcceptLendingRequest(res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to confirm lending request: %v", err)
	}

	resp := pb.LendingRequest{
		Id:              uint64(res.ID),
		LendingUserId:   uint64(res.LendingUserID),
		BorrowingUserId: uint64(res.BorrowingUserID),
		PostId:          uint64(res.PostID),
		Status:          util.MapModelToProtoStatus(res.Status),
		ActiveStatus:    res.ActiveStatus,
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) ReturnItemRequest(ctx context.Context, input *pb.ReturnItemRequest) (*pb.LendingRequest, error) {
	res, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Lending request not found: %v", err)
	}

	res, err = g.repository.ReturnItemRequest(res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to return item from lending request: %v", err)
	}

	resp := pb.LendingRequest{
		Id:              uint64(res.ID),
		LendingUserId:   uint64(res.LendingUserID),
		BorrowingUserId: uint64(res.BorrowingUserID),
		PostId:          uint64(res.PostID),
		Status:          util.MapModelToProtoStatus(res.Status),
		ActiveStatus:    res.ActiveStatus,
	}

	return &resp, nil
}
