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

	resp, _, err := h.SearchFlightUseCase.Execute(ctx, domain.SearchQuery{
		Origin:        req.Origin,
		Destination:   req.Destination,
		DepartureDate: req.DepartureDate,
		ReturnDate:    req.ReturnDate,
		Passengers:    req.Passengers,
		CabinClass:    req.CabinClass,
	})
	if err != nil {
		helper.JSON(w, ctx, nil, errs.NewWithMessage(http.StatusInternalServerError, "SearchFlightUseCase failed"))
	}
	helper.JSON(w, ctx, resp, nil)
}
