package product

import (
	"database/sql"
	"fmt"
	"log"
)

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *getUpdateProductRequest) (int64, error)
	DeleteProduct(params *getDeleteProductRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) GetProductById(productId int) (*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
				 standard_cost, list_price, category FROM products WHERE id =?`
	row := repo.db.QueryRow(sql, productId)
	product := &Product{}
	err := row.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
	if err != nil {
		panic(err)
	}
	return product, err
}

func (repo *repository) GetProducts(params *getProductRequest) ([]*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
	standard_cost, list_price, category FROM products ORDER BY id LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}

	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = `SELECT COUNT(*) FROM products`
	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}

func (repo *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const sql = `INSERT INTO products(product_code, product_name,category, description,list_price,standard_cost) VALUES(?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice, params.StandardCost)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId() //el guin bajo significa que no le interesa majejar la excepcion
	return id, nil
}

func (repo *repository) UpdateProduct(params *getUpdateProductRequest) (int64, error) {
	const sql = `UPDATE products SET product_code = ? , product_name = ? ,category = ?, description = ?,list_price = ?,standard_cost= ? WHERE id = ?`
	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice, params.StandardCost, params.ID)
	if err != nil {
		log.Fatalf("Couldn't connect to the database: %v", err)
	}

	fmt.Printf("Updated %v row(s) successfully.\n", result)
	id, _ := result.RowsAffected() //el guin bajo significa que no le interesa majejar la excepcion
	return id, nil
}

func (repo *repository) DeleteProduct(params *getDeleteProductRequest) (int64, error) {
	const sql = `DELETE FROM products WHERE id= ?`
	result, err := repo.db.Exec(sql, params.id)
	if err != nil {
		log.Fatalf("Couldn't connect to the database: %v", err)
	}
	id, _ := result.RowsAffected() //el guin bajo significa que no le interesa majejar la excepcion
	return id, nil
}
