package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type MovieData struct {
	MovieName    string
	Genre        string
	Rating       string
	Director     string
	Actor        string
	PeopleVote   string
	Year         string
	Hero_Rating  string
	movie_rating string
}

func PrintCSV(data [][]string) {
	// var shoppingList []MovieData
	t := table.NewWriter()
	for _, line := range data {
		take := strings.Split(line[0], "	")
		var copy MovieData
		for j, sol := range take {
			if j == 0 {
				copy.MovieName = sol
			} else if j == 1 {
				copy.Genre = sol
			} else if j == 2 {
				copy.Rating = sol
			} else if j == 3 {
				copy.Director = sol
			} else if j == 4 {
				copy.Actor = sol
			} else if j == 5 {
				copy.PeopleVote = sol
			} else if j == 6 {
				copy.Year = sol
			} else if j == 7 {
				copy.Hero_Rating = sol
			} else if j == 8 {
				copy.movie_rating = sol
			}

		}
		t.SetOutputMirror(os.Stdout)
		t.AppendRows([]table.Row{
			{copy.MovieName, copy.Genre, copy.Rating, copy.Director, copy.Actor, copy.PeopleVote, copy.Year, copy.Hero_Rating, copy.movie_rating},
		})
		t.AppendSeparator()

	}
	t.Render()
}

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	PrintCSV(data)

}
