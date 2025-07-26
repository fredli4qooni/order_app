package handler

import (
	"it-order-app/backend/model"
	"it-order-app/backend/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DTO untuk input membuat order
type CreateOrderInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Category    string `json:"category" binding:"required"`
}

// DTO untuk input update status
type UpdateStatusInput struct {
	Status string `json:"status" binding:"required,oneof=Pending 'In Progress' Done"`
}

type OrderHandler struct {
	orderRepository repository.OrderRepository
}

func NewOrderHandler(orderRepo repository.OrderRepository) *OrderHandler {
	return &OrderHandler{orderRepo}
}

// Handler untuk POST /api/orders
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input CreateOrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	order := model.Order{
		UserID:      userID.(uint),
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		Status:      "Pending",
	}

	savedOrder, err := h.orderRepository.Save(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, savedOrder)
}

// Handler untuk GET /api/orders (User's own orders)
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	orders, err := h.orderRepository.FindByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Handler untuk GET /api/admin/orders
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Handler untuk PUT /api/admin/orders/:id/status
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	var input UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID format"})
		return
	}

	updatedOrder, err := h.orderRepository.UpdateStatus(uint(orderID), input.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found or failed to update status"})
		return
	}

	c.JSON(http.StatusOK, updatedOrder)
}
