package util

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	pb "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
)

func MapModelToProtoStatus(status model.RequestStatus) pb.RequestStatus {
	switch status {
	case model.Pending:
		return pb.RequestStatus_PENDING
	case model.Accept:
		return pb.RequestStatus_ACCEPT
	case model.Reject:
		return pb.RequestStatus_REJECT
	default:
		return pb.RequestStatus_PENDING
	}
}
