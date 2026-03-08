package handler

import (
	"fmt"
	"net/http"

	"github.com/72sevenzy2/golang-API/internal/service"
)

func GreetHandler(g service.Greeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if method is GET or not
		if r.Method != http.MethodGet {
			http.Error(w, "invalid method", http.StatusMethodNotAllowed)
			return
		}
		name := r.URL.Query().Get("name")
		resp, err := g.Greet(name)

		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, resp)
	}
}