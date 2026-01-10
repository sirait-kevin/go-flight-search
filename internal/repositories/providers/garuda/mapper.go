package garuda

import (
	"fmt"
	"go-flight-search/internal/domain"
	"go-flight-search/pkg/helper"
)

// MapToDomain converts GarudaFlight DTO to domain.Flight
func MapToDomain(f GarudaFlight) domain.Flight {
	// Determine departure & arrival
	dep := f.Departure
	arr := f.Arrival

	// If has segments, override from segments
	if len(f.Segments) > 0 {
		dep = GarudaAirportTime{
			Airport: f.Segments[0].Departure.Airport,
			Time:    f.Segments[0].Departure.Time,
		}

		last := f.Segments[len(f.Segments)-1]
		arr = GarudaAirportTime{
			Airport: last.Arrival.Airport,
			Time:    last.Arrival.Time,
		}
	}

	carry, checked := formatBaggage(f.Baggage.CarryOn, f.Baggage.Checked)
	return domain.Flight{
		ID:           f.FlightID,
		Provider:     f.Airline,
		AirlineCode:  f.AirlineCode,
		AirlineName:  f.Airline,
		FlightNumber: f.FlightID,

		DepartureAirport: dep.Airport,
		DepartureCity:    f.Departure.City,
		DepartureTime:    dep.Time,
		DepartureTS:      helper.ParseRFC3339ToUnix(dep.Time),

		ArrivalAirport: arr.Airport,
		ArrivalCity:    f.Arrival.City,
		ArrivalTime:    arr.Time,
		ArrivalTS:      helper.ParseRFC3339ToUnix(arr.Time),

		DurationMinutes: f.DurationMinutes,
		Stops:           f.Stops,

		PriceAmount:   f.Price.Amount,
		PriceCurrency: f.Price.Currency,

		AvailableSeats: f.AvailableSeats,
		CabinClass:     f.FareClass,
		Aircraft:       &f.Aircraft,

		Amenities: f.Amenities,

		CarryOnBaggage: carry,
		CheckedBaggage: checked,
	}
}

func formatBaggage(carry int, checked int) (string, string) {
	// No baggage included at all
	if carry == 0 && checked == 0 {
		return "Additional fee", "Additional fee"
	}

	// Only cabin baggage
	if checked == 0 {
		if carry == 1 {
			return "Cabin baggage only", "Additional fee"
		}
		return fmt.Sprintf("%d cabin baggage", carry), "Additional fee"
	}

	// Both included
	return fmt.Sprintf("%d cabin baggage(s)", carry), fmt.Sprintf("%d cargo baggage(s)", checked)
}
