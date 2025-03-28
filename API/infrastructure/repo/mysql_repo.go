package repo

import (
	"ApiShortLong/domain/entities"
	"ApiShortLong/domain/repo"
	"database/sql"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) repo.ProductRepository {
	return &mysqlRepository{db: db}
}

// Optimización para AddProduct
func (r *mysqlRepository) AddProduct(product domain.Product) error {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    _, err := r.db.ExecContext(ctx,
        "INSERT INTO Producto (nombre, precio, codigo, descuento) VALUES (?, ?, ?, ?)",
        product.Nombre, product.Precio, product.Codigo, product.Descuento,
    )
    return err
}

// Optimización para GetLastAddedProducts
func (r *mysqlRepository) GetLastAddedProducts(limit int) ([]domain.Product, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    rows, err := r.db.QueryContext(ctx,
        "SELECT nombre, precio, codigo, descuento FROM Producto ORDER BY id DESC LIMIT ?",
        limit,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []domain.Product
    for rows.Next() {
        var p domain.Product
        if err := rows.Scan(&p.Nombre, &p.Precio, &p.Codigo, &p.Descuento); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}

func (r *mysqlRepository) CountProductsInDiscount() (int, error) {
	row := r.db.QueryRow("SELECT COUNT(*) FROM Producto WHERE descuento = true")

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}