package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "Rich Dad Poor Dad",
		Author:        "ABC1",
		YearPublished: 2019,
	}, {
		ID:            2,
		Title:         "Rich Dad Poor Dad 2",
		Author:        "ABC2",
		YearPublished: 2020,
	},
	{
		ID:            3,
		Title:         "Rich Dad Poor Dad 3",
		Author:        "ABC3",
		YearPublished: 2023,
	},
	{
		ID:            4,
		Title:         "Rich Dad Poor Dad 4",
		Author:        "ABC4",
		YearPublished: 2024,
	},
	{
		ID:            5,
		Title:         "Rich Dad Poor Dad 5",
		Author:        "ABC5",
		YearPublished: 2025,
	},
	{
		ID:            6,
		Title:         "Rich Dad Poor Dad 6",
		Author:        "ABC6",
		YearPublished: 2026,
	},
	{
		ID:            7,
		Title:         "Rich Dad Poor Dad 7",
		Author:        "ABC7",
		YearPublished: 2027,
	},
	{
		ID:            8,
		Title:         "Rich Dad Poor Dad 8",
		Author:        "ABC8",
		YearPublished: 2028,
	},
	{
		ID:            9,
		Title:         "Rich Dad Poor Dad 9",
		Author:        "ABC9",
		YearPublished: 2029,
	},
	{
		ID:            10,
		Title:         "Rich Dad Poor Dad 10",
		Author:        "ABC10",
		YearPublished: 2030,
	},
}

// func main() {

// 	for _, b := range books {
// 		result := b.String()
// 		fmt.Println(result)
// 	}

// }
