package generator

import (
	"encoding/json"
	"fmt"
	"l0/internal/models"
	"math/rand"
	"time"
)

func Generator() ([]byte, error) {
	randval := randomOrder()
	return json.Marshal(randval)
}

func randomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func randomDelivery() models.Delivery {
	return models.Delivery{
		Name:    randomString(8),
		Phone:   fmt.Sprintf("+7999%d", rand.Intn(10000)),
		Zip:     fmt.Sprintf("%05d", rand.Intn(100000)),
		City:    randomString(10),
		Address: randomString(15),
		Region:  randomString(10),
		Email:   randomString(8) + "@email.com",
	}
}

func randomPayment() models.Payment {
	return models.Payment{
		Transaction:  randomString(12),
		RequestID:    randomString(10),
		Currency:     "RUB",
		Provider:     randomString(8),
		Amount:       rand.Intn(100),
		PaymentDt:    time.Now().Unix(),
		Bank:         randomString(10),
		DeliveryCost: rand.Intn(20),
		GoodsTotal:   rand.Intn(100),
		CustomFee:    rand.Intn(10),
	}
}

func randomItem() models.Item {
	return models.Item{
		ChrtID:      rand.Intn(100),
		TrackNumber: randomString(8),
		Price:       rand.Intn(50),
		RID:         randomString(6),
		Name:        randomString(10),
		Sale:        rand.Intn(2),
		Size:        randomString(5),
		TotalPrice:  rand.Intn(100),
		NmID:        rand.Intn(50),
		Brand:       randomString(8),
		Status:      rand.Intn(3),
	}
}

func randomOrder() models.Order {
	return models.Order{
		OrderUID:          randomString(16),
		TrackNumber:       randomString(10),
		Entry:             randomString(5),
		Delivery:          randomDelivery(),
		Payment:           randomPayment(),
		Items:             []models.Item{randomItem(), randomItem()},
		Locale:            "en_US",
		InternalSignature: randomString(20),
		CustomerID:        randomString(8),
		DeliveryService:   randomString(10),
		Shardkey:          randomString(8),
		SMID:              rand.Intn(10),
		DateCreated:       time.Now(),
		OOFShard:          randomString(6),
	}
}
