package httphandlers

import (
	"encoding/json"
	"go-flight-search/internal/domain"
	"go-flight-search/pkg/errs"
	"go-flight-search/pkg/helper"
	"net/http"
)

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req FlightSearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.JSON(w, ctx, nil, errs.NewWithMessage(http.StatusBadRequest, "Invalid request payload"))
		return
	}

	resp, isCache, err := h.SearchFlightUseCase.Execute(ctx, domain.SearchQuery{
		Origin:        req.Origin,
		Destination:   req.Destination,
		DepartureDate: req.DepartureDate,
		ReturnDate:    req.ReturnDate,
		Passengers:    req.Passengers,
		CabinClass:    req.CabinClass,
	})
	if err != nil {
		helper.JSON(w, ctx, nil, errs.NewWithMessage(http.StatusInternalServerError, "SearchFlightUseCase failed"))
		return
	}
	helper.JSON(w, ctx, ToSearchResponse(*resp, isCache, req), nil)
}

func ToSearchResponse(result []domain.Flight, isCache bool, query FlightSearchRequest) FlightSearchResponse {
	resp := FlightSearchResponse{
		SearchCriteria: SearchCriteria{
			Origin:        query.Origin,
			Destination:   query.Destination,
			DepartureDate: query.DepartureDate,
			Passengers:    query.Passengers,
			CabinClass:    query.CabinClass,
		},
		Metadata: Metadata{
			TotalResults: len(result),
			CacheHit:     isCache,
		},
	}

	flights := make([]FlightItem, len(result))
	for i, f := range result {
		flights[i] = toFlightResponse(f)
	}

	resp.Flights = flights
	return resp
}

func toFlightResponse(f domain.Flight) FlightItem {
	// Generate stable ID
	// Nullable aircraft
	var aircraft *string
	if f.Aircraft != nil {
		aircraft = f.Aircraft
	}

	return FlightItem{
		ID:       f.ID,
		Provider: f.Provider,

		Airline: Airline{
			Name: f.AirlineName,
			Code: f.AirlineCode,
		},

		FlightNumber: f.FlightNumber,

		Departure: AirportTime{
			Airport:   f.DepartureAirport,
			City:      f.DepartureCity,
			Datetime:  f.DepartureTime,
			Timestamp: f.DepartureTS,
		},

		Arrival: AirportTime{
			Airport:   f.ArrivalAirport,
			City:      f.ArrivalCity,
			Datetime:  f.ArrivalTime,
			Timestamp: f.ArrivalTS,
		},

		Duration: Duration{
			TotalMinutes: f.DurationMinutes,
			Formatted:    helper.FormatDuration(f.DurationMinutes),
		},

		Stops: f.Stops,

		Price: Price{
			Amount:   f.PriceAmount,
			Currency: f.PriceCurrency,
		},

		AvailableSeats: f.AvailableSeats,
		CabinClass:     f.CabinClass,

		Aircraft:  aircraft,
		Amenities: f.Amenities,

		Baggage: Baggage{
			CarryOn: f.CarryOnBaggage,
			Checked: f.CheckedBaggage,
		},
	}
}
