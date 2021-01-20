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
		AccessToken: "41_LG-HTeNNdno_pfzF8UZ8XXdywpvmgHHF3neRzsx8lH54J2WYSjS8Z6Mm-vx3X8cB3vpalMlOT_Zv6SgL8nIjJ6j38BmymopIxMyAbFdDUA-Tgz-xKYz5aSnx66e0WZ-o0oca96-sGYMgAKDPHV",
	})
	keys := map[models.OrderStatus]interface{}{
		10:  "待付款",
		20:  "待发货",
		30:  "待收货",
		100: "完成",
		200: "全部商品售后之后，订单取消",
		250: "用户主动取消或待付款超时取消",
	}
	fmt.Println("小店名称：", "xxx")
	st := "2021-01-19 00:00:00"
	et := "2021-02-13 17:43:43"
	fmt.Printf("搜索时间：%s----%s\n", st, et)
	for _, item := range []models.OrderStatus{models.OrderStatusUnpaid, models.OrderStatusSendPending, models.OrderStatusReceivePending, models.OrderStatusComplete, models.OrderStatusCompleteCancel, models.OrderStatusTimeoutCancel} {
		resp, _, _ := m.GetOrder().GetOrderList(&models.OrderGetListRequest{
			StartUpdateTime: st,
			EndUpdateTime:   et,
			Page:            1,
			PageSize:        10,
			Status:          item,
		})
		fmt.Printf("%s:%d单\n", keys[item], resp.TotalNum)
	}
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
