// Package controller purchase controller for handling purchase request.
package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service/purchase"
	"swclabs/swipecore/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

// IPurchase interface for purchase controller
type IPurchase interface {
	AddToCarts(c echo.Context) error
	GetCarts(c echo.Context) error
	DeleteItem(c echo.Context) error
	CreateOrder(c echo.Context) error
	GetOrders(c echo.Context) error
}

// Purchase struct implementation of IPurchase
type Purchase struct {
	services purchase.IPurchaseService
}

var _ IPurchase = (*Purchase)(nil)

// NewPurchase creates a new Purchase object
func NewPurchase(services purchase.IPurchaseService) IPurchase {
	return &Purchase{services: services}
}

// GetOrders .
// @Description get list of orders.
// @Tags purchase
// @Accept json
// @Produce json
// @Param uid query string true "user id"
// @Success 200 {object} []domain.OrderSchema
// @Router /purchase/orders [GET]
func (purchase *Purchase) GetOrders(c echo.Context) error {
	sUserID := c.QueryParam("uid")
	if sUserID == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing 'uid' required",
		})
	}
	sLimit := c.QueryParam("limit")
	if sLimit == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing 'limit' required",
		})
	}

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "'uid' must be a positive integer",
		})
	}

	limit, err := strconv.ParseInt(sLimit, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "'limit' must be a positive integer",
		})
	}

	orders, err := purchase.services.GetOrdersByUserID(c.Request().Context(), userID, int(limit))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, orders)
}

// CreateOrder .
// @Description create order.
// @Tags purchase
// @Accept json
// @Produce json
// @Param login body domain.CreateOrderSchema true "order insert request"
// @Success 200 {object} domain.OK
// @Router /purchase/orders [POST]
func (purchase *Purchase) CreateOrder(c echo.Context) error {
	var orderReq domain.CreateOrderSchema
	if err := c.Bind(&orderReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&orderReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	msg, err := purchase.services.CreateOrders(c.Request().Context(), orderReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: fmt.Sprintf("your order %s has been created successfully", msg),
	})
}

// AddToCarts .
// @Description add item to carts.
// @Tags purchase
// @Accept json
// @Produce json
// @Param login body domain.CartInsert true "cart insert request"
// @Success 200 {object} domain.OK
// @Router /purchase/carts [POST]
func (purchase *Purchase) AddToCarts(c echo.Context) error {
	var cartReq domain.CartInsert
	if err := c.Bind(&cartReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&cartReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if err := purchase.services.AddToCart(c.Request().Context(), cartReq); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "your item has been update successfully",
	})
}

// GetCarts .
// @Description get list of items from carts
// @Tags purchase
// @Accept json
// @Produce json
// @Param uid query string true "user id"
// @Success 200 {object} domain.CartSlices
// @Router /purchase/carts [GET]
func (purchase *Purchase) GetCarts(c echo.Context) error {
	sUserID := c.QueryParam("uid")
	if sUserID == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing 'uid' required",
		})
	}
	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "'uid' must be a positive integer",
		})
	}
	items, err := purchase.services.GetCart(c.Request().Context(), userID, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, *items)
}

// DeleteItem .
// @Description delete item from carts
// @Tags purchase
// @Accept json
// @Produce json
// @Param uid query int true "user id"
// @Param wid query int true "inventories id"
// @Success 200 {object} domain.OK
// @Router /purchase/carts [DELETE]
func (purchase *Purchase) DeleteItem(c echo.Context) error {
	var (
		ids  = make(map[string]string)
		iIDs = make(map[string]int64)
	)
	const (
		uid = "uid"
		wid = "wid"
	)
	ids[uid] = c.QueryParam(uid)
	ids[wid] = c.QueryParam(wid)

	for key, param := range ids {
		if param == "" {
			return c.JSON(http.StatusBadRequest, domain.Error{
				Msg: fmt.Sprintf("missing param '%s' required", key),
			})
		}
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, domain.Error{
				Msg: fmt.Sprintf(" param '%s' must be integer", key),
			})
		}
		iIDs[key] = id
	}

	if err := purchase.services.
		DeleteItemFromCart(c.Request().Context(), iIDs[uid], iIDs[wid]); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.OK{
		Msg: "your item has been deleted successfully",
	})
}