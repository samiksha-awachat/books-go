package model

import "time"

// Book ...
type Book struct {
	Isbn        string    `json:"isbn,omitempty"`
	Title       string    `json:"title,omitempty"`
	Subtitle    string    `json:"subtitle,omitempty"`
	Author      string    `json:"author,omitempty"`
	Published   time.Time `json:"published,omitempty"`
	Publisher   string    `json:"publisher,omitempty"`
	Pages       int       `json:"pages,omitempty"`
	Description string    `json:"description,omitempty"`
	Website     string    `json:"website,omitempty"`
}

// GetBooks ...
type GetBooks struct {
	Books []Book `json:"books,omitempty"`
}
