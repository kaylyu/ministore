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
		AccessToken: "40_8DHWqetYL-mSUlViAfIkmhvmqwhX94wSjBfpUlBnG3aqpb1-3jHZrW924Psgk7WTpTjqqTapq4BGXXfElVdMNLc6iwoyevPA2w-pXAZQAj85F48X4CiGm9Fg-SskNTNXcjVUi3kWua5EdwbBVBYgAMDEKH",
	})
	resp, err := m.GetSpu().GetSpuList(&models.SpuGetListRequest{
		NeedEditSpu: models.NeedEditSpuDraft,
		Status:      models.SpuStatusPutaway,
		Page:        1,
		PageSize:    10,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestUpSpu(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_8DHWqetYL-mSUlViAfIkmhvmqwhX94wSjBfpUlBnG3aqpb1-3jHZrW924Psgk7WTpTjqqTapq4BGXXfElVdMNLc6iwoyevPA2w-pXAZQAj85F48X4CiGm9Fg-SskNTNXcjVUi3kWua5EdwbBVBYgAMDEKH",
	})
	resp, err := m.GetSpu().UpSpu(&models.SpuUpRequest{
		ProductId: 7516256,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestDownSpu(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_8DHWqetYL-mSUlViAfIkmhvmqwhX94wSjBfpUlBnG3aqpb1-3jHZrW924Psgk7WTpTjqqTapq4BGXXfElVdMNLc6iwoyevPA2w-pXAZQAj85F48X4CiGm9Fg-SskNTNXcjVUi3kWua5EdwbBVBYgAMDEKH",
	})
	resp, err := m.GetSpu().DownSpu(&models.SpuDownRequest{
		ProductId: 7516256,
	})
	fmt.Println(util.JsonEncode(resp), err)
}
