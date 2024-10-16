package util

import (
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	pb "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
)

func MapModelToProtoStatus(status model.RequestStatus) pb.RequestStatus {
	switch status {
	case model.Pending:
		return pb.RequestStatus_PENDING
	case model.Accepted:
		return pb.RequestStatus_ACCEPTED
	case model.Rejected:
		return pb.RequestStatus_REJECTED
	default:
		return pb.RequestStatus_PENDING
	}
}
