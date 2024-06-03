package ordervalidstatus

type Value int

const (
	Valid   Value = iota + 1 // 待支付
	UnValid                  // 已支付
)
