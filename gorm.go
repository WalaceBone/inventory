package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"inventory/graph/model"
	"strings"
	"time"
)

type MySQLConfig struct {
	mysql.Config
}

type Database struct {
	DB  *gorm.DB
	cfg MySQLConfig
}

func NewDatabase(cfg MySQLConfig) (*Database, error) {

	// Update by default the configuration
	cfg.AllowNativePasswords = true
	cfg.ParseTime = true

	// Build DSN
	dsn := strings.TrimPrefix(cfg.FormatDSN(), "mysql://")

	// Open the database
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Update some configuration for mysql
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)

	// Migrate the schema
	db.AutoMigrate(
		&model.Ingredient{},
		&model.Recipe{},
	)

	return &Database{
		db,
		cfg,
	}, nil
}

// Close closes the database
func (d *Database) Close() error {
	return d.DB.Close()
}

/*type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}*/
