package handler

import (
	"github.com/urlShortnerGo/pkg/dto"
	"github.com/urlShortnerGo/pkg/service"
	"net/http"
)

type UrlHandler struct {
	svc service.UrlShortnerService
}

func (h *UrlHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	urlReq := new(dto.GetUrl)
	if err := urlReq.ValidateRequest(); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	}

}
