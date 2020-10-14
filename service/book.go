package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/samiksha-awachat/books-go/model"
)

const (
	booksURL = "https://gist.githubusercontent.com/samiksha-awachat/69c085f70040f80c0ed3229f08367ffe/raw/3d287fa1d778a21e2787930b807771270a83c1b2/books.json"
)

// GetBooks ...
func GetBooks() ([]model.Book, error) {
	response, err := http.Get(booksURL)
	if err != nil {
		fmt.Printf("error while fetching books: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error while reading books: %v", err)
		return nil, err
	}

	var books model.GetBooks
	err = json.Unmarshal(responseBody, &books)
	if err != nil {
		fmt.Printf("error while parsing books: %v", err)
		return nil, err
	}

	return books.Books, nil
}

// SortBooksByAuthor ...
func SortBooksByAuthor() ([]model.Book, error) {
	books, err := GetBooks()
	if err != nil {
		return nil, err
	}

	sort.Slice(books, func(i, j int) bool {
		return strings.Compare(books[i].Author, books[j].Author) <= 0
	})

	return books, nil
}

// SearchBooksByTitle ...
func SearchBooksByTitle(title string) ([]model.Book, error) {
	books, err := GetBooks()
	if err != nil {
		return nil, err
	}

	var filteredBooks []model.Book

	title = strings.ToLower(title)
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), title) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	return filteredBooks, nil
}
