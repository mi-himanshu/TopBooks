package tests

import (
    "testing"
    "topbooks/models"
)

func Test_GetBooksByName(t *testing.T) {
    book, err := models.GetBookByName("Harry Potter")
    if err != nil {
        t.Errorf("Call failed")
    }

    if book.Title != "Harry Potter" {
        t.Errorf("Book fetched wrongly")
    }
}

func Test_GetAllBooks(t *testing.T) {
    books := models.GetAllBooks()
    if len(books) != 1 {
        t.Fatalf("Failed to fetch books")
    }
}