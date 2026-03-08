package health

import (
	"net/http"
	"fmt"
)

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ignore requests which arent GET
		if r.Method != http.MethodGet {
			http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		}

		fmt.Fprintln(w, "fully operational API")
	}
}