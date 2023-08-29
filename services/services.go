package services

import (
	"database/sql"
	"fmt"

	dal "goGrocery/dal"
	. "goGrocery/entities"

	"github.com/gofiber/fiber/v2"
)

func GetActiveGroceryList(owner string) (gl GroceryList, err error) {
	gl, err = dal.SelectGroceryList(owner)

	if err != nil {
		if err == sql.ErrNoRows {
			return gl, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("The requested grocery list wasn't found %s", err.Error()),
			)
		} else {
			return gl, fiber.NewError(fiber.StatusInternalServerError,
				fmt.Sprintf("An error has occurred while searching for active grocery list %s", err.Error()),
			)
		}

	}

	return
}

func NewGroceryList(owner string) (gl GroceryList, err error) {
	_, err = dal.SelectGroceryList(owner)

	if err == nil {
		return gl, fiber.NewError(fiber.StatusForbidden,
			"User already have an active Grocery List, finish it or delete it to create a new one",
		)
	}

	_, err = dal.InsertGroceryList(owner)

	if err != nil {
		return gl, fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("An error has occurred while creating new grocery list %s", err.Error()),
		)
	}

	gl, err = dal.SelectGroceryList(owner)

	if err != nil {
		if err == sql.ErrNoRows {
			return gl, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("The recent created grocery list wasn't found %s", err.Error()),
			)
		} else {
			return gl, fiber.NewError(fiber.StatusInternalServerError,
				fmt.Sprintf("An error has occurred while searching the new grocery list %s", err.Error()),
			)
		}
	}

	return
}

func FinishActiveGroceryList(listID int64) (err error) {
	rows, err := dal.UpdateGroceryList(listID)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("An error has occurred while finishing the grocery list %s", err.Error()),
		)
	}

	if rows != 1 {
		return fiber.NewError(fiber.StatusInternalServerError,
			"Update instruction affected zero or more than one row",
		)
	}

	return
}

func DeleteGroceryList(listID int64) (err error) {
	rows, err := dal.DeleteGroceryList(listID)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("An error has occurred while deleting the grocery list %s", err.Error()),
		)
	}

	if rows != 1 {
		return fiber.NewError(fiber.StatusInternalServerError,
			"Delete for the grocery list instruction affected zero or more than one row",
		)
	}

	return
}

func GetGroceryItems(listID int64) (gis []GroceryItem, err error) {
	lst, err := dal.SelectSingleGroceryList(listID)

	if err != nil {
		return gis, fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("An error has occurred while searching for the active the list %s", err.Error()),
		)
	}

	if !lst.Active || !lst.DeletedAt.Valid {
		return gis, fiber.NewError(fiber.StatusForbidden,
			"Target list is not active or is deleted.",
		)
	}

	gis, err = dal.SelectGroceryItems(listID)

	if err != nil {
		return gis, fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("An error has occurred while searching for the grocery items from active the list %s", err.Error()),
		)
	}

	return
}

func NewGroceryItem(listID int64, name string, quantity int) (gi GroceryItem, err error) {
	id, err := dal.InsertGroceryItem(listID, name, quantity)

	if err != nil {
		return gi, fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("An error has occurred while inserting the new grocery item %s", err.Error()),
		)
	}

	gi, err = dal.SelectGroceryItem(id, listID)

	if err != nil {
		if err == sql.ErrNoRows {
			return gi, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("The recent created grocery item wasn't found %s", err.Error()),
			)
		} else {
			return gi, fiber.NewError(fiber.StatusInternalServerError,
				fmt.Sprintf("An error has occurred while searching the new grocery item %s", err.Error()),
			)
		}
	}

	return
}

func UpdateGroceryItem(groceryItemID int64, listID int64, name string, quantity int) (gi GroceryItem, err error) {
	gi, err = dal.UpdateGroceryItem(groceryItemID, listID, name, quantity)

	if err != nil {
		if err == sql.ErrNoRows {
			return gi, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("The recent updated grocery item wasn't found %s", err.Error()),
			)
		} else {
			return gi, fiber.NewError(fiber.StatusInternalServerError,
				fmt.Sprintf("An error has occurred while updating the grocery item %s", err.Error()),
			)
		}
	}

	return
}

func UpdateCheckGroceryItem(groceryItemID int64, listID int64) (gi GroceryItem, err error) {
	gi, err = dal.UpdateCheckGroceryItem(groceryItemID, listID)

	if err != nil {
		if err == sql.ErrNoRows {
			return gi, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("The recent updated grocery item wasn't found %s", err.Error()),
			)
		} else {
			return gi, fiber.NewError(fiber.StatusInternalServerError,
				fmt.Sprintf("An error has occurred while updating the check state of the grocery item %s", err.Error()),
			)
		}
	}

	return
}

func DeleteGroceryItem(groceryItemID int64, listID int64) (err error) {
	rows, err := dal.DeleteGroceryItem(groceryItemID, listID)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError,
			fmt.Sprintf("An error has occurred while deleting the grocery item %s", err.Error()),
		)
	}

	if rows != 1 {
		return fiber.NewError(fiber.StatusInternalServerError,
			"Delete instruction for the grocery item affected zero or more than one row",
		)
	}

	return
}
