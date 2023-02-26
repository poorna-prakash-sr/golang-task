package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title      string `json:"title"`
	Director   string `json:"director"`
	Hero       string `json:"hero"`
	RelaseDate string `json:"reasedate"`
	Filename   string `json:"filename"`
}
