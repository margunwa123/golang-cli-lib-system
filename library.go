package main

import (
	"bufio"
	"fmt"
	"library/constants"
	"os"
	"strings"
)

type Book struct {
	name string
	author string
}

func main() {
	fmt.Println("Welcome to the go library!");
	fmt.Println("This library purpose is for learning GOLANG as well to create a random project. ");
	fmt.Println("Basically, you can CRUD library in here (the data will only be saved while the program is running).");
	fmt.Println();
	help();
	mainLoop();
}

func mainLoop() {
	books := makeEmptyBookArray()
	command := "0"
	for command != "-1" && command != constants.EXIT {
		fmt.Println("")
		fmt.Println("Enter a commmand (or type \"help\" for more information) ")
		fmt.Print("> ")
		command = readLine()
		evaluateCommand(command, &books)
	}
}

func makeEmptyBookArray() []Book {
	return make([]Book, 0);
}

// Reset the database then add some placeholder books
func initBooks(books* []Book) {
	resetBooks(books)
	book1 := Book{name: "haha", author: "tunaktun"}
	book2 := Book{name: "hehe", author: "daler"}
	book3 := Book{name: "hoho", author: "pler"}
	*books = append(*books, book1)
	*books = append(*books, book2)
	*books = append(*books, book3)
}

func evaluateCommand(command string, books *[]Book) {
	switch command {
		case "-1", constants.EXIT:
			fmt.Println("Goodbye!");

		// HELP
		case "0", constants.HELP:
			help()

		// SEE ALL BOOK
		case "1", constants.SEE_ALL:
			seeAllBook(*books)

		// OPEN
		case "2", constants.OPEN:
			openBook(*books)

		// DELETE
		case "3", constants.DELETE:
			deleteBook(books)

		// CREATE
		case "4", constants.CREATE:
			createBook(books)

		// UPDATE
		case "5", constants.UPDATE:
			updateBook(books)

		// RESET
		case "6", constants.RESET:
			resetBooks(books)

		// HELP
		case "7", constants.INIT_BOOKS:
			initBooks(books)

		default:
			fmt.Println("That is not a command");
	}
}

func resetBooks(books *[]Book) {
	*books = makeEmptyBookArray()

	fmt.Println("Successfully resetted the database");
}

func createBook(books *[]Book) {
	bookname := promptBookName()
	author := promptBookAuthor()

	newBook := Book{name: bookname, author: author}
	
	*books = append(*books, newBook)

	fmt.Printf("Successfully created the book %s \n", author);
}

func updateBook(books *[]Book) {
	bookname := promptBookName()
	theBook, idx := findBook(*books, bookname)

	if(idx == -1) {
		fmt.Println("Book name not found");
		return;
	}

	fmt.Println("Enter new author name");
	fmt.Print("> ")
	newAuthor := readLine();

	(*books)[idx] = Book{name: theBook.name, author: newAuthor}
}

// returns book and the index
func findBook(books []Book, bookname string) (Book, int) {
	for idx, value := range books {
		if(value.name == bookname) {
			return value, idx
		}
	}

	defaultBook := Book{}
	return defaultBook, -1
}

func openBook(books []Book) {
	bookname := promptBookName()

	for _,value := range books {
		if(value.name == bookname) {
			fmt.Printf("Book %s, Author %s\n", value.name, value.author)
		}
	}
}

func deleteBook(books *[]Book) {
	bookname := promptBookName()
	for idx, value := range *books {
		if(value.name == bookname) {
			*books = append((*books)[:idx], (*books)[idx + 1:]...)
			fmt.Println("Successfully deleted book");
			return
		}
	}
	fmt.Println("Book not found");
}

func promptBookAuthor() string {
	fmt.Print("Enter the author\n> ")
	bookname := readLine()
	return bookname
}

func promptBookName() string {
	fmt.Print("Enter the bookname\n> ")
	bookname := readLine()
	return bookname
}

func seeAllBook(books []Book) {
	for _,value := range books {
		fmt.Printf("Book %s \n", value.name);
	}
}

func help() {
	lines := []string{
		"COMMANDS: ",
		"1. See all (see all)",
		"2. Book detail (open)",
		"3. Delete a book (delete)",
		"4. Create a book (create)",
		"5. Update a book (update)",
		"6. Reset book database (reset)",
		"7. Help (help)",
		"Execute a command by typing its number or the text in parentheses",
		"Exit the program by typing Exit (exit) or \"-1\"",
	}
	for _, value := range lines {
		fmt.Println(value);
	}
}

func readLine() string {
  reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	command = strings.TrimSuffix(command, "\n")
	return command
}