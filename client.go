package httpclient

import (
	"bitbucket.org/snapmartinc/httpclient/middleware"
	"bitbucket.org/snapmartinc/logger"
	"net/http"
	"time"
)

func NewDefaultHttpClient(defaultEntry logger.Entry, timeout time.Duration) *http.Client {
	c := &http.Client{
		Timeout: timeout,
	}

	c.Transport = middleware.WithMiddleware(
		c.Transport,
		middleware.NewResponseLogger(defaultEntry),
		middleware.NewRequestLogger(defaultEntry),
	)

	return c
}
