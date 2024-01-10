package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
	})
}

func createHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"title":   "Create a new book",
		"message": "Book created successfully",
	})
}

func listHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{
		"title": "List of books",
	})
}

func deleteHandler(c *gin.Context) {
	isbn := c.Param("isbn")
	err := deleteBook(supabaseClient, isbn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete book"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/list")
	c.HTML(http.StatusOK, "index.html", gin.H{"message": "Book deleted successfully"})
}

func editHandler(c *gin.Context) {
	isbn := c.Param("isbn")
	book, err := getBook(supabaseClient, isbn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch book"})
		return
	}
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"title": "Edit book",
		"book":  book,
	})
}

func updateHandler(c *gin.Context) {
	isbn := c.Param("isbn")
	var newBook Books
	if error := c.ShouldBind(&newBook); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := updateBook(supabaseClient, isbn, newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update book"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/list")
}

func getBooksHandler(c *gin.Context) {
	searchTerm := c.Query("search")
	books, err := getBooks(supabaseClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch books"})
		return
	}
	if searchTerm != "" {
		books = filterBooks(books, searchTerm)
	}
	c.HTML(http.StatusOK, "books.html", gin.H{"books": books})
}

func postBooksHandler(c *gin.Context) {
	var newBook Books
	if error := c.ShouldBind(&newBook); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := insertBook(supabaseClient, newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert book"})
		return
	}
	books, err := getBooks(supabaseClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch books"})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"books": books})
}
