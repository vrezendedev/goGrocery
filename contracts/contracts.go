package contracts

type (
	User struct {
		Name string `json:"name" validate:"required,min=1,max=30"`
	}

	FinishGroceryList struct {
		ListID int64 `json:"list_id" validate:"required"`
	}

	GroceryItem struct {
		ID       int64  `json:"id"`
		ListID   int64  `json:"list_id" validate:"required"`
		Name     string `json:"name" validate:"required,min=1"`
		Quantity int    `json:"quantity" validate:"required,min=1"`
	}

	CheckGroceryItem struct {
		ListID        int64 `json:"list_id" validate:"required"`
		GroceryItemID int64 `json:"grocery_item_id" validate:"required"`
	}

	EntityID struct {
		ID int64 `json:"id" validate:"required"`
	}

	EntitiesIDS struct {
		ChildID  int64 `json:"child_id" validate:"required"`
		ParentID int64 `json:"parent_id" validate:"required"`
	}

	ValidateErrors struct {
		Field string
		Value interface{}
		Tag   string
	}
)
