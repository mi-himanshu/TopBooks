package models

import (
    "errors"
)

type Book struct {
    Id int64 `json:id`
    Author string `json:author`
    Title string `json:title`
    Desc string `json:desc`
    Ratings float32 `json:ratings`
    Reviews int32 `json:reviews`
}

func GetAllBooks() []Book {

    var books = []Book {
        {
            Id: 1,
            Author: "J.K. Rowling",
            Title: "Harry Potter",
            Desc: "Magic school wizardry",
            Ratings: 4.8,
            Reviews: 50,
        },
    }

    return books
}

func GetBookByName(name string) (Book, error) {
    // If no name was given, return an error with message
    if name == "" {
        return Book{}, errors.New("a name is required")
    } else if name != "Harry Potter" {
        return Book{}, errors.New("book not in stock")
    }

    return GetAllBooks()[0], nil
}