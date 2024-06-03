package order

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (e *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouter := Router.Group("order")
	orderApi := v1.ApiGroupApp.OrderApiGroup.OrderApi
	{
		orderRouter.GET("list", orderApi.OrderList)
		orderRouter.POST("add-or-save", orderApi.AddOrSave)
		orderRouter.POST("remove", orderApi.Remove)
	}
}
