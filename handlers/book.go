package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/y-yagi/go-api-template/database"
	"github.com/y-yagi/go-api-template/ent"
	"github.com/y-yagi/go-api-template/ent/book"
)

func GetBooks(c *fiber.Ctx) error {
	books, err := database.Client.Book.Query().Order(ent.Asc(book.FieldID)).All(c.UserContext())
	if err != nil {
		return err
	}
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	book, err := database.Client.Book.Query().Where(book.ID(id)).Only(c.UserContext())
	if err != nil {
		return err
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(ent.Book)
	if err := c.BodyParser(&book); err != nil {
		return err
	}

	_, err := database.Client.Book.Create().SetName(book.Name).SetAuthor(book.Author).Save(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	book := new(ent.Book)
	if err := c.BodyParser(&book); err != nil {
		return err
	}

	_, err := database.Client.Book.UpdateOne(book).Save(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	return database.Client.Book.DeleteOneID(id).Exec(c.UserContext())
}
