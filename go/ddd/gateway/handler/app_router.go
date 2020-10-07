package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type facebookHandler struct {
	*App
	decoder *schema.Decoder

	// services
}

var (
	// GET http method
	GET = []string{http.MethodOptions, http.MethodGet}

	// POST http method
	POST = []string{http.MethodOptions, http.MethodPost}

	// PATCH http method
	PATCH = []string{http.MethodOptions, http.MethodPatch}

	// DELETE http method
	DELETE = []string{http.MethodOptions, http.MethodDelete}
)

func NewHandler(app *App) *facebookHandler {
	var decoder = schema.NewDecoder()
	return &facebookHandler{
		// extra
		decoder: decoder,
	}
}

// InternalHandler is all handler static data
func (h *facebookHandler) InternalHandler(app *App, r *mux.Router) {
	// r route to /rules
	r.Handle("/", appHandler{app, h.startCampaign}).Methods(POST...)
}
