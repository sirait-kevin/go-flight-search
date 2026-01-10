package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"go-flight-search/internal/domain"
	"go-flight-search/pkg/errs"
	"go-flight-search/pkg/helper"
	"net/http"
	"sort"
	"time"
)

func (u *SearchFlightsUsecase) Execute(ctx context.Context, q domain.SearchQuery) (*[]domain.Flight, bool, error) {
	start := time.Now()

	cacheKey := u.makeCacheKey(q)

	if u.Cache != nil {
		if data, err := u.Cache.Get(cacheKey); err == nil {
			var cached []domain.Flight
			if err := json.Unmarshal(data, &cached); err == nil {
				return &cached, true, nil
			}
		}
	}

	type result struct {
		flights *[]domain.Flight
		err     error
	}

	ch := make(chan result, len(u.Providers))

	for _, p := range u.Providers {
		go func(p FlightProvider) {
			flights, err := p.Search(q)
			ch <- result{flights: flights, err: err}
		}(p)
	}

	var all []domain.Flight
	var err error

	timeout := time.After(5 * time.Second)

	for i := 0; i < len(u.Providers); i++ {
		select {
		case r := <-ch:
			if r.err == nil {
				all = append(all, *r.flights...)
				err = nil
			} else {
				err = errs.NewWithMessage(http.StatusInternalServerError, r.err.Error())
			}
		case <-timeout:
			break
		}
	}

	if err != nil {
		return nil, false, err
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].PriceAmount < all[j].PriceAmount
	})

	// 5️⃣ Store to cache
	if u.Cache != nil {
		if raw, err := json.Marshal(all); err == nil {
			_ = u.Cache.Set(cacheKey, raw, 300) // 5 min TTL
		}
	}

	fmt.Println("Search time:", time.Since(start))

	return &all, false, nil
}

// makeCacheKey generates deterministic cache key
func (u *SearchFlightsUsecase) makeCacheKey(q domain.SearchQuery) string {
	raw := fmt.Sprintf(
		"%s|%s|%s|%v|%d|%s",
		q.Origin,
		q.Destination,
		q.DepartureDate,
		q.ReturnDate,
		q.Passengers,
		q.CabinClass,
	)

	return helper.Hash(raw)
}
