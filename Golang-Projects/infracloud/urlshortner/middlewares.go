package urlshortner

// import (
// 	"context"
// 	"time"

// 	"github.com/go-kit/log"
// )

// // Middleware describes a service (as opposed to endpoint) middleware.
// type Middleware func(ServiceURL) ServiceURL

// func LoggingMiddleware(logger log.Logger) Middleware {
// 	return func(next ServiceURL) ServiceURL {
// 		return &loggingMiddleware{
// 			next:   next,
// 			logger: logger,
// 		}
// 	}
// }

// type loggingMiddleware struct {
// 	next   ServiceURL
// 	logger log.Logger
// }

// func (mw loggingMiddleware) ShortenUrl(ctx context.Context, p UrlModel) (url string, err error) {
// 	defer func(begin time.Time) {
// 		mw.logger.Log("method", "ShortenUrl", "id", p.URL, "took", time.Since(begin), "err", err)
// 	}(time.Now())
// 	return mw.next.ShortenUrl(ctx, p)
// }
