package main

import (
	"log"
	"strings"

	supa "github.com/nedpals/supabase-go"
)

func getBooks(client *supa.Client) ([]Books, error) {
	var results []Books
	err := client.DB.From("Books").Select("*").Execute(&results)
	if err != nil {
		log.Fatal("error")
	}
	return results, nil
}

func getBook(client *supa.Client, isbn string) (Books, error) {
	var results []Books
	err := client.DB.From("Books").Select("*").Eq("ISBN", isbn).Execute(&results)
	if err != nil {
		log.Fatal("error")
	}
	return results[0], nil
}

func insertBook(client *supa.Client, book Books) error {
	var results []Books
	err := client.DB.From("Books").Insert(book).Execute(&results)
	if err != nil {
		log.Fatal("error inserting", err)
	}
	return nil
}

func deleteBook(client *supa.Client, isbn string) error {
	var results []Books
	err := client.DB.From("Books").Delete().Eq("ISBN", isbn).Execute(&results)
	if err != nil {
		log.Fatal("error deleting", err)
	}
	return nil
}

func updateBook(client *supa.Client, isbn string, book Books) error {
	var results []Books
	err := client.DB.From("Books").Update(book).Eq("ISBN", isbn).Execute(&results)
	if err != nil {
		log.Fatal("error updating", err)
	}
	return nil
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
