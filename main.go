package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

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

	router := gin.Default()
	router.LoadHTMLGlob("ui/html/*.html")

	//Serve static files
	router.Static("/static", "./ui/static")

	// Define your routes here
	router.GET("/", indexHandler)
	router.GET("/create", createHandler)
	router.GET("/list", listHandler)
	router.POST("/delete/:isbn", deleteHandler)
	router.GET("/edit/:isbn", editHandler)
	router.POST("/update/:isbn", updateHandler)
	router.GET("/books", getBooksHandler)
	router.POST("/books", postBooksHandler)

	router.Run("localhost:8080")
}
