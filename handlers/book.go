package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/y-yagi/go-api-template/database"
	"github.com/y-yagi/go-api-template/ent"
	"github.com/y-yagi/go-api-template/ent/author"
	"github.com/y-yagi/go-api-template/ent/book"
)

type BookParam struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	AuthorID int    `json:"author_id,omitempty"`
}

func GetBooks(c *fiber.Ctx) error {
	books, err := database.Client.Book.Query().Order(ent.Asc(book.FieldID)).WithAuthor().All(c.UserContext())
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

	book, err := database.Client.Book.Query().Where(book.ID(id)).WithAuthor().Only(c.UserContext())
	if err != nil {
		return err
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	params := new(BookParam)
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	author, err := database.Client.Author.Query().Where(author.ID(params.AuthorID)).Only(c.UserContext())
	if err != nil {
		return err
	}

	book, err := database.Client.Book.Create().SetName(params.Name).AddAuthor(author).Save(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	params := new(BookParam)
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	book, err := database.Client.Book.UpdateOneID(params.ID).SetName(params.Name).Save(c.UserContext())
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
