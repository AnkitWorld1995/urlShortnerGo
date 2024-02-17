package service

import (
	"github.com/urlShortnerGo/pkg/domain"
	"github.com/urlShortnerGo/pkg/dto"
)

type urlService struct {
	repo domain.UrlRepo
}

func NewUrlService(r domain.UrlRepo) *urlService {
	return &urlService{
		repo: r,
	}
}

type UrlShortnerService interface {
}

func (svc *urlService) Get(req *dto.GetUrl) {
	// req.LongUrl
}

func (svc *urlService) Post() {}
