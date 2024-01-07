package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

type Books struct {
	//ID     int8    `json:"id"`
	Title  string  `form:"title"`
	Author string  `form:"author"`
	Price  float64 `form:"price"`
}

var (
	supabaseUrl    = "https://xayhcgyazivmkvyxzoyz.supabase.co"
	supabaseKey    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InhheWhjZ3lheml2bWt2eXh6b3l6Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDQ1NTg4OTgsImV4cCI6MjAyMDEzNDg5OH0.z94D5ceoDvrdYJyjbJPBw8qi-2gsFQCSDwDZjB85Sok"
	supabaseClient *supa.Client
)

// func postBooks(c *gin.Context) {
// 	var newBook Books

// 	if error := c.BindJSON(&newBook); error != nil {
// 		return
// 	}

// 	books = append(books, newBook)
// 	c.IndentedJSON(http.StatusCreated, newBook)
// }

// locating the item trough the id
// func getBookByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range books {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

func main() {

	supabaseClient = supa.CreateClient(supabaseUrl, supabaseKey)

	if supabaseClient == nil {
		log.Fatal("Unable to connect to Supabase")
	}

	// handler for the endpoint path; the getBooks function handles requests to the /books endpoint path
	router := gin.Default()
	router.LoadHTMLGlob("*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/books", getBooksHandler)
	router.POST("/books", postBooksHandler)
	// router.GET("/books/:id", getBookByID)

	router.Run("localhost:8080")

}

func getBooksHandler(c *gin.Context) {

	searchTerm := c.Query("search")

	books, err := getBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch books"})
		return
	}

	if searchTerm != "" {
		books = filterBooks(books, searchTerm)
	}

	//c.JSON(http.StatusOK, books)
	c.HTML(http.StatusOK, "books.html", gin.H{"books": books})
}

func getBooks() ([]Books, error) {

	var results []Books
	err := supabaseClient.DB.From("Books").Select("*").Execute(&results)
	if err != nil {
		log.Fatal("error")
	}

	return results, nil
}

func insertBook(book Books) error {
	var results []Books
	err := supabaseClient.DB.From("Books").Insert(book).Execute(&results)
	if err != nil {
		log.Fatal("error inserting", err)
	}

	return nil
}

func postBooksHandler(c *gin.Context) {
	var newBook Books

	if error := c.ShouldBind(&newBook); error != nil {
		return
	}

	err := insertBook(newBook)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert book"})
		return
	}

	// Fetch the updated list of books after a new book is inserted
	books, err := getBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch books"})
		return
	}

	// Return the updated list of books to the client
	c.HTML(http.StatusOK, "index.html", gin.H{"books": books})
}

func filterBooks(books []Books, searchTerm string) []Books {
	var results []Books

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(searchTerm)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(searchTerm)) {
			results = append(results, book)
		}
	}

	return results
}
