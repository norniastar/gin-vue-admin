package orderstatus

type Value int

const (
	Wait   Value = iota + 1 // 待支付
	Paid                    // 已支付
	Closed                  // 已关闭
	Refund                  // 已退款
)
