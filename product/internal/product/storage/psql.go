package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	"venvision/internal/config"
	"venvision/protofiles"
)

type PSQL struct {
	DB *sqlx.DB
}

// InsertOne inserts a single product into the database.
func (ps *PSQL) InsertOne(ctx context.Context, p *protofiles.Product) error {
	// Write the SQL query to insert a single product into the database.
	query := `INSERT INTO products (uuid, shop_uuid, name, description, price_value, price_currency_code, images_url, categories) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Convert the Price message to the corresponding database columns.
	priceValue := p.Price.GetValue()
	priceCurrencyCode := p.Price.GetCurrencyCode()

	// Execute the SQL query and return any error that occurs.
	_, err := ps.DB.ExecContext(ctx, query, p.Uuid, p.ShopUUID, p.Name, p.Description, priceValue, priceCurrencyCode,
		p.ImagesURL, p.Categories)
	return err
}

// FindByUUID finds a single product by its UUID.
func (ps *PSQL) FindByUUID(ctx context.Context, uuid string) (*protofiles.Product, error) {
	// Declare a variable to hold the product that will be returned.
	var product protofiles.Product

	// Write the SQL query to find a single product by its UUID.
	query := `SELECT * FROM products WHERE uuid = $1`

	// Execute the SQL query and scan the results into the product variable.
	err := ps.DB.GetContext(ctx, &product, query, uuid)
	if err != nil {
		return &protofiles.Product{}, err
	}

	// Convert the database columns to the corresponding message fields.
	price := &protofiles.Price{
		Value:        product.Price.Value,
		CurrencyCode: product.Price.CurrencyCode,
	}

	// Convert the string arrays to repeated fields.
	var imagesURL []string
	var categories []string

	// Set the converted fields on the product message.
	product.Price = price
	product.ImagesURL = imagesURL
	product.Categories = categories

	// Return the product and any error that occurred.
	return &product, nil
}

// FindAll finds all products in the database.
func (ps *PSQL) FindAll(ctx context.Context) ([]*protofiles.Product, error) {
	// Declare a variable to hold the slice of products that will be returned.
	var products []*protofiles.Product

	// Write the SQL query to find all products in the database.
	query := `SELECT * FROM products`

	// Execute the SQL query and scan the results into the products slice.
	err := ps.DB.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}

	// Convert the database columns to the corresponding message fields for each product.
	for i := range products {
		price := &protofiles.Price{
			Value:        products[i].Price.Value,
			CurrencyCode: products[i].Price.CurrencyCode,
		}

		var imagesURL []string
		var categories []string

		products[i].Price = price
		products[i].ImagesURL = imagesURL
		products[i].Categories = categories
	}

	// Return the products slice and any error that occurred.
	return products, nil
}

// Update updates a single product in the database.
func (ps *PSQL) Update(ctx context.Context, uuid string, p *protofiles.Product) error {
	// Write the SQL query to update a single product in the database.
	query := `UPDATE products SET shop_uuid = $1, name = $2, description = $3, price_value = $4, price_currency_code = $5, images_url = $6, categories = $7 WHERE uuid = $8`
	// Convert the Price message to the corresponding database columns.
	priceValue := p.Price.GetValue()
	priceCurrencyCode := p.Price.GetCurrencyCode()

	// Execute the SQL query and return any error that occurs.
	_, err := ps.DB.ExecContext(ctx, query, p.ShopUUID, p.Name, p.Description, priceValue, priceCurrencyCode, p.ImagesURL, p.Categories, uuid)
	return err
}

// Delete deletes a single product from the database.
func (ps *PSQL) Delete(ctx context.Context, uuid string) error {
	// Write the SQL query to delete a single product from the database.
	query := `DELETE FROM products WHERE uuid = $1`

	// Execute the SQL query and return any error that occurs.
	_, err := ps.DB.ExecContext(ctx, query, uuid)
	return err
}

// DeleteAll deletes all products from the database.
func (ps *PSQL) DeleteAll(ctx context.Context) error {
	// Write the SQL query to delete all products from the database.
	query := `DELETE FROM products`
	// Execute the SQL query and return any error that occurs.
	_, err := ps.DB.ExecContext(ctx, query)
	return err
}

func NewPSQLClient(ctx context.Context, attempts int, c config.StorageConfig) (s Storage, err error) {
	var db *sqlx.DB
	for i := 0; i < attempts; i++ {
		db, err = sqlx.ConnectContext(ctx, "postgres", c.Database)
		if err == nil {
			break
		}
	}

	if err != nil {
		return
	}

	s = &PSQL{
		DB: db,
	}
	return
}
