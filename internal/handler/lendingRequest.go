package handler

import (
	"context"
	"fmt"
	"strconv"

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

func (g *LendingRequestGRPC) CreateLendingRequest(ctx context.Context, input *pb.CreateLendingRequestInput) (*pb.CreateLendingRequestResponse, error) {
	data := model.LendingRequest{
		LendingUserID:   uint64(input.LendingUserId),
		BorrowingUserID: uint64(input.BorrowingUserId),
		PostID:          uint64(input.PostId),
		Status:          model.Pending,
		ActiveStatus:    true,
	}

	res, err := g.repository.CreateLendingRequest(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create lending request: %v", err)
	}
	postIDStr := strconv.FormatUint(res.PostID, 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Reserve items from lending post %s.", postIDStr)
	err = util.CallActivityLogService(uint64(res.BorrowingUserID), logDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", err)
	}
	resp := pb.CreateLendingRequestResponse{
		Message: "created successfully",
	}

	return &resp, nil
}

func (g *LendingRequestGRPC) GetLendingRequestById(ctx context.Context, input *pb.GetLendingRequestInput) (*pb.LendingRequest, error) {
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

func (g *LendingRequestGRPC) RejectLendingRequest(ctx context.Context, input *pb.RejectLendingRequestInput) (*pb.LendingRequest, error) {
	req, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Lending request not found: %v", err)
	}

	res, err := g.repository.RejectLendingRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to reject lending request: %v", err)
	}

	requestIdStr := strconv.FormatUint(uint64(res.ID), 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Reject lending request %s.", requestIdStr)
	err = util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", err)
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

func (g *LendingRequestGRPC) AcceptLendingRequest(ctx context.Context, input *pb.AcceptLendingRequestInput) (*pb.LendingRequest, error) {
	req, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Lending request not found: %v", err)
	}

	res, err := g.repository.AcceptLendingRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to accept lending request: %v", err)
	}

	requestIdStr := strconv.FormatUint(uint64(res.ID), 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Accept lending request %s.", requestIdStr)
	err = util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", err)
	}

	err = util.CallPostService("LendingPost", uint64(res.PostID), false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update lending post status: %v", err)
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

func (g *LendingRequestGRPC) ReturnItemLendingRequest(ctx context.Context, input *pb.ReturnItemLendingRequestInput) (*pb.LendingRequest, error) {
	req, err := g.repository.GetLendingRequestById(input.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Lending request not found: %v", err)
	}

	res, err := g.repository.ReturnItemLendingRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to return item from lending request: %v", err)
	}
	requestIdStr := strconv.FormatUint(uint64(res.ID), 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Return Item form lending request %s.", requestIdStr)
	err = util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", err)
	}
	err = util.CallPostService("LendingPost", uint64(res.PostID), true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update lending post status: %v", err)
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
