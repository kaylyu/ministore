package delivery_test

import (
	"fmt"
	"github.com/kaylyu/ministore"
	"github.com/kaylyu/ministore/config"
	"github.com/kaylyu/ministore/models"
	"github.com/kaylyu/ministore/util"
	"testing"
)

func TestGetDeliveryCompanyList(t *testing.T) {
	m := ministore.New(&config.Config{
		LogPrefix:   "ministore",
		AccessToken: "40_IK3jwO1yAitZpPqnb1Tzp5Mdv_CylhWV4Ue3UalRgVMMUalPd4dB_O0A5ssHmVIg90__L77zOUNpji5BqLkLS23BSFH27uVOeqLXdQJgv72D0jRYEVCIxKXkn0t7KWAs8tXMQfnzc2alXYctZKRaAHDEPW",
	})
	resp, _, err := m.GetDelivery().GetDeliveryCompanyList(&models.DeliveryGetCompanyRequest{})
	fmt.Println(util.JsonEncode(resp), err)
}
