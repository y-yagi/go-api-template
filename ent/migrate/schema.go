// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthorsColumns holds the columns for the "authors" table.
	AuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AuthorsTable holds the schema information for the "authors" table.
	AuthorsTable = &schema.Table{
		Name:        "authors",
		Columns:     AuthorsColumns,
		PrimaryKey:  []*schema.Column{AuthorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:        "books",
		Columns:     BooksColumns,
		PrimaryKey:  []*schema.Column{BooksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// AuthorBooksColumns holds the columns for the "author_books" table.
	AuthorBooksColumns = []*schema.Column{
		{Name: "author_id", Type: field.TypeInt},
		{Name: "book_id", Type: field.TypeInt},
	}
	// AuthorBooksTable holds the schema information for the "author_books" table.
	AuthorBooksTable = &schema.Table{
		Name:       "author_books",
		Columns:    AuthorBooksColumns,
		PrimaryKey: []*schema.Column{AuthorBooksColumns[0], AuthorBooksColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "author_books_author_id",
				Columns:    []*schema.Column{AuthorBooksColumns[0]},
				RefColumns: []*schema.Column{AuthorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "author_books_book_id",
				Columns:    []*schema.Column{AuthorBooksColumns[1]},
				RefColumns: []*schema.Column{BooksColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthorsTable,
		BooksTable,
		AuthorBooksTable,
	}
)

func init() {
	AuthorBooksTable.ForeignKeys[0].RefTable = AuthorsTable
	AuthorBooksTable.ForeignKeys[1].RefTable = BooksTable
}
