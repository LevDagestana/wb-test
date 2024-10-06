package cache

import (
	"sync"
	"wb/models"
)

type cache struct {
	orders map[string]models.Order
	sync.RWMutex
}

var Cache = cache{
	orders: map[string]models.Order{},
}

func (c *cache) SetCache(order models.Order) {
	c.Lock()
	defer c.Unlock()
	c.orders[order.OrderUID] = order
}

func (c *cache) GetCache(orderUID string) (models.Order, bool) {

	c.RLock()
	defer c.RUnlock()
	order, found := c.orders[orderUID]

	return order, found
}
