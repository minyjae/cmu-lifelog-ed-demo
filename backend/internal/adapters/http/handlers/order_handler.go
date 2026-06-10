package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
	"github.com/minyjae/cmu-life-long-ed-api/pkg/utils"
)

type OrderHandler struct {
	orderService services.OrderService
	res          utils.IResponse
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: s, res: utils.NewResponse()}
}

// CreateOrder godoc
// @Summary สร้างคำสั่งซื้อใหม่
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body entities.Order true "Order Request"
// @Success 201 {object} entities.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/order [post]
// @Security BearerAuth
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var req entities.Order
	if err := c.BodyParser(&req); err != nil {
		return h.res.BadRequest(c, "Cannot parse JSON", utils.CodeInvalidRequest)
	}

	newOrder, err := h.orderService.CreateOrder(&req)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to create new order", err.Error(), utils.CodeInternalError)
	}

	return h.res.Created(c, "Create order successfully", newOrder)
}

// UpdateOrder godoc
// @Summary อัปเดตการจับคู่คำสั่งซื้อกับคิว
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body entities.OrderMapping true "Order Mapping Update"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/order [put]
// @Security BearerAuth
func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	var req entities.OrderMapping
	if err := c.BodyParser(&req); err != nil {
		return h.res.BadRequest(c, "Cannot parse JSON", utils.CodeInvalidRequest)
	}

	if _, err := h.orderService.UpdateOrder(&req); err != nil {
		return h.res.InternalServerError(c, "Failed to update order mapping", err.Error(), utils.CodeInternalError)
	}

	return h.res.Updated(c, "Order mapping updated successfully", fiber.Map{
		"order_id":      req.OrderID,
		"list_queue_id": req.ListQueueID,
		"checked":       req.Checked,
	})
}

// UpdateOrderName godoc
// @Summary อัปเดตชื่อคำสั่งซื้อ
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body object true "Order ID and new title"
// @Success 200 {object} entities.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/order/name [put]
// @Security BearerAuth
func (h *OrderHandler) UpdateOrderName(c *fiber.Ctx) error {
	var req struct {
		OrderID uint   `json:"order_id"`
		Title   string `json:"title"`
	}
	if err := c.BodyParser(&req); err != nil {
		return h.res.BadRequest(c, "Cannot parse JSON", utils.CodeInvalidRequest)
	}
	updatedOrder, err := h.orderService.UpdateOrderName(req.OrderID, req.Title)
	if err != nil {
		return h.res.InternalServerError(c, "Failed to update order name", err.Error(), utils.CodeInternalError)
	}
	return h.res.Updated(c, "Update order name successfully", updatedOrder)
}

// RemoveOrder godoc
// @Summary ลบคำสั่งซื้อ
// @Tags Orders
// @Param id path int true "Order ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/order/{id} [delete]
// @Security BearerAuth
func (h *OrderHandler) RemoveOrder(c *fiber.Ctx) error {
	orderID, err := c.ParamsInt("id")
	if err != nil {
		return h.res.BadRequest(c, "Invalid order ID", utils.CodeInvalidID)
	}

	if err := h.orderService.RemoveOrder(uint(orderID)); err != nil {
		return h.res.InternalServerError(c, "Failed to remove order", err.Error(), utils.CodeInternalError)
	}

	return h.res.Deleted(c, "Order removed successfully")
}

// GetOrderFromListQueueID godoc
// @Summary ดึงคำสั่งซื้อทั้งหมดจาก ListQueue ที่ระบุ
// @Tags Orders
// @Param id path int true "ListQueue ID"
// @Produce json
// @Success 200 {array} entities.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/order/{id} [get]
// @Security BearerAuth
func (h *OrderHandler) GetOrderFromListQueueID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return h.res.BadRequest(c, "Invalid listQueue ID", utils.CodeInvalidID)
	}

	list, err := h.orderService.GetOrderFromListQueueID(uint(id))
	if err != nil {
		return h.res.InternalServerError(c, "Failed to get order from listQueue ID", err.Error(), utils.CodeInternalError)
	}

	return h.res.Item(c, "Get orders successfully", list)
}
