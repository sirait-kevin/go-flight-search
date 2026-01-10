package domain

type SearchQuery struct {
	Origin        string
	Destination   string
	DepartureDate string
	ReturnDate    *string
	Passengers    int
	CabinClass    string
}
