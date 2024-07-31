package validators

import (
	"bibliosphere_gin/domain"
	"errors"
	"regexp"
	"time"
)

type BookValidator interface {
	Validate(book *domain.Book) error
}

type bookValidator struct{}

func NewBookValidator() BookValidator {
	return &bookValidator{}
}

func (v *bookValidator) Validate(book *domain.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}
	if book.Author == "" {
		return errors.New("author is required")
	}
	if book.ISBN == "" {
		return errors.New("ISBN is required")
	}
	if book.Publisher == "" {
		return errors.New("publisher is required")
	}
	if book.PageCount <= 0 {
		return errors.New("page count must be a positive number")
	}

	if !isValidISBN(book.ISBN) {
		return errors.New("invalid ISBN format")
	}

	parsedDate, err := time.Parse("2006-01-02", book.PublicationYear)
	if err != nil {
		return errors.New("invalid publication year format")
	}

	if !isValidPublicationYear(parsedDate) {
		return errors.New("invalid publication year format")
	}

	return nil
}

func isValidISBN(isbn string) bool {
	const isbn13Pattern = `^\d{13}$`
	matched, _ := regexp.MatchString(isbn13Pattern, isbn)
	return matched
}

func isValidPublicationYear(publicationYear time.Time) bool {
	currentYear := time.Now().Year()
	upperBound := currentYear + 10 // Allow some leeway for future dates or reprints
	year := publicationYear.Year()
	return year <= upperBound
}
