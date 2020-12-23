package order_test

import (
	"fmt"
	"github.com/kaylyu/ministore"
	"github.com/kaylyu/ministore/config"
	"github.com/kaylyu/ministore/models"
	"github.com/kaylyu/ministore/util"
	"testing"
)

func TestGetOrderList(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_1hUbUTP68GLnLMQ2I4vbDGwq-VPgr7HAfCROktjBuO6y46YFTj085O73Jj4rsNE2gSIIqeHAzTsooiaOeEhTm0Lt3j5k6DefoCcDV-Fw2sVGpxVwGFqRCmkwffHNhYt1Od03LclVTaGZy-7GGFNgAGDJIQ",
	})
	resp, err := m.GetOrder().GetOrderList(&models.OrderGetListRequest{
		StartCreateTime: "2020-12-22 14:40:43",
		EndCreateTime:   "2020-12-22 17:43:43",
		Page:            1,
		PageSize:        10,
		Status:          models.OrderStatusUnpaid,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetOrderDetail(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_1hUbUTP68GLnLMQ2I4vbDGwq-VPgr7HAfCROktjBuO6y46YFTj085O73Jj4rsNE2gSIIqeHAzTsooiaOeEhTm0Lt3j5k6DefoCcDV-Fw2sVGpxVwGFqRCmkwffHNhYt1Od03LclVTaGZy-7GGFNgAGDJIQ",
	})
	resp, err := m.GetOrder().GetOrderDetail(&models.OrderDetailRequest{
		OrderId: 1675794656629807,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetOrderSearch(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_1hUbUTP68GLnLMQ2I4vbDGwq-VPgr7HAfCROktjBuO6y46YFTj085O73Jj4rsNE2gSIIqeHAzTsooiaOeEhTm0Lt3j5k6DefoCcDV-Fw2sVGpxVwGFqRCmkwffHNhYt1Od03LclVTaGZy-7GGFNgAGDJIQ",
	})
	resp, err := m.GetOrder().GetOrderSearch(&models.OrderSearchRequest{
		Page:     1,
		PageSize: 10,
		Status:   models.OrderStatusUnpaid,
	})
	fmt.Println(util.JsonEncode(resp), err)
}
