package phone_repos

import (
	"AISale/database"
	"errors"
	"regexp"
)

func Create(title string, authors string, isbn string, docPath string) (Book, error) {
	db := database.GetDB()

	var book Book
	book.Title = title
	book.ISBN = isbn
	book.BookPdfFile = docPath

	re := regexp.MustCompile(`[^\w\d.]+`)

	if authors != "" {
		book.Authors = append(book.Authors, Author{
			Name: re.ReplaceAllString(authors, ""),
		})
	}

	if err := db.Create(&book).Error; err != nil {
		return Book{}, errors.New("невозможно создать новый объект")
	}

	return book, nil
}

func CheckIfExist(title string, docPath string) bool {
	db := database.GetDB()
	var book Book
	field := "title = ?"
	value := title

	if title == "" {
		field = "book_pdf_file = ?"
		value = docPath
	}
	if err := db.Where(field, value).First(&book).Error; err != nil {
		return false
	}

	return true
}

func UpdateEmptyFields(title string, authors string, isbn string, docPath string) (Book, []string, error) {
	value := title
	field := "title"

	if title == "" {
		value = docPath
		field = "book_pdf_file = ?"
	}

	book, err := GetBy(value, field)
	if err != nil {
		return Book{}, []string{}, err
	}

	re := regexp.MustCompile(`[^\w\d.]+`)
	var replaced []string

	if book.Title == "" && title != "" {
		book.Title = title
		replaced = append(replaced, "title")
	}
	if book.ISBN == "" && isbn != "" {
		book.ISBN = isbn
		replaced = append(replaced, "isbn")
	}
	if book.BookPdfFile == "" && docPath != "" {
		book.BookPdfFile = docPath
		replaced = append(replaced, "book_pdf_file")
	}
	if len(book.Authors) == 0 && authors != "" {
		book.Authors = append(book.Authors, Author{
			Name: re.ReplaceAllString(authors, ""),
		})
		replaced = append(replaced, "author")
	}

	return book, replaced, nil
}

func GetBy(value string, field string) (Book, error) {
	db := database.GetDB()

	var book Book
	if err := db.First(&book, field, value).Error; err != nil {
		return Book{}, errors.New("невозможно найти объект")
	}

	return book, nil
}

func Update(book Book) error {
	db := database.GetDB()

	if err := db.Save(&book).Error; err != nil {
		return errors.New("невозможно сохранить объект")
	}

	return nil
}
