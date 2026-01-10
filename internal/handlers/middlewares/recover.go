package middlewares

import (
	"go-flight-search/pkg/errs"
	"go-flight-search/pkg/helper"
	"log"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			ctx := r.Context()
			if err := recover(); err != nil {
				log.Println("Recovered from panic:", err)
				helper.JSON(w, ctx, nil, errs.NewWithMessage(http.StatusInternalServerError, "Internal Server Error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
