package dal

import (
	"goGrocery/db"
	. "goGrocery/entities"
)

func InsertGroceryList(owner string) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO grocery_list(owner, active) VALUES ($1, $2) RETURNING id`

	err = conn.QueryRow(sql, owner, true).Scan(&id)

	return
}

func SelectGroceryList(owner string) (gl GroceryList, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM grocery_list WHERE owner = $1 and active = true and deleted_at = null`

	err = conn.QueryRow(sql, owner).Scan(&gl.ID, &gl.Owner, &gl.Active, &gl.CreatedAt, &gl.UpdatedAt)

	return
}

func SelectSingleGroceryList(listID int64) (gl GroceryList, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM grocery_list WHERE id = $1`

	err = conn.QueryRow(sql, listID).Scan(&gl.ID, &gl.Owner, &gl.Active, &gl.CreatedAt, &gl.UpdatedAt)

	return
}

func SelectGroceryItem(groceryItemID int64, groceryListID int64) (gi GroceryItem, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM grocery_item WHERE id = $1 and grocery_list_id = $2 and deleted_at = null`

	err = conn.QueryRow(sql, groceryItemID, groceryListID).Scan(&gi.ID, &gi.GroceryListID, &gi.Name, &gi.Quantity, &gi.Checked, &gi.CreatedAt, &gi.UpdatedAt)

	return
}

func SelectGroceryItems(groceryListID int64) (gis []GroceryItem, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT * FROM grocery_item WHERE grocery_list_id = $1 and deleted_at = null`

	rows, err := conn.Query(sql, groceryListID)

	if err != nil {
		return
	}

	for rows.Next() {
		var gi GroceryItem

		err = rows.Scan(&gi.ID, &gi.GroceryListID, &gi.Name, &gi.Quantity, &gi.Checked, &gi.CreatedAt, &gi.UpdatedAt)

		if err != nil {
			continue
		}

		gis = append(gis, gi)
	}

	return
}

func InsertGroceryItem(groceryListID int64, name string, quantity int) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO grocery_item(grocery_list_id, name, quantity, checked) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, groceryListID, name, quantity, false).Scan(&id)

	return
}

func UpdateGroceryList(groceryListID int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `UPDATE grocery_list SET active = false WHERE id = $1`

	rows, err := conn.Exec(sql, groceryListID)

	if err != nil {
		return 0, err
	}

	return rows.RowsAffected()
}

func UpdateGroceryItem(groceryItemID int64, listID int64, groceryItemName string, quantity int) (gi GroceryItem, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `UPDATE grocery_item SET name = $1, quantity = $2 WHERE id = $3 and grocery_list_id = $4`

	_, err = conn.Exec(sql, groceryItemName, quantity, groceryItemID, listID)

	if err != nil {
		return
	}

	gi, err = SelectGroceryItem(groceryItemID, listID)

	return
}

func UpdateCheckGroceryItem(groceryItemID int64, listID int64) (gi GroceryItem, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `UPDATE grocery_item SET checked = true WHERE id = $1 and grocery_list_id = $2`

	_, err = conn.Exec(sql, groceryItemID, listID)

	if err != nil {
		return
	}

	gi, err = SelectGroceryItem(groceryItemID, listID)

	return
}

func DeleteGroceryList(groceryListID int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `UPDATE grocery_list SET deleted_at = NOW() WHERE id = $1`

	rows, err := conn.Exec(sql, groceryListID)

	if err != nil {
		return 0, err
	}

	return rows.RowsAffected()
}

func DeleteGroceryItem(groceryItemID int64, listID int64) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	sql := `UPDATE grocery_item SET deleted_at = NOW() WHERE id = $1 and grocery_list_id = $2`

	rows, err := conn.Exec(sql, groceryItemID, listID)

	if err != nil {
		return 0, err
	}

	return rows.RowsAffected()
}
