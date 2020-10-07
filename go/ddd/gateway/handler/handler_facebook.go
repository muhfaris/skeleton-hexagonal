package handler

import (
	"net/http"

	"github.com/muhfaris/adsrobot/internal/response"
)

// facebookHandler request to facebook
func (h *facebookHandler) startCampaign(w http.ResponseWriter, r *http.Request) response.ResponseRequest {
	return response.ResponseRequest{
		Result: "Sussess",
	}
}
