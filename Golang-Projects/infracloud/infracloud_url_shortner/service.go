package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type InputUrl struct {
	url string `json:"url"`
}

type memService struct {
	baseUrl string
	mtx     sync.RWMutex
	urlMap  map[string]int
}

type Service interface {
	ShortenUrl(ctx context.Context, url string) error
}

func (s *memService) ShortenUrl(ctx context.Context, url string) (error, string) {
	s.mtx.RLock()
	s.baseUrl = "https://cflair.com/"
	defer s.mtx.RUnlock()
	if _, ok := s.urlMap[url]; ok {
		shortenUrl := GetShortenUrl(url, s.urlMap)
		return nil, s.baseUrl + string(shortenUrl) // POST = create, don't overwrite
	}
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Int()
	s.urlMap[url] = randInt
	// return the URL https://cflair.com/randomHash
	return nil, s.baseUrl + string(randInt)
}

func GetShortenUrl(url string, urlMap map[string]int) int {
	shortUrl, ok := urlMap[url]
	if !ok {
		fmt.Println("Not Possible!")
	}
	return shortUrl

}
