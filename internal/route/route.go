package route

import (
	"github.com/KKhimmoon/yuemnoi-reserve/config"
	handler "github.com/KKhimmoon/yuemnoi-reserve/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	borrowingRequestRestHandler *handler.BorrowingRequestRestHandler
	lendingRequestRestHandler   *handler.LendingRequestRestHandler
	requestRestHandler          *handler.RequestRestHandler
}

func NewHandler(borrowingRequestRestHandler *handler.BorrowingRequestRestHandler, lendingRequestRestHandler *handler.LendingRequestRestHandler, requestRestHandler *handler.RequestRestHandler) *Handler {
	return &Handler{
		borrowingRequestRestHandler: borrowingRequestRestHandler,
		lendingRequestRestHandler:   lendingRequestRestHandler,
		requestRestHandler:          requestRestHandler,
	}
}

func (h *Handler) RegisterRouter(r fiber.Router, cfg *config.Config) {
	{
		borrowingRequestRouter := r.Group("/borrowing-requests")
		borrowingRequestRouter.Get("/my-requests", h.borrowingRequestRestHandler.GetMyBorrowingRequests)
		borrowingRequestRouter.Get("/my-lending-posts", h.borrowingRequestRestHandler.GetMyLendingPosts)
		borrowingRequestRouter.Post("/accept/:requestId", h.borrowingRequestRestHandler.AcceptBorrowingRequest)
		borrowingRequestRouter.Post("/reject/:requestId", h.borrowingRequestRestHandler.RejectBorrowingRequest)
		borrowingRequestRouter.Post("return/:requestId", h.borrowingRequestRestHandler.ReturnItemBorrowingRequest)
		borrowingRequestRouter.Post("/", h.borrowingRequestRestHandler.CreateBorrowingRequest)
	}
	{
		lendingRequestRouter := r.Group("/lending-requests")
		lendingRequestRouter.Get("/my-requests", h.lendingRequestRestHandler.GetMyLendingRequests)
		lendingRequestRouter.Get("/my-borrowing-posts", h.lendingRequestRestHandler.GetMyBorrowingPosts)
		lendingRequestRouter.Post("/accept/:requestId", h.lendingRequestRestHandler.AcceptLendingRequest)
		lendingRequestRouter.Post("/reject/:requestId", h.lendingRequestRestHandler.RejectLendingRequest)
		lendingRequestRouter.Post("return/:requestId", h.lendingRequestRestHandler.ReturnItemLendingRequest)
		lendingRequestRouter.Post("/", h.lendingRequestRestHandler.ReturnItemLendingRequest)
	}
	{
		r.Get("/active-request", h.requestRestHandler.GetMyActiveRequest)
		r.Get("/history-request", h.requestRestHandler.GetMyHistoryRequest)
	}
}
