package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/omarghandour/test1/db"
	"github.com/omarghandour/test1/models"
	"gorm.io/gorm"
)

// GetBooks godoc
// @Summary      List all books
// @Description  Get a list of all books in the database
// @Tags         books
// @Produce      json
// @Success      200  {array}   models.Book
// @Failure      500  {object}  map[string]string "Internal Server Error"
// @Router       /api/books [get]
func GetBooks(c fiber.Ctx) error {
	var books []models.Book
	if err := db.DB.Find(&books).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}
	return c.JSON(books)
}

// GetBook godoc
// @Summary      Get book by ID
// @Description  Retrieve a single book from the database by its ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      404  {object}  map[string]string "Book not found"
// @Failure      500  {object}  map[string]string "Internal Server Error"
// @Router       /api/books/{id} [get]
func GetBook(c fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := db.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	return c.JSON(book)
}

// CreateBook godoc
// @Summary      Create a new book
// @Description  Add a new book to the database
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.BookInput  true  "Book data"
// @Success      201   {object}  models.Book
// @Failure      400   {object}  map[string]string "Invalid request payload"
// @Failure      500   {object}  map[string]string "Internal Server Error"
// @Router       /api/books [post]
func CreateBook(c fiber.Ctx) error {
	var input models.BookInput
	if err := c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if input.Title == "" || input.Author == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and Author are required fields",
		})
	}

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
		Year:   input.Year,
	}

	if err := db.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create book",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

// UpdateBook godoc
// @Summary      Update a book by ID
// @Description  Update details of an existing book by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      int               true  "Book ID"
// @Param        book  body      models.BookInput  true  "Book data to update"
// @Success      200   {object}  models.Book
// @Failure      400   {object}  map[string]string "Invalid request payload"
// @Failure      404   {object}  map[string]string "Book not found"
// @Failure      500   {object}  map[string]string "Internal Server Error"
// @Router       /api/books/{id} [put]
func UpdateBook(c fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := db.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	var input models.BookInput
	if err := c.Bind().JSON(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if input.Title != "" {
		book.Title = input.Title
	}
	if input.Author != "" {
		book.Author = input.Author
	}
	if input.Year != 0 {
		book.Year = input.Year
	}

	if err := db.DB.Save(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update book",
		})
	}

	return c.JSON(book)
}

// DeleteBook godoc
// @Summary      Delete a book by ID
// @Description  Remove a book from the database by its ID
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  map[string]string "Book deleted successfully"
// @Failure      404  {object}  map[string]string "Book not found"
// @Failure      500  {object}  map[string]string "Internal Server Error"
// @Router       /api/books/{id} [delete]
func DeleteBook(c fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if err := db.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	if err := db.DB.Delete(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete book",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
