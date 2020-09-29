package datastore

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/matt-FFFFFF/bookdata-api/loader"
)

type Books struct {
	Store *[]*loader.BookData `json:"store"`
}

// Initialize is the method used to populate the file datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {

	start := time.Now()
	books := []*loader.BookData{}
	csvFile, err := os.Open("assets/books.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// using gocsv
	// if err := gocsv.UnmarshalFile(csvFile, &books); err != nil {
	// 	panic(err)
	// }

	// manual deserialization
	r := csv.NewReader(csvFile)

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, &loader.BookData{
			BookID:        row[0],
			Title:         row[1],
			Authors:       row[2],
			AverageRating: atof64(row[3]),
			ISBN:          row[4],
			ISBN13:        row[5],
			LanguageCode:  row[6],
			NumPages:      atoi(row[7]),
			Ratings:       atoi(row[8]),
			Reviews:       atoi(row[9]),
		})

	}

	fmt.Println("execution time:", time.Since(start))

	b.Store = &books
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func atof64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}
