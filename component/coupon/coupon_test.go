package coupon_test

import (
	"fmt"
	"github.com/kaylyu/ministore"
	"github.com/kaylyu/ministore/config"
	"github.com/kaylyu/ministore/models"
	"github.com/kaylyu/ministore/util"
	"testing"
)

func TestCreateCoupon(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_xKoMXQzhVQhEwBpZzqYt5UuPli1Zroc5zdd3LU8BK6cuWj_6IDDuswCaAzNfwD5OhQmWiU_VzyJBsNBP-66OshVzXVm4UsrV1-BHcfsaxzj0wp_9Q4FwT5vfL174Ns2LfKZ7zTgGZ1XcQcKYCDOgAKDKLW",
	})
	resp, _, err := m.GetCoupon().CreateCoupon(&models.CouponCreateRequest{
		Type: models.CouponTypeShopFull,
		CouponInfo: &models.CouponInfo{
			Name: "测试优惠券kaylv",
			DiscountInfo: &models.DiscountInfo{
				DiscountCondition: &models.DiscountCondition{
					ProductCnt:   2,
					ProductIds:   []uint64{7516256},
					ProductPrice: 50,
				},
				DiscountFee: 20,
				DiscountNum: 10,
			},
		},
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetCouponList(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_xKoMXQzhVQhEwBpZzqYt5UuPli1Zroc5zdd3LU8BK6cuWj_6IDDuswCaAzNfwD5OhQmWiU_VzyJBsNBP-66OshVzXVm4UsrV1-BHcfsaxzj0wp_9Q4FwT5vfL174Ns2LfKZ7zTgGZ1XcQcKYCDOgAKDKLW",
	})
	resp, _, err := m.GetCoupon().GetCouponList(&models.CouponGetListRequest{
		StartCreateTime: "2020-12-20 12:05:25",
		EndCreateTime:   "2021-2-25 12:05:25",
		Page:            1,
		PageSize:        10,
		Status:          models.CouponStatusValid,
	})
	fmt.Println(util.JsonEncode(resp), err)
}
