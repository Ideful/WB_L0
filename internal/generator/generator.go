package generator

import (
	"encoding/json"
	"fmt"
	"l0/internal/models"
	"math/rand"
	"regexp"
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

func randomEmail() string {
	next := rand.Intn(4)
	if next > 0 {
		return randomString(5) + "@mail.com"
	}
	return randomString(15)
}

func randomDelivery() models.Delivery {
	return models.Delivery{
		Name:    randomString(8),
		Phone:   fmt.Sprintf("+7999054%d", rand.Intn(10000)),
		Zip:     fmt.Sprintf("%05d", rand.Intn(100000)),
		City:    randomString(10),
		Address: randomString(15),
		Region:  randomString(10),
		Email:   randomEmail(),
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

func randomItem() []models.Item {
	len := rand.Intn(4) + 1
	a := make([]models.Item, len)
	for i := 0; i < len; i++ {
		a[i].ChrtID = rand.Intn(100)
		a[i].TrackNumber = randomString(8)
		a[i].Price = rand.Intn(50)
		a[i].RID = randomString(6)
		a[i].Name = randomString(10)
		a[i].Sale = rand.Intn(2)
		a[i].Size = randomString(5)
		a[i].TotalPrice = rand.Intn(100)
		a[i].NmID = rand.Intn(50)
		a[i].Brand = randomString(8)
		a[i].Status = rand.Intn(3)
	}
	return a
}
func randomOrder() models.Order {
	return models.Order{
		OrderUID:          randomString(16),
		TrackNumber:       randomString(10),
		Entry:             randomString(5),
		Delivery:          randomDelivery(),
		Payment:           randomPayment(),
		Items:             randomItem(),
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

func Valid(o models.Order) bool {
	d := o.Delivery

	if len(d.Phone) != 12 {
		return false
	}
	if !isValidEmail(d.Email) {
		return false
	}
	return true
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return true
	}
	return false
}
