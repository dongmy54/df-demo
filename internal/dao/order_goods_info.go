// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-demo/internal/dao/internal"
)

// internalOrderGoodsInfoDao is internal type for wrapping internal DAO implements.
type internalOrderGoodsInfoDao = *internal.OrderGoodsInfoDao

// orderGoodsInfoDao is the data access object for table order_goods_info.
// You can define custom methods on it to extend its functionality as you wish.
type orderGoodsInfoDao struct {
	internalOrderGoodsInfoDao
}

var (
	// OrderGoodsInfo is globally public accessible object for table order_goods_info operations.
	OrderGoodsInfo = orderGoodsInfoDao{
		internal.NewOrderGoodsInfoDao(),
	}
)

// Fill with you ideas below.
