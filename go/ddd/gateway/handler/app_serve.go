package handler

import (
	"encoding/json"
	"net/http"

	"github.com/muhfaris/adsrobot/internal/logger"
	"github.com/muhfaris/adsrobot/internal/response"
)

type appHandler struct {
	*App
	HandlerFunc func(http.ResponseWriter, *http.Request) response.ResponseRequest
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	data := ah.HandlerFunc(w, r)
	if data.Error != nil {
		response := data.ParseError()
		w.WriteHeader(response.Status)
		serveResponse(w, r, response)
		return
	}

	response := data.Result.(response.DataRequestResponse)
	w.WriteHeader(response.Status)
	serveResponse(w, r, response)
}

func serveResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		logger.NewLogRequest(r).Error(err)
		return
	}
}
