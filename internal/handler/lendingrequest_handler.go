package handler

import (
	"context"

	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	pb "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	data := model.Request{
		RequestUserID:  input.RequestUserId,
		ResponseUserID: input.ResponseUserId,
		PostID:         input.PostId,
		Type:           model.Lending,
		Status:         model.Pending,
	}

	_, err := g.repository.CreateLendingRequest(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateLendingRequestResponse{
		Message: "created",
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) GetLendingRequestDetail(ctx context.Context, input *pb.GetLendingRequestDetailRequest) (*pb.Request, error) {
	request, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.Request{
		Id:             uint64(request.ID),
		RequestUserId:  request.RequestUserID,
		ResponseUserId: request.ResponseUserID,
		PostId:         request.PostID,
		Status:         mapModelToProtoStatus(request.Status),
		UpdatedAt:      timestamppb.New(request.UpdatedAt),
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) RejectLendingRequest(ctx context.Context, input *pb.RejectLendingRequestRequest) (*pb.Request, error) {
	request, err := g.repository.RejectLendingRequest(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.Request{
		Id:             uint64(request.ID),
		RequestUserId:  request.RequestUserID,
		ResponseUserId: request.ResponseUserID,
		PostId:         request.PostID,
		Status:         mapModelToProtoStatus(model.Reject),
		UpdatedAt:      timestamppb.New(request.UpdatedAt),
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) AcceptLendingRequest(ctx context.Context, input *pb.AcceptLendingRequestRequest) (*pb.Request, error) {
	request, err := g.repository.AcceptLendingRequest(input.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	resp := pb.Request{
		Id:             uint64(request.ID),
		RequestUserId:  request.RequestUserID,
		ResponseUserId: request.ResponseUserID,
		PostId:         request.PostID,
		Status:         mapModelToProtoStatus(model.Accept),
		UpdatedAt:      timestamppb.New(request.UpdatedAt),
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
