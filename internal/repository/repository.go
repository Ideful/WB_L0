package repository

import (
	"encoding/json"
	"fmt"
	"l0/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type MyDB struct {
	Db *sqlx.DB
}

const (
	ordertalbe    = "orders"
	deliverytable = "delivery_info"
	paymenttable  = "payment_info"
	itemstable    = "order_items"
)

type Config struct {
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func InitConfig() Config {
	viper.SetConfigFile("config/db.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	port := viper.GetString("Port")
	username := viper.GetString("Username")
	password := viper.GetString("Password")
	dbname := viper.GetString("DBName")
	sslmode := viper.GetString("SSLMode")
	return Config{port, username, password, dbname, sslmode}
}

func CreatePostgresDB() (*MyDB, error) {
	var cfg = InitConfig()
	var db MyDB
	var err error
	db.Db, err = sqlx.Open("postgres", fmt.Sprintf("port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Port,
		cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Db.Ping()
	if err != nil {
		return nil, err
	}
	return &db, nil
}

func (db MyDB) InsertQuery(o *models.Order) error {
	if err := db.orders_insert(o); err != nil {
		return err
	}
	if err := db.deliveries_insert(o); err != nil {
		return err
	}
	if err := db.payments_insert(o); err != nil {
		return err
	}
	if err := db.items_insert(o); err != nil {
		return err
	}
	return nil
}

func (db MyDB) orders_insert(o *models.Order) error {
	var query = `INSERT INTO orders (order_UID, track_number, entry, locale, Internal_Signature, Customer_ID,
		Delivery_Service, Shardkey, SM_ID, Date_Created, OOF_Shard)  
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, (to_timestamp($10, 'YYYY-MM-DD HH24:MI:SS')), $11) RETURNING id`
	row := db.Db.DB.QueryRow(query, o.OrderUID, o.TrackNumber, o.Entry, o.Locale, o.InternalSignature, o.CustomerID, o.DeliveryService,
		o.Shardkey, o.SMID, (o.DateCreated), o.OOFShard)
	if err := row.Scan(&o.ID); err != nil {
		return err
	}
	return nil
}

func (db MyDB) deliveries_insert(o *models.Order) error {
	query := `INSERT INTO deliveries (name, phone, zip, city, address, region,
		email) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err := db.Db.DB.Exec(query, o.Delivery.Name, o.Delivery.Phone, o.Delivery.Zip,
		o.Delivery.City, o.Delivery.Address, o.Delivery.Region, o.Delivery.Email); err != nil {
		return err
	}
	return nil
}

func (db MyDB) payments_insert(o *models.Order) error {
	query := `INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt,
		bank,delivery_cost,goods_total,custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := db.Db.DB.Exec(query, o.Payment.Transaction, o.Payment.RequestID, o.Payment.Currency, o.Payment.Provider, o.Payment.Amount, o.Payment.PaymentDt,
		o.Payment.Bank, o.Payment.DeliveryCost, o.Payment.GoodsTotal, o.Payment.CustomFee); err != nil {
		return err
	}
	return nil
}

func (db MyDB) items_insert(o *models.Order) error {
	for _, i := range o.Items {
		query := `INSERT INTO items (order_id, chrt_id, track_number, price, rid, name,
			sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
		if _, err := db.Db.DB.Exec(query, o.ID, i.ChrtID, i.TrackNumber, i.Price, i.RID, i.Name, i.Sale, i.Size, i.TotalPrice,
			i.NmID, i.Brand, i.Status); err != nil {
			return err
		}
	}
	return nil
}

func (db *MyDB) GetOrder(id int) ([]byte, error) {
	o := models.Order{}

	query := `SELECT * FROM orders WHERE id = $1`
	err := db.Db.Get(&o, query, id)
	if err != nil {
		return nil, err
	}

	query = `SELECT * FROM deliveries WHERE id = $1`
	err = db.Db.Get(&o.Delivery, query, id)
	if err != nil {
		return nil, err
	}

	query = `SELECT * FROM payments WHERE id = $1`
	err = db.Db.Get(&o.Payment, query, id)
	if err != nil {
		return nil, err
	}

	query = `SELECT * FROM items WHERE order_id = $1`
	err = db.Db.Select(&o.Items, query, id)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.MarshalIndent(o, "", "	")

	return jsonBytes, err
}

func (db *MyDB) GetOrdersAmount() (int, error) {
	var id int
	query := `SELECT COUNT (*) FROM orders`
	err := db.Db.Get(&id, query)
	if err != nil {
		return -1, err
	}
	return id, err
}
