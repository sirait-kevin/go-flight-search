package httphandlers

import (
	"encoding/json"
	"go-flight-search/pkg/errs"
	"go-flight-search/pkg/helper"
	"net/http"
)

func (*SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req FlightSearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.JSON(w, ctx, nil, errs.NewWithMessage(http.StatusBadRequest, "Invalid request payload"))
		return
	}

}
