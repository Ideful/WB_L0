package models

import "time"

type Delivery struct {
	ID       int    `db:"id" json:"id"`
	Order_ID int    `db:"order_id" json:"order_id"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Zip      string `json:"zip" db:"zip"`
	City     string `json:"city" db:"city"`
	Address  string `json:"address" db:"address"`
	Region   string `json:"region" db:"region"`
	Email    string `json:"email" db:"email"`
}

type Payment struct {
	Order_ID     int    `db:"order_id" json:"order_id"`
	ID           int    `db:"id" json:"id"`
	Transaction  string `db:"transaction" json:"transaction"`
	RequestID    string `db:"request_id" json:"request_id"`
	Currency     string `db:"currency" json:"currency"`
	Provider     string `db:"provider" json:"provider"`
	Amount       int    `db:"amount" json:"amount"`
	PaymentDt    int64  `db:"payment_dt" json:"payment_dt"`
	Bank         string `db:"bank" json:"bank"`
	DeliveryCost int    `db:"delivery_cost" json:"delivery_cost"`
	GoodsTotal   int    `db:"goods_total" json:"goods_total"`
	CustomFee    int    `db:"custom_fee" json:"custom_fee"`
}

type Item struct {
	Order_ID    int    `db:"order_id" json:"order_id"`
	ID          int    `db:"id" json:"id"`
	ChrtID      int    `db:"chrt_id" json:"chrt_id"`
	TrackNumber string `db:"track_number" json:"track_number"`
	Price       int    `db:"price" json:"price"`
	RID         string `db:"rid" json:"rid"`
	Name        string `db:"name" json:"name"`
	Sale        int    `db:"sale" json:"sale"`
	Size        string `db:"size" json:"size"`
	TotalPrice  int    `db:"total_price" json:"total_price"`
	NmID        int    `db:"nm_id" json:"nm_id"`
	Brand       string `db:"brand" json:"brand"`
	Status      int    `db:"status" json:"status"`
}

type Order struct {
	ID                int       `db:"id" json:"id"`
	OrderUID          string    `db:"order_uid" json:"order_uid"`
	TrackNumber       string    `db:"track_number" json:"track_number"`
	Entry             string    `db:"entry" json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `db:"locale" json:"locale"`
	InternalSignature string    `db:"internal_signature" json:"internal_signature"`
	CustomerID        string    `db:"customer_id" json:"customer_id"`
	DeliveryService   string    `db:"delivery_service" json:"delivery_service"`
	Shardkey          string    `db:"shardkey" json:"shardkey"`
	SMID              int       `db:"sm_id" json:"sm_id"`
	DateCreated       time.Time `db:"date_created" json:"date_created"`
	OOFShard          string    `db:"oof_shard" json:"oof_shard"`
}
