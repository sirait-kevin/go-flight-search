package domain

type Flight struct {
	ID           string
	Provider     string
	AirlineCode  string
	AirlineName  string
	FlightNumber string

	DepartureAirport string
	DepartureCity    string
	DepartureTime    string
	DepartureTS      int64

	ArrivalAirport string
	ArrivalCity    string
	ArrivalTime    string
	ArrivalTS      int64

	DurationMinutes int
	Stops           int

	PriceAmount   int
	PriceCurrency string

	AvailableSeats int
	CabinClass     string
	Aircraft       *string

	Amenities []string

	CarryOnBaggage string
	CheckedBaggage string
}
