package handler

import (
	"strconv"

	dto "github.com/KKhimmoon/yuemnoi-reserve/internal/dto"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	repo "github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/util"
	"github.com/gofiber/fiber/v2"
)

type RequestRestHandler struct {
	borrowingRequestRepository repo.BorrowingRequestRepository
	lendingRequestRepository   repo.LendingRequestRepository
}

func NewRequestRestHandler(borrowingRequestRepository repo.BorrowingRequestRepository, lendingRequestRepository repo.LendingRequestRepository) *RequestRestHandler {
	return &RequestRestHandler{
		lendingRequestRepository:   lendingRequestRepository,
		borrowingRequestRepository: borrowingRequestRepository,
	}
}

func (h *RequestRestHandler) GetMyActiveRequest(c *fiber.Ctx) error {
	userIdString := c.Get("X-User-Id")
	if userIdString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}
	borrowingReqs, err := h.borrowingRequestRepository.GetMyActiveBorrowingRequests(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve active borrowing requests",
		})
	}
	lendingReqs, err := h.lendingRequestRepository.GetMyActiveLendingRequests(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve active lending requests",
		})
	}

	var requestIdsBorrowingReq []uint64
	for _, request := range borrowingReqs {
		requestIdsBorrowingReq = append(requestIdsBorrowingReq, uint64(request.PostID))
	}

	lendingPosts, err := util.GetLendingPostsByIds(requestIdsBorrowingReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	lendingPostMap := make(map[uint64]interface{})
	for _, post := range lendingPosts.Posts {
		lendingPostMap[post.Id] = post
	}

	var response []dto.ActiveRequestResponse
	for _, request := range borrowingReqs {
		role := "borrower"
		if request.LendingUserID == uint(userId) {
			role = "lender"
		}
		name, err := util.GetUserById(uint(request.BorrowingUserID))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable retrieve user name",
			})
		}
		lendingPost, found := lendingPostMap[uint64(request.PostID)]
		if !found {
			continue
		}
		response = append(response, dto.ActiveRequestResponse{
			RequestType:     dto.BorrowingRequest,
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			PostID:          request.PostID,
			Role:            role,
			Post:            lendingPost,
			Borrower:        name,
		})
	}

	var requestIdsLendingReq []uint64
	for _, request := range lendingReqs {
		requestIdsLendingReq = append(requestIdsLendingReq, uint64(request.LendingPostID))
	}

	lendingPosts, err = util.GetLendingPostsByIds(requestIdsLendingReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	lendingPostMap = make(map[uint64]interface{})
	for _, post := range lendingPosts.Posts {
		lendingPostMap[post.Id] = post
	}
	for _, request := range lendingReqs {
		role := "lender"
		if request.BorrowingUserID == uint(userId) {
			role = "borrower"
		}
		name, err := util.GetUserById(uint(request.BorrowingUserID))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable retrieve user name",
			})
		}
		lendingPost, found := lendingPostMap[uint64(request.LendingPostID)]
		if !found {
			continue
		}
		response = append(response, dto.ActiveRequestResponse{
			RequestType:     dto.LendingRequest,
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			PostID:          request.LendingPostID,
			Role:            role,
			Post:            lendingPost,
			Borrower:        name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (h *RequestRestHandler) GetMyHistoryRequest(c *fiber.Ctx) error {
	userIdString := c.Get("X-User-Id")
	if userIdString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid userId",
		})
	}
	borrowingReqs, err := h.borrowingRequestRepository.GetMyHistorryBorrowingRequests(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve history borrowing requests",
		})
	}
	lendingReqs, err := h.lendingRequestRepository.GetMyHistorryLendingRequests(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve history lending requests",
		})
	}
	var requestIdsBorrowingReq []uint64
	for _, request := range borrowingReqs {
		requestIdsBorrowingReq = append(requestIdsBorrowingReq, uint64(request.PostID))
	}

	lendingPosts, err := util.GetLendingPostsByIds(requestIdsBorrowingReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	lendingPostMap := make(map[uint64]interface{})
	for _, post := range lendingPosts.Posts {
		lendingPostMap[post.Id] = post
	}
	var response []dto.HistoryRequestResponse
	for _, request := range borrowingReqs {
		isReject := request.Status == model.Rejected
		name, err := util.GetUserById(uint(request.BorrowingUserID))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable retrieve user name",
			})
		}
		lendingPost, found := lendingPostMap[uint64(request.PostID)]
		if !found {
			continue
		}
		response = append(response, dto.HistoryRequestResponse{
			RequestType:     dto.BorrowingRequest,
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			PostID:          request.PostID,
			IsReject:        isReject,
			Post:            lendingPost,
			Borrower:        name,
		})
	}
	var requestIdsLendingReq []uint64
	for _, request := range lendingReqs {
		requestIdsLendingReq = append(requestIdsLendingReq, uint64(request.LendingPostID))
	}

	lendingPosts, err = util.GetLendingPostsByIds(requestIdsLendingReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	lendingPostMap = make(map[uint64]interface{})
	for _, post := range lendingPosts.Posts {
		lendingPostMap[post.Id] = post
	}
	for _, request := range lendingReqs {
		isReject := request.Status == model.Rejected
		name, err := util.GetUserById(uint(request.BorrowingUserID))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable retrieve user name",
			})
		}
		lendingPost, found := lendingPostMap[uint64(request.LendingPostID)]
		if !found {
			continue
		}
		response = append(response, dto.HistoryRequestResponse{
			RequestType:     dto.LendingRequest,
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			PostID:          request.LendingPostID,
			IsReject:        isReject,
			Post:            lendingPost,
			Borrower:        name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
