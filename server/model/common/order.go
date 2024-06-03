package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	orderpaytype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_pay_type"
	ordersourcetype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_source_type"
	orderstatus "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_status"
	ordertype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_type"
	ordervalidstatus "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_valid_status"
	"github.com/shopspring/decimal"
)

type Order struct {
	global.GVA_MODEL
	OrderSN    string                 `json:"order_sn"`
	Type       ordertype.Value        `json:"type"`
	Status     orderstatus.Value      `json:"status"`
	Price      decimal.Decimal        `json:"price"`
	PayPrice   decimal.Decimal        `json:"pay_price"`
	PayType    orderpaytype.Value     `json:"pay_type"`
	SourceType ordersourcetype.Value  `json:"source_type"`
	IsValid    ordervalidstatus.Value `json:"is_valid"`
	Remark     string                 `json:"remark"`
}
