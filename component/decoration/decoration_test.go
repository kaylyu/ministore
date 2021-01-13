package decoration_test

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
	resp, _, err := m.GetDecoration().DecorationSwitch(&models.DecorationSwitchRequest{})
	fmt.Println(util.JsonEncode(resp), err)
}

func TestGetCouponList(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_xKoMXQzhVQhEwBpZzqYt5UuPli1Zroc5zdd3LU8BK6cuWj_6IDDuswCaAzNfwD5OhQmWiU_VzyJBsNBP-66OshVzXVm4UsrV1-BHcfsaxzj0wp_9Q4FwT5vfL174Ns2LfKZ7zTgGZ1XcQcKYCDOgAKDKLW",
	})
	resp, _, err := m.GetDecoration().DecorationExperience(&models.DecorationExperienceRequest{})
	fmt.Println(util.JsonEncode(resp), err)
}
