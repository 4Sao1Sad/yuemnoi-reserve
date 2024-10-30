package util

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateRequest(reqStatus model.RequestStatus, activeStatus bool) error {
	if reqStatus != model.Pending {
		return status.Errorf(codes.FailedPrecondition, "Request is not in pending status")
	}
	if !activeStatus {
		return status.Errorf(codes.FailedPrecondition, "Request is not active")
	}
	return nil
}

func ValidateReturnItemRequest(reqStatus model.RequestStatus, activeStatus bool) error {
	if reqStatus != model.Accepted {
		return status.Errorf(codes.FailedPrecondition, "Request is not in accepted status")
	}
	if !activeStatus {
		return status.Errorf(codes.FailedPrecondition, "Request is not active")
	}
	return nil
}
