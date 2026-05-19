package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/ports/services"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(s services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: s}
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
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var req entities.Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	newOrder, err := h.orderService.CreateOrder(&req)
	if err != nil {
		log.Printf("[NewOrder] failed to create new order: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create new order",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newOrder)
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
// @Router /orders/mapping [put]
func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	var req entities.OrderMapping
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if _, err := h.orderService.UpdateOrder(&req); err != nil {
		log.Printf("[UpdateOrder] failed to update order mapping: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order mapping",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Order mapping updated successfully",
		"order_id":      req.OrderID,
		"list_queue_id": req.ListQueueID,
		"checked":       req.Checked,
	})
}

func (h *OrderHandler) UpdateOrderName(c *fiber.Ctx) error {
	var req struct {
		OrderID uint   `json:"order_id"`
		Title   string `json:"title"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	updatedOrder, err := h.orderService.UpdateOrderName(req.OrderID, req.Title)
	if err != nil {
		log.Printf("[UpdateOrderName] failed to update order name: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order name",
		})
	}
	return c.Status(fiber.StatusOK).JSON(updatedOrder)
}

// RemoveOrder godoc
// @Summary ลบคำสั่งซื้อ
// @Tags Orders
// @Param id path int true "Order ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [delete]
func (h *OrderHandler) RemoveOrder(c *fiber.Ctx) error {
	orderID, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("[RemoveOrder] failed to convert id to int: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	err = h.orderService.RemoveOrder(uint(orderID))
	if err != nil {
		log.Printf("[RemoveOrder] failed to remove order: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to remove order",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order removed successfully",
	})
}

// GetOrderFromListQueueID godoc
// @Summary ดึงคำสั่งซื้อทั้งหมดจาก ListQueue ที่ระบุ
// @Tags Orders
// @Param id path int true "ListQueue ID"
// @Produce json
// @Success 200 {array} entities.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/list-queue/{id} [get]
func (h *OrderHandler) GetOrderFromListQueueID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("[GetOrderFromListQueueID] failed to convert id to int: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid listQueue ID",
		})
	}

	list, err := h.orderService.GetOrderFromListQueueID(uint(id))
	if err != nil {
		log.Printf("[GetOrderFromListQueueID] failed to get order from listQueue ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get order from listQueue ID",
		})
	}

	return c.JSON(list)
}
