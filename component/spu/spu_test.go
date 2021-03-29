// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spu_test

import (
	"fmt"
	"github.com/kaylyu/ministore"
	"github.com/kaylyu/ministore/config"
	"github.com/kaylyu/ministore/models"
	"github.com/kaylyu/ministore/util"
	"testing"
)

func TestGetSpuList(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "43_qAxTOftN_eEqtdSoQnbuYjWDqorgys3phCVrHBl8_JPrJpFUVq6mr3vyMMSY1Hr5wecPkmfGkNRSCV--Y8_6j4dy8C1-rytrzCIxpPQOPaHmH2hIiFjE_yAr4XVYaXCKJq9t5gqPYoT4fm45GBHeAJDQLR",
	})
	resp, _, err := m.GetSpu().GetSpuList(&models.SpuGetListRequest{
		NeedEditSpu: models.NeedEditSpuOnline,
		Status:      models.SpuStatusPutaway,
		Page:        1,
		PageSize:    10,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestSearchSpu(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_R2De-VYJKpTnEN42Z83neFP8Kmav9ARQ2liHZ2eUBEvSxmqIlh5z4IhQUYgPc0Zt-pXzciYeN381GRfcwr53zZRZJSZzF645HZV98ryPegGNR7w-uQNBM7ZgrOLOxLMRXG3LNq1JNe0j2n6_LFBaAKDGAW",
	})
	resp, _, err := m.GetSpu().SearchSpu(&models.SpuSearchRequest{
		NeedEditSpu: models.NeedEditSpuDraft,
		Status:      models.SpuStatusPutaway,
		Page:        1,
		PageSize:    10,
		Keyword:     "乐品乐茶茶叶",
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestUpSpu(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_8DHWqetYL-mSUlViAfIkmhvmqwhX94wSjBfpUlBnG3aqpb1-3jHZrW924Psgk7WTpTjqqTapq4BGXXfElVdMNLc6iwoyevPA2w-pXAZQAj85F48X4CiGm9Fg-SskNTNXcjVUi3kWua5EdwbBVBYgAMDEKH",
	})
	resp, _, err := m.GetSpu().UpSpu(&models.SpuUpRequest{
		ProductId: 7516256,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestDownSpu(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_8DHWqetYL-mSUlViAfIkmhvmqwhX94wSjBfpUlBnG3aqpb1-3jHZrW924Psgk7WTpTjqqTapq4BGXXfElVdMNLc6iwoyevPA2w-pXAZQAj85F48X4CiGm9Fg-SskNTNXcjVUi3kWua5EdwbBVBYgAMDEKH",
	})
	resp, _, err := m.GetSpu().DownSpu(&models.SpuDownRequest{
		ProductId: 7516256,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestSpuLimitedDiscountAdd(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "43_qAxTOftN_eEqtdSoQnbuYjWDqorgys3phCVrHBl8_JPrJpFUVq6mr3vyMMSY1Hr5wecPkmfGkNRSCV--Y8_6j4dy8C1-rytrzCIxpPQOPaHmH2hIiFjE_yAr4XVYaXCKJq9t5gqPYoT4fm45GBHeAJDQLR",
	})
	resp, _, err := m.GetSpu().SpuLimitedDiscountAdd(&models.SpuLimitedDiscountRequest{
		ProductId: 16037732,
		StartTime: 1616986800,
		EndTime:   1617764400,
		LimitedDiscountSkuList: []*models.LimitedDiscountSkuList{
			&models.LimitedDiscountSkuList{
				SkuId:     18006617,
				SalePrice: 0,
				SaleStock: 12,
			},
		},
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestSpuLimitedGetList(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "43_qAxTOftN_eEqtdSoQnbuYjWDqorgys3phCVrHBl8_JPrJpFUVq6mr3vyMMSY1Hr5wecPkmfGkNRSCV--Y8_6j4dy8C1-rytrzCIxpPQOPaHmH2hIiFjE_yAr4XVYaXCKJq9t5gqPYoT4fm45GBHeAJDQLR",
	})
	resp, _, err := m.GetSpu().SpuLimitedGetList(&models.SpuLimitedListRequest{})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestSpuLimitedDiscountUpdate(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "43_qAxTOftN_eEqtdSoQnbuYjWDqorgys3phCVrHBl8_JPrJpFUVq6mr3vyMMSY1Hr5wecPkmfGkNRSCV--Y8_6j4dy8C1-rytrzCIxpPQOPaHmH2hIiFjE_yAr4XVYaXCKJq9t5gqPYoT4fm45GBHeAJDQLR",
	})
	resp, _, err := m.GetSpu().SpuLimitedDiscountUpdate(&models.SpuLimitedDiscountUpdateRequest{
		TaskId: 377486,
		Status: 1,
	})
	fmt.Println(util.JsonEncode(resp), err)
}
