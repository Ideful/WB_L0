package models

import (
	"encoding/json"
	"errors"

	"l0/internal/models"
	"l0/internal/repository"
	"sync"
)

type Cache struct {
	sync.RWMutex
	orders map[int]models.Order
	pod_id int
	i_id   int
}

func NewCache() *Cache {
	orders := make(map[int]models.Order)

	cache := Cache{
		orders: orders,
		pod_id: 1,
		i_id:   1,
	}
	return &cache
}

func (c *Cache) AddToCache(order *models.Order) {
	c.Lock()
	defer c.Unlock()

	order.ID = c.pod_id
	order.Delivery.ID = c.pod_id
	order.Payment.ID = c.pod_id
	for i := range order.Items {
		order.Items[i].Order_ID = c.i_id
		c.i_id++
	}

	c.orders[c.pod_id] = *order
	c.pod_id++
}

func (c *Cache) GetFromCache(key int) ([]byte, error) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.orders[key]
	if !found {
		return nil, errors.ErrUnsupported
	}

	val, err := json.MarshalIndent(item, "", "	")
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *Cache) FillCache(db *repository.MyDB) error {
	max_id, err := db.GetOrdersAmount()
	if err != nil {
		return err
	}
	for i := 1; i < max_id; i++ {
		v, err := db.GetOrder(i)
		if err != nil {
			return err
		}
		order := models.Order{}
		json.Unmarshal(v, &order)
		c.orders[i] = order
	}
	c.i_id = (max_id + 1) * 2
	c.pod_id = (max_id + 1)
	return nil
}
