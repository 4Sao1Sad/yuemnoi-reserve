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

type LendingRequestRestHandler struct {
	lendingRequestRepository repo.LendingRequestRepository
}

func NewLendingRequestRestHandler(lendingRequestRepository repo.LendingRequestRepository) *LendingRequestRestHandler {
	return &LendingRequestRestHandler{
		lendingRequestRepository: lendingRequestRepository,
	}
}

func (h *LendingRequestRestHandler) CreateLendingRequest(c *fiber.Ctx) error {
	body := new(dto.CreateLendingRequestInput)
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
	lendingRequest := model.LendingRequest{
		LendingUserID:   uint(userId),
		BorrowingUserID: body.BorrowingUserID,
		LendingPostID:   body.LendingPostID,
		BorrowingPostID: body.BorrowingPostID,
		Status:          model.Pending,
		ActiveStatus:    true,
	}

	borrowingPostID := uint64(lendingRequest.BorrowingPostID)
	if err := util.ValidatePostExists(&borrowingPostID, uint64(lendingRequest.LendingPostID)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Posts don't exist",
		})
	}
	_, err = util.GetUserById(uint(lendingRequest.BorrowingUserID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User don't exist",
		})
	}
	_, err = util.GetUserById(uint(lendingRequest.LendingUserID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User don't exist",
		})
	}
	if err := util.CheckPostIsReady(&borrowingPostID, uint64(lendingRequest.LendingPostID)); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Posts are not ready",
		})
	}
	request, err := h.lendingRequestRepository.CreateLendingRequest(lendingRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create Lending request",
		})
	}
	PostId := strconv.FormatUint(uint64(request.BorrowingUserID), 10)
	logDetail := "Reservation Service: [success] Offer item to borrowing post id = " + PostId
	errLog := util.CallActivityLogService(uint64(request.LendingUserID), logDetail)
	if errLog != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}

	requestFromLendingNoti := dto.NotificationRequest{
		Message: "You get a new offer, please check your Request list.",
		UserIds: []int{int(request.BorrowingUserID)},
	}
	event.SendNotification(requestFromLendingNoti)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": request,
	})
}
func (h *LendingRequestRestHandler) GetMyLendingRequests(c *fiber.Ctx) error {
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
	requests, err := h.lendingRequestRepository.GetMyLendingRequests(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve Lending requests",
		})
	}

	var lendingPostids []uint64
	for _, request := range requests {
		lendingPostids = append(lendingPostids, uint64(request.LendingPostID))
	}

	lendingPosts, err := util.GetLendingPostsByIds(lendingPostids)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}

	var borrowingPostids []uint64
	for _, request := range requests {
		borrowingPostids = append(borrowingPostids, uint64(request.BorrowingPostID))
	}

	borrowingPosts, err := util.GetBorrowingPostsByIds(borrowingPostids)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	var response []dto.GetMyLendingRequestsResponse
	for i, request := range requests {
		response = append(response, dto.GetMyLendingRequestsResponse{
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			BorrowingPostID: request.BorrowingPostID,
			LendingPostID:   request.LendingPostID,
			Status:          request.Status,
			ActiveStatus:    request.ActiveStatus,
			LendingPost:     lendingPosts.Posts[i],
			BorrowingPost:   borrowingPosts.BorrowingPost[i],
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": response,
	})
}

func (h *LendingRequestRestHandler) GetMyBorrowingPosts(c *fiber.Ctx) error {
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
	requests, err := h.lendingRequestRepository.GetMyBorrowingPosts(uint(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	var lendingPostids []uint64
	for _, request := range requests {
		lendingPostids = append(lendingPostids, uint64(request.LendingPostID))
	}

	lendingPosts, err := util.GetLendingPostsByIds(lendingPostids)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	var borrowingPostids []uint64
	for _, request := range requests {
		borrowingPostids = append(borrowingPostids, uint64(request.BorrowingPostID))
	}

	borrowingPosts, err := util.GetBorrowingPostsByIds(borrowingPostids)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve lending posts",
		})
	}
	var response []dto.GetMyBorrowingPostsResponse
	for i, request := range requests {
		response = append(response, dto.GetMyBorrowingPostsResponse{
			ID:              request.ID,
			BorrowingUserID: request.BorrowingUserID,
			LendingUserID:   request.LendingUserID,
			BorrowingPostID: request.BorrowingPostID,
			LendingPostID:   request.LendingPostID,
			Status:          request.Status,
			ActiveStatus:    request.ActiveStatus,
			LendingPost:     lendingPosts.Posts[i],
			BorrowingPost:   borrowingPosts.BorrowingPost[i],
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (h *LendingRequestRestHandler) AcceptLendingRequest(c *fiber.Ctx) error {
	requestIdStr := c.Params("requestId")
	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID format",
		})
	}
	req, err := h.lendingRequestRepository.GetLendingRequestById(uint(requestId))
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
	borrowingPostID := uint64(req.BorrowingPostID)
	if err := util.CheckPostIsReady(&borrowingPostID, uint64(req.LendingPostID)); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Post is not ready for lending",
		})
	}
	res, err := h.lendingRequestRepository.AcceptLendingRequest(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to accept Lending request",
		})
	}
	err = util.UpdatePostService("BorrowingPost", uint64(res.BorrowingPostID), false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update status on post",
		})
	}

	err = util.UpdatePostService("LendingPost", uint64(res.LendingPostID), false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update status on post",
		})
	}

	logDetail := "Reservation Service: [success] Accept lending request id = " + requestIdStr
	errLog := util.CallActivityLogService(uint64(res.BorrowingUserID), logDetail)
	if errLog != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}
	acceptLendingRequestNoti := dto.NotificationRequest{
		Message: "Your offer is accepted, please check the Active items.",
		UserIds: []int{int(res.LendingUserID)},
	}
	event.SendNotification(acceptLendingRequestNoti)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (h *LendingRequestRestHandler) RejectLendingRequest(c *fiber.Ctx) error {
	requestIdStr := c.Params("requestId")
	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID format",
		})
	}
	req, err := h.lendingRequestRepository.GetLendingRequestById(uint(requestId))
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
	res, err := h.lendingRequestRepository.RejectLendingRequest(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to reject lending request",
		})
	}
	logDetail := "Reservation Service: [success] Reject lending request id  = " + requestIdStr
	errLog := util.CallActivityLogService(uint64(res.BorrowingUserID), logDetail)
	if errLog != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"warning": "failed to create activity log",
		})
	}
	rejectLendingRequestNoti := dto.NotificationRequest{
		Message: "Your offer is rejected.",
		UserIds: []int{int(res.LendingUserID)},
	}
	event.SendNotification(rejectLendingRequestNoti)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (h *LendingRequestRestHandler) ReturnItemLendingRequest(c *fiber.Ctx) error {
	requestIdStr := c.Params("requestId")
	requestId, err := strconv.Atoi(requestIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request ID format",
		})
	}
	req, err := h.lendingRequestRepository.GetLendingRequestById(uint(requestId))
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
	res, err := h.lendingRequestRepository.ReturnItemLendingRequest(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to returnItem lending request",
		})
	}
	err = util.UpdatePostService("LendingPost", uint64(res.LendingPostID), true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update status on post",
		})
	}

	logDetail := "Reservation Service: [success] Return Item form lending request id = " + requestIdStr
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
