package app

import "net/http"

func Run() *http.Server {
	// TODO:
	// - load config
	// - init redis
	// - init providers
	// - init usecase
	// - init router

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
