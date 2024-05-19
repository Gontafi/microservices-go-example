package repository

import (
	"context"
	"database/sql"
	"orders/internal/models"
)

type Repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) (int64, error) {
	queryOrderCreate := `INSERT INTO orders(user_id, status, total_amount, created_at, updated_at)
				VALUES (?, ?, ?, ?, ?)`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, queryOrderCreate,
		order.UserID,
		order.Status,
		order.TotalAmount,
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		return 0, err
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	queryOrderItemCreate := `INSERT INTO order_items(order_id, product_id, quantity, unit_price, created_at, updated_at)
				VALUES (?, ?, ?, ?, ?, ?)`

	for _, item := range items {
		_, err := tx.ExecContext(ctx, queryOrderItemCreate,
			orderID,
			item.ProductID,
			item.Quantity,
			item.UnitPrice,
			item.CreatedAt,
			item.UpdatedAt,
		)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (r *Repo) GetOrderByID(ctx context.Context, orderID int64) (*models.Order, error) {
	query := `SELECT id, user_id, status, total_amount, created_at, updated_at
			  FROM orders
			  WHERE id = ?`

	order := models.Order{}
	err := r.db.QueryRowContext(ctx, query, orderID).Scan(
		&order.ID,
		&order.UserID,
		&order.Status,
		&order.TotalAmount,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *Repo) GetOrderItemsByOrderID(ctx context.Context, orderID int64) ([]models.OrderItem, error) {
	query := `SELECT id, order_id, product_id, quantity, unit_price, created_at, updated_at
			  FROM order_items
			  WHERE order_id = ?`

	rows, err := r.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		item := models.OrderItem{}
		err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.ProductID,
			&item.Quantity,
			&item.UnitPrice,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repo) UpdateOrderStatus(ctx context.Context, orderID int64, status string) error {
	query := `UPDATE orders
			  SET status = ?, updated_at = CURRENT_TIMESTAMP
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, status, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteOrderByID(ctx context.Context, orderID int64) error {
	query := `DELETE FROM orders
			  WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, orderID)
	if err != nil {
		return err
	}

	return nil
}
