package urlshortner

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type ServiceURL interface {
	ShortenUrl(ctx context.Context, url UrlModel) (string, error)
	GetLongUrl(ctx context.Context, model UrlModel) (string, error)
}

type UrlModel struct {
	URL    string `json:"url"`
	GENUrl string
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type inmemService struct {
	mtx     sync.RWMutex
	baseUrl string `default:"https://cflair.com/"`
	urlMap  map[string]UrlModel
}

func NewInmemService() ServiceURL {
	return &inmemService{
		urlMap: map[string]UrlModel{},
	}
}

// Now need to more-correct this method, that's All for now :)
func (s *inmemService) ShortenUrl(ctx context.Context, model UrlModel) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s.baseUrl == "" {
		s.baseUrl = "https://cflair.com/"
	}
	// Case if GENUrl key exists.
	if _, ok := s.urlMap[model.URL]; ok {
		shortenUrl, err := GetShortenUrl(model.URL, s.urlMap)
		if err != nil {
			ctx.Err()
			//ctx.logger.Log(ErrAlreadyExists)
		}
		return shortenUrl.GENUrl, nil
	}
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Int()
	model.GENUrl = s.baseUrl + strconv.Itoa(randInt)
	s.urlMap[model.URL] = model
	return model.GENUrl, nil
}

func (s *inmemService) GetLongUrl(ctx context.Context, model UrlModel) (string, error) {

	// Get Key from Value of URL.
	longUrl := s.urlMap[model.URL]
	//http.Redirect(longUrl)
	return longUrl.URL, nil
}

func GetShortenUrl(url string, urlMap map[string]UrlModel) (UrlModel, error) {
	shortUrl, ok := urlMap[url]
	if !ok {
		return shortUrl, ErrNotFound
	}
	return shortUrl, nil
}
