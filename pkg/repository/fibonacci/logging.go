package fibonacci

import (
	"time"

	"github.com/go-kit/log"

	"github.com/Kuwerin/fibonacci/pkg/domain"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Repository
}

func LoggingMiddleware(logger log.Logger) func(Repository) Repository {
	return func(next Repository) Repository {
		return &loggingMiddleware{logger, next}
	}
}

func (l *loggingMiddleware) Save(fibonacciNum domain.Fibonacci) (err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"entity", "repository",
			"method", "save",
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.next.Save(fibonacciNum)
}

func (l *loggingMiddleware) Find(key uint64) (fibonacciNum domain.Fibonacci, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"entity", "repository",
			"method", "find",
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return l.next.Find(key)
}
