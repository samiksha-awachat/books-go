package service

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"testing"
)

func Test_GetBooks_Success(t *testing.T) {
	// Arrange
	fetchBooks = func() (*http.Response, error) {
		body := getMockBooksJSON()
		response := &http.Response{
			Status:        "200 OK",
			StatusCode:    200,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
			ContentLength: int64(len(body)),
			// Request:       req,
			Header: make(http.Header, 0),
		}
		return response, nil
	}

	// Act
	books, err := GetBooks()

	// Assert
	if err != nil {
		t.Errorf("expected error nil, got error = %v", err)
		return
	}
	if len(books) != 3 {
		t.Errorf("expected 3 books, got %d books", len(books))
		return
	}
}

func Test_GetBooks_Error(t *testing.T) {
	// Arrange
	fetchBooks = func() (*http.Response, error) {
		return nil, errors.New("network error")
	}

	// Act
	books, err := GetBooks()

	// Assert
	if err == nil {
		t.Errorf("expected error, got nil error")
		return
	}
	if err.Error() != "network error" {
		t.Errorf("expected `network error` error, got error = %v", err.Error())
		return
	}
	if len(books) != 0 {
		t.Errorf("expected 0 books, got %d books", len(books))
		return
	}
}

func Test_SearchBooksByTitle_Success(t *testing.T) {
	// Arrange
	fetchBooks = func() (*http.Response, error) {
		body := getMockBooksJSON()
		response := &http.Response{
			Status:        "200 OK",
			StatusCode:    200,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
			ContentLength: int64(len(body)),
			// Request:       req,
			Header: make(http.Header, 0),
		}
		return response, nil
	}

	// Act
	books, err := SearchBooksByTitle("javascript")

	// Assert
	if err != nil {
		t.Errorf("expected error not nil, got error = %v", err)
		return
	}
	if len(books) != 2 {
		t.Errorf("expected 2 books, got %d books", len(books))
		return
	}
}

func Test_SearchBooksByTitle_Error(t *testing.T) {
	// Arrange
	fetchBooks = func() (*http.Response, error) {
		return nil, errors.New("network error")
	}

	// Act
	books, err := SearchBooksByTitle("javascript")

	// Assert
	if err == nil {
		t.Errorf("expected error, got nil error")
		return
	}
	if err.Error() != "network error" {
		t.Errorf("expected `network error` error, got error = %v", err.Error())
		return
	}
	if len(books) != 0 {
		t.Errorf("expected 0 books, got %d books", len(books))
		return
	}
}

func Test_SortBooksByAuthor_Success(t *testing.T) {
	// Arrange
	fetchBooks = func() (*http.Response, error) {
		body := getMockBooksJSON()
		response := &http.Response{
			Status:        "200 OK",
			StatusCode:    200,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
			ContentLength: int64(len(body)),
			// Request:       req,
			Header: make(http.Header, 0),
		}
		return response, nil
	}

	// Act
	books, err := SortBooksByAuthor()

	// Assert
	if err != nil {
		t.Errorf("expected error not nil, got error = %v", err)
		return
	}
	if len(books) != 3 {
		t.Errorf("expected 3 books, got %d books", len(books))
		return
	}

	if !sort.SliceIsSorted(books, func(i, j int) bool {
		return strings.Compare(books[i].Author, books[j].Author) <= 0
	}) {
		t.Errorf("expected books to be sorted by author, got unsorted books")
		return
	}
}

func Test_SortBooksByAuthor_Error(t *testing.T) {
	// Arrange
	fetchBooks = func() (*http.Response, error) {
		return nil, errors.New("network error")
	}

	// Act
	books, err := SortBooksByAuthor()

	// Assert
	if err == nil {
		t.Errorf("expected error, got nil error")
		return
	}
	if err.Error() != "network error" {
		t.Errorf("expected `network error` error, got error = %v", err.Error())
		return
	}
	if len(books) != 0 {
		t.Errorf("expected 0 books, got %d books", len(books))
		return
	}
}

func getMockBooksJSON() string {
	return `
	{
		"books": [
		  {
			"isbn": "9781593275846",
			"title": "Eloquent JavaScript, Second Edition",
			"subtitle": "A Modern Introduction to Programming",
			"author": "Marijn Haverbeke",
			"published": "2014-12-14T00:00:00.000Z",
			"publisher": "No Starch Press",
			"pages": 472,
			"description": "JavaScript lies at the heart of almost every modern web application, from social apps to the newest browser-based games. Though simple for beginners to pick up and play with, JavaScript is a flexible, complex language that you can use to build full-scale applications.",
			"website": "http://eloquentjavascript.net/"
		  },
		  {
			"isbn": "9781449331818",
			"title": "Learning JavaScript Design Patterns",
			"subtitle": "A JavaScript and jQuery Developer's Guide",
			"author": "Addy Osmani",
			"published": "2012-07-01T00:00:00.000Z",
			"publisher": "O'Reilly Media",
			"pages": 254,
			"description": "With Learning JavaScript Design Patterns, you'll learn how to write beautiful, structured, and maintainable JavaScript by applying classical and modern design patterns to the language. If you want to keep your code efficient, more manageable, and up-to-date with the latest best practices, this book is for you.",
			"website": "http://www.addyosmani.com/resources/essentialjsdesignpatterns/book/"
		  },
		  {
			"isbn": "9781449325862",
			"title": "Git Pocket Guide",
			"subtitle": "A Working Introduction",
			"author": "Richard E. Silverman",
			"published": "2013-08-02T00:00:00.000Z",
			"publisher": "O'Reilly Media",
			"pages": 234,
			"description": "This pocket guide is the perfect on-the-job companion to Git, the distributed version control system. It provides a compact, readable introduction to Git for new users, as well as a reference to common commands and procedures for those of you with Git experience.",
			"website": "http://chimera.labs.oreilly.com/books/1230000000561/index.html"
		  }
		]
	}
	`
}
