package response

import (
	orderpaytype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_pay_type"
	ordersourcetype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_source_type"
	orderstatus "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_status"
	ordertype "github.com/flipped-aurora/gin-vue-admin/server/model/enum/order_type"
	"github.com/shopspring/decimal"
	"time"
)

type OrderList struct {
	ID         uint                  `json:"id"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
	OrderSN    string                `json:"order_sn"`
	Type       ordertype.Value       `json:"type"`
	Status     orderstatus.Value     `json:"status"`
	Price      decimal.Decimal       `json:"price"`
	PayPrice   decimal.Decimal       `json:"pay_price"`
	PayType    orderpaytype.Value    `json:"pay_type"`
	SourceType ordersourcetype.Value `json:"source_type"`
}
