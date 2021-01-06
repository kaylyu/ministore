package store_test

import (
	"fmt"
	"github.com/kaylyu/ministore"
	"github.com/kaylyu/ministore/config"
	"github.com/kaylyu/ministore/models"
	"github.com/kaylyu/ministore/util"
	"testing"
)

func TestStoreRegister(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_xKoMXQzhVQhEwBpZzqYt5UuPli1Zroc5zdd3LU8BK6cuWj_6IDDuswCaAzNfwD5OhQmWiU_VzyJBsNBP-66OshVzXVm4UsrV1-BHcfsaxzj0wp_9Q4FwT5vfL174Ns2LfKZ7zTgGZ1XcQcKYCDOgAKDKLW",
	})
	resp, err := m.GetStore().StoreRegister(&models.StoreRegisterRequest{}, "123333")
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetStoreInfo(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_xKoMXQzhVQhEwBpZzqYt5UuPli1Zroc5zdd3LU8BK6cuWj_6IDDuswCaAzNfwD5OhQmWiU_VzyJBsNBP-66OshVzXVm4UsrV1-BHcfsaxzj0wp_9Q4FwT5vfL174Ns2LfKZ7zTgGZ1XcQcKYCDOgAKDKLW",
	})
	resp, err := m.GetStore().GetStoreInfo(&models.StoreGetInfoRequest{})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestRegisterCheckAuditStatus(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_Vc8cMTQq-2ZpxpCecjvsmRRBT7HXuHdgGPxk2L1WBEnTAQOwoK8lNRh3RvZ_p-LvepkimQfbvod6aebJjfIoJ7D34N0kLpyRp-tpV5DcOgzNq6_T1iH-OYGCB_1fKGbDP5cHeompH1uZhiryGNHaAEDZGA",
	})
	resp, err := m.GetStore().RegisterCheckAuditStatus(&models.StoreRegisterCheckAuditStatusRequest{})
	fmt.Println(util.JsonEncode(resp), err)
}
