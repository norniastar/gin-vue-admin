package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
)

type OrderService struct{}

func (o *OrderService) GetOrderList(req orderReq.OrderList) (res []*common.Order, total int64, err error) {
	db := global.GVA_DB

	if req.OrderSN != "" {
		db = db.Where("order_sn = ?", req.OrderSN)
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}
	if req.PayType != "" {
		db = db.Where("pay_type = ?", req.PayType)
	}
	if req.SourceType != "" {
		db = db.Where("source_type = ?", req.SourceType)
	}

	err = db.Model(common.Order{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	orders := make([]*common.Order, 0)
	err = db.Offset(req.PageSize * (req.Page - 1)).Limit(req.PageSize).Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (o *OrderService) AddOrSave(req orderReq.OrderList) error {
	db := global.GVA_DB

	if req.Id == 0 {
		return db.Create(&common.Order{
			OrderSN:    req.OrderSN,
			Type:       req.Type,
			Status:     req.Status,
			Price:      req.Price,
			PayPrice:   req.PayPrice,
			PayType:    req.PayType,
			SourceType: req.SourceType,
			IsValid:    req.IsValid,
			Remark:     req.Remark,
		}).Error
	}

	return global.GVA_DB.Where("id = ?", req.Id).UpdateColumns(&common.Order{
		OrderSN:    req.OrderSN,
		Type:       req.Type,
		Status:     req.Status,
		Price:      req.Price,
		PayPrice:   req.PayPrice,
		PayType:    req.PayType,
		SourceType: req.SourceType,
		IsValid:    req.IsValid,
		Remark:     req.Remark,
	}).Error
}

func (o *OrderService) Remove(req orderReq.Remove) error {
	db := global.GVA_DB
	return db.Where("id in (?)", req.Ids).Delete(&common.Order{}).Error
}
