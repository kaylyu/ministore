package sku_test

import (
	"fmt"
	"github.com/kaylyu/ministore"
	"github.com/kaylyu/ministore/config"
	"github.com/kaylyu/ministore/models"
	"github.com/kaylyu/ministore/util"
	"testing"
)

func TestGetSku(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_1hUbUTP68GLnLMQ2I4vbDGwq-VPgr7HAfCROktjBuO6y46YFTj085O73Jj4rsNE2gSIIqeHAzTsooiaOeEhTm0Lt3j5k6DefoCcDV-Fw2sVGpxVwGFqRCmkwffHNhYt1Od03LclVTaGZy-7GGFNgAGDJIQ",
	})
	resp, err := m.GetSku().GetSku(&models.SkuGetRequest{
		SkuId: 12566346,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetListSku(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_1hUbUTP68GLnLMQ2I4vbDGwq-VPgr7HAfCROktjBuO6y46YFTj085O73Jj4rsNE2gSIIqeHAzTsooiaOeEhTm0Lt3j5k6DefoCcDV-Fw2sVGpxVwGFqRCmkwffHNhYt1Od03LclVTaGZy-7GGFNgAGDJIQ",
	})
	resp, err := m.GetSku().GetListSku(&models.SkuGetListRequest{
		ProductId:     7521949,
		NeedRealStock: models.NeedRealStockOnline,
		NeedEditSku:   models.NeedEditSkuOnline,
	})
	fmt.Println(util.JsonEncode(resp), err)
}
