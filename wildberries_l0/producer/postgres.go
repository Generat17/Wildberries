package producer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	db *sql.DB
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var (
	once        sync.Once
	dbClient    *PostgresClient
	dbClientErr error
)

// NewPostgresClient установка соединения с Postgresql
func NewPostgresClient(config PostgresConfig) (*PostgresClient, error) {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.DBName)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			dbClientErr = err
			return
		}

		err = db.Ping()
		if err != nil {
			dbClientErr = err
			return
		}

		dbClient = &PostgresClient{
			db: db,
		}
	})

	if dbClientErr != nil {
		return nil, dbClientErr
	}

	return dbClient, nil
}

// Close закрытие соединения с Postgresql
func (c *PostgresClient) Close() {
	c.db.Close()
}

// InsertOrder вставляет Order в таблицу orders
func (c *PostgresClient) InsertOrder(order Order) error {
	query, values := generateInsertQuery("orders", order)

	_, err := c.db.Exec(query, values...)
	if err != nil {
		return err
	}

	return nil
}

// generateInsertQuery генерируем sql запрос для вставки order в таблицу
func generateInsertQuery(tableName string, entity interface{}) (string, []interface{}) {
	var columns []string
	var placeholders []string
	var values []interface{}

	// получаем тип и значение объекта
	entityType := reflect.TypeOf(entity)
	entityValue := reflect.ValueOf(entity)

	// Итерируемся по полям структуры
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldValue := entityValue.Field(i).Interface()

		// Пропускаем поля, которые не экспортируются или для которых для тега "db" установлено значение "-"
		if field.PkgPath != "" || field.Tag.Get("db") == "-" {
			continue
		}

		var insertValue interface{}
		if field.Tag.Get("marshal") != "" {
			insertValue, _ = json.Marshal(fieldValue)
		} else {
			insertValue = fieldValue
		}

		columns = append(columns, field.Tag.Get("db"))
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(columns)))
		values = append(values, insertValue)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	return query, values
}

// GetOrderFromPostgres получает Order по id
func (c *PostgresClient) GetOrderFromPostgres(id string) (*Order, error) {
	query := "SELECT * FROM orders WHERE order_uid = $1"
	row := c.db.QueryRow(query, id)

	order := &Order{}
	var delivery, payment, items []byte

	err := row.Scan(
		&order.OrderUID, &order.TrackNumber, &order.Entry,
		&delivery, &payment, &items,
		&order.Locale, &order.InternalSignature, &order.CustomerID,
		&order.DeliveryService, &order.ShardKey, &order.SMID,
		&order.DateCreated, &order.OOFShard,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(delivery, &order.Delivery)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(payment, &order.Payment)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(items, &order.Items)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetOrdersFromPostgres получает все order из таблицы orders
func (c *PostgresClient) GetOrdersFromPostgres() ([]Order, error) {
	query := "SELECT * FROM orders"
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	orders := []Order{}

	for rows.Next() {

		order := Order{}
		var delivery, payment, items []byte

		err := rows.Scan(
			&order.OrderUID, &order.TrackNumber, &order.Entry,
			&delivery, &payment, &items,
			&order.Locale, &order.InternalSignature, &order.CustomerID,
			&order.DeliveryService, &order.ShardKey, &order.SMID,
			&order.DateCreated, &order.OOFShard,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(delivery, &order.Delivery)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(payment, &order.Payment)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(items, &order.Items)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)

	}

	return orders, nil
}
