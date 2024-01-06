package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

type Books struct {
	ID     int8    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
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

	// getBooks()

	// row := Books{
	// 	ID:     2,
	// 	Title:  "The Lord of the Rings",
	// 	Author: "J.R.R. Tolkien",
	// 	Price:  9.99,
	// }

	// var results []Books

	// err := supabase.DB.From("Books").Insert(row).Execute(&results)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(results)

	// handler for the endpoint path; the getBooks function handles requests to the /books endpoint path
	router := gin.Default()
	router.GET("/books", getBooksHandler)
	// router.POST("/books", postBooks)
	// router.GET("/books/:id", getBookByID)

	router.Run("localhost:8080")

}

func getBooksHandler(c *gin.Context) {
	books, err := getBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

func getBooks() ([]Books, error) {

	var results []Books
	err := supabaseClient.DB.From("Books").Select("*").Execute(&results)
	if err != nil {
		log.Fatal("error")
	}

	return results, nil
}
