package service

import (
	"github.com/urlShortnerGo/pkg/domain"
	"github.com/urlShortnerGo/pkg/dto"
)

type UrlService struct {
	repo domain.UrlRepo
}

func NewUrlService(r domain.UrlRepo) *UrlService {
	return &UrlService{
		repo: r,
	}
}

type UrlShortnerService interface {
}

func (svc *UrlService) Get(req *dto.GetUrl) {
	// req.LongUrl
}

func (svc *UrlService) Post() {}
