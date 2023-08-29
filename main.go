package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"

	. "goGrocery/contracts"
	. "goGrocery/services"
)

type XValidator struct {
	Validator *validator.Validate
}

var validate = validator.New()

func (v XValidator) Validate(data interface{}) []ValidateErrors {
	validationErrors := []ValidateErrors{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {

			var elem ValidateErrors

			elem.Field = err.Field()
			elem.Value = err.Value()
			elem.Tag = err.Tag()

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func StringifyErrors(errs []ValidateErrors) string {
	msgs := make([]string, len(errs))

	for _, err := range errs {
		msgs = append(msgs, fmt.Sprintf(
			"[%s]: '%v' ... needs to implement '%s'",
			err.Field,
			err.Value,
			err.Tag,
		))
	}

	return strings.Join(msgs, ",")[1:]
}

func main() {
	vld := &XValidator{
		Validator: validate,
	}

	app := fiber.New()

	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.SendString("Running ⚡")
	})

	app.Get("/api/grocery-list/active", func(c *fiber.Ctx) error {
		qrs := c.Queries()

		if len(qrs) == 0 {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Expected Query params but received none.")
		}

		user := new(User)
		user.Name = c.Query("name")

		errs := vld.Validate(user)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		gl, err := GetActiveGroceryList(user.Name)

		if err != nil {
			return err
		}

		return c.JSON(gl)
	})

	app.Post("/api/grocery-list", func(c *fiber.Ctx) error {
		u := new(User)

		c.BodyParser(u)

		errs := vld.Validate(u)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		gl, err := NewGroceryList(u.Name)

		if err != nil {
			return err
		}

		return c.JSON(gl)
	})

	app.Put("/api/grocery-list/finish", func(c *fiber.Ctx) error {
		fgl := new(FinishGroceryList)

		c.BodyParser(fgl)

		errs := vld.Validate(fgl)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		err := FinishActiveGroceryList(fgl.ListID)

		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Delete("/api/grocery-list", func(c *fiber.Ctx) error {
		eid := new(EntityID)

		c.BodyParser(eid)

		errs := vld.Validate(eid)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		err := DeleteGroceryList(eid.ID)

		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/api/grocery-item", func(c *fiber.Ctx) error {

		qrs := c.Queries()

		if len(qrs) == 0 {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Expected Query params but received none.")
		}

		eid := new(EntityID)
		eid.ID, _ = strconv.ParseInt(c.Query("list_id"), 0, 64)

		errs := vld.Validate(eid)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		gis, err := GetGroceryItems(eid.ID)

		if err != nil {
			return err
		}

		return c.JSON(gis)
	})

	app.Post("/api/grocery-item", func(c *fiber.Ctx) error {

		gi := new(GroceryItem)

		c.BodyParser(gi)

		errs := vld.Validate(gi)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		grci, err := NewGroceryItem(gi.ListID, gi.Name, gi.Quantity)

		if err != nil {
			return err
		}

		return c.JSON(grci)
	})

	app.Put("/api/grocery-item", func(c *fiber.Ctx) error {
		cgi := new(GroceryItem)

		c.BodyParser(cgi)

		errs := vld.Validate(cgi)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		gi, err := UpdateGroceryItem(cgi.ID, cgi.ListID, cgi.Name, cgi.Quantity)

		if err != nil {
			return err
		}

		return c.JSON(gi)
	})

	app.Put("/api/grocery-item/check", func(c *fiber.Ctx) error {
		cgi := new(CheckGroceryItem)

		c.BodyParser(cgi)

		errs := vld.Validate(cgi)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		gi, err := UpdateCheckGroceryItem(cgi.GroceryItemID, cgi.ListID)

		if err != nil {
			return err
		}

		return c.JSON(gi)
	})

	app.Delete("/api/grocery-item", func(c *fiber.Ctx) error {
		eid := new(EntitiesIDS)

		c.BodyParser(eid)

		errs := vld.Validate(eid)

		if len(errs) > 0 {
			return fiber.NewError(fiber.StatusBadRequest, StringifyErrors(errs))
		}

		err := DeleteGroceryItem(eid.ChildID, eid.ParentID)

		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	})

	app.Listen(":3000")
}
