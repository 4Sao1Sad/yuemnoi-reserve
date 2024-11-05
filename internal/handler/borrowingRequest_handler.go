package handler

import (
	"strconv"

	"github.com/KKhimmoon/yuemnoi-reserve/internal/dto"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/event"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/model"
	repo "github.com/KKhimmoon/yuemnoi-reserve/internal/repository"
	"github.com/KKhimmoon/yuemnoi-reserve/internal/util"
	"github.com/gofiber/fiber/v2"
)

type BorrowingRequestRestHandler struct {
	borrowingRequestRepository repo.BorrowingRequestRepository
}

func NewBorrowingRequestRestHandler(borrowingRequestRepository repo.BorrowingRequestRepository) *BorrowingRequestRestHandler {
	return &BorrowingRequestRestHandler{
		borrowingRequestRepository: borrowingRequestRepository,
	}
}
func (h *BorrowingRequestRestHandler) CreateBorrowingRequest(c *fiber.Ctx) error {
	body := new(dto.CreateBorrowingRequestInput)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input format",
		})
	}
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
	borrowingRequest := model.BorrowingRequest{
		LendingUserID:   body.LendingUserID,
		BorrowingUserID: uint(userId),
		PostID:          body.PostID,
		Status:          model.Pending,
		ActiveStatus:    true,
	}
	if err := util.ValidatePostExists(nil, uint64(borrowingRequest.PostID)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post doesn't exist",
		})
	}
	_, err = util.GetUserById(uint(borrowingRequest.BorrowingUserID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User doesn't exist",
		})
	}
	_, err = util.GetUserById(uint(borrowingRequest.LendingUserID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User doesn't exist",
		})
	}
	if err := util.CheckPostIsReady(nil, uint64(borrowingRequest.PostID)); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Posts are not ready",
		})
	}
	request, err := h.borrowingRequestRepository.CreateBorrowingRequest(borrowingRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create borrowing request",
		})
	}
	postId := strconv.FormatUint(uint64(request.PostID), 10)
	logDetail := "Reservation Service: [success] Reserve items from lending post id = " + postId
	if errLog := util.CallActivityLogService(uint64(borrowingRequest.BorrowingUserID), logDetail); errLog != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}
	notification := dto.NotificationRequest{
		Message: "You get a new request, please check your Request list.",
		UserIds: []int{int(borrowingRequest.LendingUserID)},
	}
	event.SendNotification(notification)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": request,
	})
}
func (h *BorrowingRequestRestHandler) GetMyBorrowingRequests(c *fiber.Ctx) error {
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

	requests, err := h.borrowingRequestRepository.GetMyBorrowingRequests(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve borrowing requests",
		})
	}

	var requestIdsBorrowingReq []uint64
	for _, request := range requests {
		requestIdsBorrowingReq = append(requestIdsBorrowingReq, uint64(request.PostID))
	}
	lendingPosts, err := util.GetLendingPostsByIds(requestIdsBorrowingReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	var response []dto.GetMyBorrowingRequestsResponse
	for i, request := range requests {
		response = append(response, dto.GetMyBorrowingRequestsResponse{
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			PostID:          request.PostID,
			Status:          request.Status,
			ActiveStatus:    request.ActiveStatus,
			Post:            lendingPosts.Posts[i],
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": response,
	})
}

func (h *BorrowingRequestRestHandler) GetMyLendingPosts(c *fiber.Ctx) error {
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

	requests, err := h.borrowingRequestRepository.GetMyLendingPosts(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}

	var lendingPostids []uint64
	for _, request := range requests {
		lendingPostids = append(lendingPostids, uint64(request.PostID))
	}

	lendingPosts, err := util.GetLendingPostsByIds(lendingPostids)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	var response []dto.GetMyLendingPostsResponse
	for i, request := range requests {
		name, err := util.GetUserById(uint(userId))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable retrieve user name",
			})
		}
		response = append(response, dto.GetMyLendingPostsResponse{
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			PostID:          request.PostID,
			Status:          request.Status,
			ActiveStatus:    request.ActiveStatus,
			Post:            lendingPosts.Posts[i],
			Borrower:        name,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (h *BorrowingRequestRestHandler) AcceptBorrowingRequest(c *fiber.Ctx) error {
	requestIdStr := c.Params("requestId")
	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID format",
		})
	}
	req, err := h.borrowingRequestRepository.GetBorrowingRequestById(uint(requestId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Request not found",
		})
	}
	if err := util.ValidateRequest(req.Status, req.ActiveStatus); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Invalid request status",
		})
	}
	if err := util.CheckPostIsReady(nil, uint64(req.PostID)); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Post is not ready for lending",
		})
	}
	res, err := h.borrowingRequestRepository.AcceptBorrowingRequest(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to accept borrowing request",
		})
	}
	err = util.UpdatePostService("LendingPost", uint64(res.PostID), false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update status on post",
		})

	}
	logDetail := "Reservation Service: [success] Accept borrowing request id = " + requestIdStr
	err = util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}

	acceptBorrowingRequestNoti := dto.NotificationRequest{
		Message: "Your request is accepted, please check the Active items.",
		UserIds: []int{int(res.BorrowingUserID)},
	}
	event.SendNotification(acceptBorrowingRequestNoti)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (h *BorrowingRequestRestHandler) RejectBorrowingRequest(c *fiber.Ctx) error {
	requestIdStr := c.Params("requestId")
	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID format",
		})
	}
	req, err := h.borrowingRequestRepository.GetBorrowingRequestById(uint(requestId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Request not found",
		})
	}
	if err := util.ValidateRequest(req.Status, req.ActiveStatus); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Invalid request status",
		})
	}
	res, err := h.borrowingRequestRepository.RejectBorrowingRequest(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to reject borrowing request",
		})
	}
	logDetail := "Reservation Service: [success] Reject borrowing request id  = " + requestIdStr
	err = util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}
	rejectBorrowingRequestNoti := dto.NotificationRequest{
		Message: "Your request is rejected.",
		UserIds: []int{int(res.BorrowingUserID)},
	}
	event.SendNotification(rejectBorrowingRequestNoti)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (h *BorrowingRequestRestHandler) ReturnItemBorrowingRequest(c *fiber.Ctx) error {
	requestIdStr := c.Params("requestId")
	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID format",
		})
	}
	req, err := h.borrowingRequestRepository.GetBorrowingRequestById(uint(requestId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Request not found",
		})
	}
	if err := util.ValidateReturnItemRequest(req.Status, req.ActiveStatus); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Invalid return request",
		})
	}
	res, err := h.borrowingRequestRepository.ReturnItemBorrowingRequest(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to returnItem borrowing request",
		})
	}
	err = util.UpdatePostService("LendingPost", uint64(res.PostID), true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update status on post",
		})
	}
	logDetail := "Reservation Service: [success] Return Item form borrowing request id = " + requestIdStr
	err = util.CallActivityLogService(uint64(res.LendingUserID), logDetail)
	if err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}

	returnItemRequestNoti := dto.NotificationRequest{
		Message: "Your Items returning is confirmed.",
		UserIds: []int{int(res.LendingUserID)},
	}
	event.SendNotification(returnItemRequestNoti)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}
