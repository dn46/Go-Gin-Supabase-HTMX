package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

type Books struct {
	Title  string  `form:"title"`
	Author string  `form:"author"`
	Price  float64 `form:"price"`
}

var (
	supabaseClient *supa.Client
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		supabaseUrl = os.Getenv("SUPABASE_URL")
		supabaseKey = os.Getenv("SUPABASE_KEY")
	)

	supabaseClient = supa.CreateClient(supabaseUrl, supabaseKey)

	if supabaseClient == nil {
		log.Fatal("Unable to connect to Supabase")
	}

	// handler for the endpoint path; the getBooks function handles requests to the /books endpoint path
	router := gin.Default()
	router.LoadHTMLGlob("ui/html/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Books",
		})
	})

	router.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", gin.H{
			"title": "Create a new book",
		})
	})

	router.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"title": "List of books",
		})
	})

	//router.GET("/create", createBookHandler)

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

func createBookHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}
