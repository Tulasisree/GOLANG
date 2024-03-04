package media

import "strings"

type Catalogable interface {
	NewMovie(title string, rating Rating, boxoffice float32)
	GetTitle() string
	GetBoxOffice() string
	GetRating() string
	SetTitle(newTitle string)
	SetBoxOffice(newBoxoffice float32)
	SetRating(newRating string)
}

type Movie struct {
	title     string
	rating    Rating
	boxoffice float32
}

type Rating string

const (
	R    = "R (Restricted)"
	G    = "G (Genereal audiences)"
	PG   = "PG (Parental Guidance)"
	PG13 = "PG-13 (Parental Caution)"
	NC17 = "NC-17 (No children under 17)"
)

func (m *Movie) NewMovie(title string, rating Rating, boxoffice float32) {

	m.title = title
	m.rating = rating
	m.boxoffice = boxoffice

}

func (m *Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m *Movie) GetRating() Rating {
	return m.rating
}

func (m *Movie) GetBoxOffice() float32 {
	return m.boxoffice
}

func (m *Movie) SetTitle(title string) {
	m.title = title
}

func (m *Movie) SetRating(rating Rating) {
	m.rating = rating
}

func (m *Movie) SetBoxOffice(boxoffice float32) {
	m.boxoffice = boxoffice
}
