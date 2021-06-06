package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/y-yagi/go-api-template/database"
	"github.com/y-yagi/go-api-template/ent"
	"github.com/y-yagi/go-api-template/ent/author"
)

func GetAuthors(c *fiber.Ctx) error {
	authors, err := database.Client.Author.Query().Order(ent.Asc(author.FieldID)).All(c.UserContext())
	if err != nil {
		return err
	}
	return c.JSON(authors)
}

func GetAuthor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	author, err := database.Client.Author.Query().Where(author.ID(id)).Only(c.UserContext())
	if err != nil {
		return err
	}
	return c.JSON(author)
}

func CreateAuthor(c *fiber.Ctx) error {
	author := new(ent.Author)
	if err := c.BodyParser(&author); err != nil {
		return err
	}

	_, err := database.Client.Author.Create().SetName(author.Name).Save(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	author := new(ent.Author)
	if err := c.BodyParser(&author); err != nil {
		return err
	}

	fmt.Printf("%v\n", author)
	_, err := database.Client.Author.UpdateOne(author).SetName(author.Name).Save(c.UserContext())
	if err != nil {
		return err
	}

	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	return database.Client.Author.DeleteOneID(id).Exec(c.UserContext())
}
