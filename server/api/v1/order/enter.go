package order

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	OrderApi
}

var (
	orderService = service.ServiceGroupApp.OrderServiceGroup.OrderService
)
