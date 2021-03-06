package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/samiksha-awachat/books-go/enums"
	"github.com/samiksha-awachat/books-go/model"
	"github.com/samiksha-awachat/books-go/service"
)

func main() {
	option, err := ShowOptions()
	if err != nil {
		fmt.Println(err)
	}

	switch option {
	case enums.GetBooks:
		books, _ := service.GetBooks()
		PrintBooks(books)
	case enums.SortByAuthor:
		books, _ := service.SortBooksByAuthor()
		PrintBooks(books)
	case enums.SearchByTitle:
		fmt.Println("Enter the title: ")
		title := ""
		fmt.Scanln(&title)
		books, _ := service.SearchBooksByTitle(title)
		PrintBooks(books)
	default:
		fmt.Println("Enter a valid choice")
	}
}

// ShowOptions ...
func ShowOptions() (enums.Option, error) {
	message := `Select from the below given options:
1. List of Books
2. Sort by Author
3. Search by Title
Input: `
	fmt.Printf("%s", message)

	var input string
	fmt.Scanln(&input)

	inputOption, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("error while parsing user input: %v", err)
		return 0, err
	}

	return enums.Option(inputOption), nil
}

// PrintBooks ...
func PrintBooks(books []model.Book) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "TITLE\tAUTHOR\tDESCRIPTION")
	for _, book := range books {
		fmt.Fprintf(w, "%s\t%s\t%s...\t", book.Title, book.Author, book.Description[0:50])
		fmt.Fprintln(w)
	}
	w.Flush()
}
