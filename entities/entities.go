package entities

import (
	"database/sql"
	"time"
)

type (
	GroceryList struct {
		ID        int64        `json:"id"`
		Owner     string       `json:"owner"`
		Active    bool         `json:"active"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		DeletedAt sql.NullTime `json:"deleted_at"`
	}

	GroceryItem struct {
		ID            int64        `json:"id"`
		GroceryListID int64        `json:"grocery_list_id"`
		Name          string       `json:"name"`
		Quantity      int          `json:"quantity"`
		Checked       bool         `json:"checked"`
		CreatedAt     time.Time    `json:"created_at"`
		UpdatedAt     time.Time    `json:"updated_at"`
		DeletedAt     sql.NullTime `json:"deleted_at"`
	}
)
