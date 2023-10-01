package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// type Middleware func(Service) Service

// func LoggingMiddleware(logger log.Logger) Middleware {
// 	return func(next Service) Service {
// 		return &loggingMiddleware{
// 			next:   next,
// 			logger: logger,
// 		}
// 	}
// }

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) PostProfile(ctx context.Context, url string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PostProfile", "URL", url, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.ShortenUrl(ctx, url)
}
