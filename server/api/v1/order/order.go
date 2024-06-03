package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/gin-gonic/gin"
)

type OrderApi struct{}

func (o *OrderApi) OrderList(c *gin.Context) {
	var req orderReq.OrderList
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := orderService.GetOrderList(req)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)
}

func (o *OrderApi) AddOrSave(c *gin.Context) {
	var req orderReq.OrderList
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = orderService.AddOrSave(req)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed("ok", "获取成功", c)
}

func (o *OrderApi) Remove(c *gin.Context) {
	var req orderReq.Remove
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = orderService.Remove(req)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed("ok", "获取成功", c)
}
