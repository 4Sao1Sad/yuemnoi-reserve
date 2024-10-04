package handler

import (
    "context"

    "github.com/KKhimmoon/yuemnoi-reserve/internal/model"
    "github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
    pb "github.com/KKhimmoon/yuemnoi-reserve/proto/reserve"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type ItemGRPC struct {
    pb.UnimplementedItemServiceServer
    repository repository.ItemRepository
}

func NewItemGRPC(repo repository.ItemRepository) *ItemGRPC {
    return &ItemGRPC{
        repository: repo,
    }
}

func (g *ItemGRPC) CreateItem(ctx context.Context, input *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
    data := model.Item{
        Name: input.Name,
    }
    err := g.repository.CreateItem(data)
    if err != nil {
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }
    resp := pb.CreateItemResponse{
        Message: "created",
    }

    return &resp, nil
}
