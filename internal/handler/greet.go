package handler

import (
	"net/http"
	"encoding/json"
	"github.com/72sevenzy2/golang-API/internal/service"
	"github.com/72sevenzy2/golang-API/internal/response"
	"github.com/72sevenzy2/golang-API/internal/configs"
)

func GreetHandler(g service.Greeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req configs.GreetRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid json")
			return
		}

		msg, count, err := g.Greet(req.Name)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		resp := configs.GreetResponse{
			Message: msg,
			Count: count,
		}

		response.Json(w, http.StatusOK, resp)
	}
}