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
		AccessToken: "41_gaLbN9dhFEcuv-Wjpe31zXGQk893IfAr7sey0kHT_p0MoUPBYvq5---hs4GCVm_UycXx0XSKyamSLXLppfqtwnh4cQGAopjDvEZd1-Kpv2sw2JmQS3rcdc_7RkyTLHX0lEmCzvv8C4kI7OZ7LCIaADDXOU",
	})
	resp, _, err := m.GetOrder().GetOrderList(&models.OrderGetListRequest{
		StartUpdateTime: "2020-12-22 14:40:43",
		EndUpdateTime:   "2021-01-13 17:43:43",
		Page:            1,
		PageSize:        10,
		Status:          models.OrderStatusComplete,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetOrderDetail(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_1hUbUTP68GLnLMQ2I4vbDGwq-VPgr7HAfCROktjBuO6y46YFTj085O73Jj4rsNE2gSIIqeHAzTsooiaOeEhTm0Lt3j5k6DefoCcDV-Fw2sVGpxVwGFqRCmkwffHNhYt1Od03LclVTaGZy-7GGFNgAGDJIQ",
	})
	resp, _, err := m.GetOrder().GetOrderDetail(&models.OrderDetailRequest{
		OrderId: 1675794656629807,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetOrderSearch(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_1hUbUTP68GLnLMQ2I4vbDGwq-VPgr7HAfCROktjBuO6y46YFTj085O73Jj4rsNE2gSIIqeHAzTsooiaOeEhTm0Lt3j5k6DefoCcDV-Fw2sVGpxVwGFqRCmkwffHNhYt1Od03LclVTaGZy-7GGFNgAGDJIQ",
	})
	resp, _, err := m.GetOrder().GetOrderSearch(&models.OrderSearchRequest{
		Page:     1,
		PageSize: 10,
		Status:   models.OrderStatusUnpaid,
	})
	fmt.Println(util.JsonEncode(resp), err)
}
