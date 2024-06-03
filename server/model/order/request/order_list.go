package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	orderpaytype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_pay_type"
	ordersourcetype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_source_type"
	orderstatus "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_status"
	ordertype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_type"
	ordervalidstatus "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_valid_status"
	"github.com/shopspring/decimal"
	"time"
)

type OrderList struct {
	request.PageInfo
	Id         uint                   `json:"ID" form:"ID"`
	OrderSN    string                 `json:"order_sn" form:"order_sn"`
	Type       ordertype.Value        `json:"type" form:"type"`
	Status     orderstatus.Value      `json:"status" form:"status"`
	Price      decimal.Decimal        `json:"price" form:"price"`
	PayPrice   decimal.Decimal        `json:"pay_price" form:"pay_price"`
	PayType    orderpaytype.Value     `json:"pay_type" form:"pay_type"`
	SourceType ordersourcetype.Value  `json:"source_type" form:"source_type"`
	IsValid    ordervalidstatus.Value `json:"is_valid" form:"is_valid"`
	Remark     string                 `json:"remark" form:"remark"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

type Remove struct {
	Ids []uint `json:"ids"`
}
