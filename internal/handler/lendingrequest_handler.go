package handler

import (
	"context"
	"log"

	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
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
		LendingUserID:   input.LendingUserId,
		BorrowingUserID: input.BorrowingUserId,
		PostID:          input.PostId,
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
		Message: "created",
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
		LendingUserId:   request.LendingUserID,
		BorrowingUserId: request.BorrowingUserID,
		PostId:          request.PostID,
		Status:          mapModelToProtoStatus(request.Status),
		ActiveStatus:    request.ActiveStatus,
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) RejectLendingRequest(ctx context.Context, input *pb.RejectLendingRequestRequest) (*pb.LendingRequest, error) {
	request, err := g.repository.RejectLendingRequest(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.LendingRequest{
		Id:              uint64(request.ID),
		LendingUserId:   request.LendingUserID,
		BorrowingUserId: request.BorrowingUserID,
		PostId:          request.PostID,
		Status:          mapModelToProtoStatus(model.Reject),
		ActiveStatus:    false,
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) AcceptLendingRequest(ctx context.Context, input *pb.AcceptLendingRequestRequest) (*pb.LendingRequest, error) {
	request, err := g.repository.AcceptLendingRequest(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.LendingRequest{
		Id:              uint64(request.ID),
		LendingUserId:   request.LendingUserID,
		BorrowingUserId: request.BorrowingUserID,
		PostId:          request.PostID,
		Status:          mapModelToProtoStatus(model.Accept),
		ActiveStatus:    true,
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) ReturnItem(ctx context.Context, input *pb.ReturnItemRequest) (*pb.LendingRequest, error) {
	request, err := g.repository.ReturnItem(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.LendingRequest{
		Id:              uint64(request.ID),
		LendingUserId:   request.LendingUserID,
		BorrowingUserId: request.BorrowingUserID,
		PostId:          request.PostID,
		Status:          mapModelToProtoStatus(model.Accept),
		ActiveStatus:    false,
	}

	return &resp, nil
}

func mapModelToProtoStatus(status model.RequestStatus) pb.LendingRequestStatus {
	switch status {
	case model.Pending:
		return pb.LendingRequestStatus_PENDING
	case model.Accept:
		return pb.LendingRequestStatus_ACCEPT
	case model.Reject:
		return pb.LendingRequestStatus_REJECT
	default:
		return pb.LendingRequestStatus_PENDING
	}
}
