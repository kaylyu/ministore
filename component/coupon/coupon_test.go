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
		AccessToken: "43_s-VfxF-OkMmygASaj-wqCFwIoSn2sor44GzPOs9zNdxAQBlvFA1MX1lA23eCrKKo4u-ecAIK1v5BKuVNspQCwVPN_Wj2BK5PX8wRtiL7W68Y-VXducQLzfs3FVBj6I98a4_pKpkc3t8GYse6BNYcAJDNZK",
	})
	resp, _, err := m.GetCoupon().GetCouponList(&models.CouponGetListRequest{
		StartCreateTime: "2020-12-20 12:05:25",
		EndCreateTime:   "2021-10-25 12:05:25",
		Page:            1,
		PageSize:        10,
		Status:          models.CouponStatusValid,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestUpdateCoupon(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_xKoMXQzhVQhEwBpZzqYt5UuPli1Zroc5zdd3LU8BK6cuWj_6IDDuswCaAzNfwD5OhQmWiU_VzyJBsNBP-66OshVzXVm4UsrV1-BHcfsaxzj0wp_9Q4FwT5vfL174Ns2LfKZ7zTgGZ1XcQcKYCDOgAKDKLW",
	})
	resp, _, err := m.GetCoupon().UpdateCoupon(&models.CouponUpdateRequest{})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestUpdateStatusCoupon(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "43_qAxTOftN_eEqtdSoQnbuYjWDqorgys3phCVrHBl8_JPrJpFUVq6mr3vyMMSY1Hr5wecPkmfGkNRSCV--Y8_6j4dy8C1-rytrzCIxpPQOPaHmH2hIiFjE_yAr4XVYaXCKJq9t5gqPYoT4fm45GBHeAJDQLR",
	})
	resp, _, err := m.GetCoupon().UpdateStatusCoupon(&models.CouponUpdateStatusRequest{
		CouponId: 13050417,
		Status:   2,
	})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetCoupon(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "43_qAxTOftN_eEqtdSoQnbuYjWDqorgys3phCVrHBl8_JPrJpFUVq6mr3vyMMSY1Hr5wecPkmfGkNRSCV--Y8_6j4dy8C1-rytrzCIxpPQOPaHmH2hIiFjE_yAr4XVYaXCKJq9t5gqPYoT4fm45GBHeAJDQLR",
		Debug:       true,
	})
	resp, _, err := m.GetCoupon().GetCoupon(&models.CouponGetRequest{
		CouponId: 13488154,
	})
	fmt.Println(util.JsonEncode(resp), err)
}
