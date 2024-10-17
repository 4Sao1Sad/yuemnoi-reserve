package handler

import (
	"context"
	"fmt"
	"strconv"

	"github.com/KKhimmoon/yuemnoi-reserve/dto"

	"github.com/KKhimmoon/yuemnoi-reserve/internal/event"
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

func (h *BorrowingGRPC) CreateBorrowingRequest(ctx context.Context, input *pb.CreateBorrowingRequestInput) (*pb.CreateBorrowingRequestResponse, error) {
	data := model.BorrowingRequest{
		LendingUserID:   uint(input.LendingUserId),
		BorrowingUserID: uint(input.BorrowingUserId),
		LendingPostID:   uint(input.LendingPostId),
		BorrowingPostID: uint(input.BorrowingPostId),
		Status:          model.Pending,
		ActiveStatus:    true,
	}
	res, err := h.repository.CreateBorrowingRequest(data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create borrowing request: %v", err)
	}
	PostIdStr := strconv.FormatUint(uint64(res.BorrowingPostID), 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Offer item to borrowing post %s.", PostIdStr)
	errLog := util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if errLog != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", errLog)
	}
	requestFromBorrowingNoti := dto.NotificationRequest{
		Message: "You get a new offer, please check your Request list.",
		UserIds: []int{int(input.BorrowingUserId)},
	}
	event.SendNotification(requestFromBorrowingNoti)

	response := pb.CreateBorrowingRequestResponse{
		Message: "created successfully",
	}

	return &response, nil
}

func (h *BorrowingGRPC) GetBorrowingRequestById(ctx context.Context, input *pb.GetBorrowingRequestInput) (*pb.BorrowingRequest, error) {
	res, err := h.repository.GetBorrowingRequestById(uint(input.Id))
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

func (h *BorrowingGRPC) AcceptBorrowingRequest(ctx context.Context, input *pb.AcceptBorrowingRequestInput) (*pb.BorrowingRequest, error) {
	req, err := h.repository.GetBorrowingRequestById(uint(input.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Borrowing request not found: %v", err)
	}

	res, err := h.repository.AcceptBorrowingRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to Accept borrowing request: %v", err)
	}
	requestIdStr := strconv.FormatUint(uint64(res.ID), 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Accept borrowing request %s", requestIdStr)
	errLog := util.CallActivityLogService(uint64(res.BorrowingUserID), logDetail)
	if errLog != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", errLog)
	}

	err = util.CallPostService("BorrowingPost", uint64(res.BorrowingPostID), false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update borrowing post status: %v", err)
	}

	err = util.CallPostService("LendingPost", uint64(res.LendingPostID), false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update lending post status: %v", err)
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
	req, err := h.repository.GetBorrowingRequestById(uint(input.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Borrowing request not found: %v", err)
	}

	res, err := h.repository.RejectBorrowingRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to reject borrowing request: %v", err)
	}
	requestIdStr := strconv.FormatUint(uint64(res.ID), 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Reject borrowing request %s", requestIdStr)
	errLog := util.CallActivityLogService(uint64(res.BorrowingUserID), logDetail)
	if errLog != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", errLog)
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

func (h *BorrowingGRPC) ReturnItemBorrowingRequest(ctx context.Context, input *pb.ReturnItemBorrowingRequestInput) (*pb.BorrowingRequest, error) {
	req, err := h.repository.GetBorrowingRequestById(uint(input.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Borrowing request not found: %v", err)
	}

	res, err := h.repository.ReturnItemBorrowingRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to return item from borrowing request: %v", err)
	}
	requestIdStr := strconv.FormatUint(uint64(res.ID), 10)
	logDetail := fmt.Sprintf("Reservation Service: [success] Return Item form borrowing request %s.", requestIdStr)
	err = util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to log activity: %v", err)
	}

	err = util.CallPostService("LendingPost", uint64(res.LendingPostID), true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update lending post status: %v", err)
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
