package ministore

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/component/coupon"
	"github.com/kaylyu/ministore/component/delivery"
	"github.com/kaylyu/ministore/component/order"
	"github.com/kaylyu/ministore/component/product"
	"github.com/kaylyu/ministore/component/service"
	"github.com/kaylyu/ministore/component/sku"
	"github.com/kaylyu/ministore/component/spu"
	"github.com/kaylyu/ministore/component/store"
	"github.com/kaylyu/ministore/config"
	"github.com/kaylyu/ministore/context"
	"github.com/kaylyu/ministore/util"
	"github.com/sirupsen/logrus"
	"os"
)

/*
MiniStore 实例
*/
type MiniStore struct {
	context *context.Context
	core    *component.Core
}

/*
创建小商店实例
*/
func New(config *config.Config) *MiniStore {
	ctx := &context.Context{
		Config: config,
		//Logger: log.New(os.Stdout, fmt.Sprintf("[%s] ", config.LogPrefix), log.LstdFlags|log.Llongfile),
		Logger: &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    &util.CustomerFormatter{Prefix: config.LogPrefix},
			Level:        logrus.DebugLevel,
			ExitFunc:     os.Exit,
			ReportCaller: true,
		},
	}
	return &MiniStore{ctx, component.NewCore(ctx)}
}

//接入商品前必须接口
func (m *MiniStore) GetProduct() *product.Product {
	return product.NewProduct(m.core)
}

//服务市场
func (m *MiniStore) GetService() *service.Service {
	return service.NewService(m.core)
}

//SPU
func (m *MiniStore) GetSpu() *spu.Spu {
	return spu.NewSpu(m.core)
}

//SKU
func (m *MiniStore) GetSku() *sku.Sku {
	return sku.NewSku(m.core)
}

//Order
func (m *MiniStore) GetOrder() *order.Order {
	return order.NewOrder(m.core)
}

//Delivery
func (m *MiniStore) GetDelivery() *delivery.Delivery {
	return delivery.NewDelivery(m.core)
}

//Coupon
func (m *MiniStore) GetCoupon() *coupon.Coupon {
	return coupon.NewCoupon(m.core)
}

//Store
func (m *MiniStore) GetStore() *store.Store {
	return store.NewStore(m.core)
}
