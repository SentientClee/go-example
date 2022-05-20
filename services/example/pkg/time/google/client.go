package google

import (
	"errors"
	"net/http"
	"net/url"
	stdtime "time"

	"github.com/timstudd/go-example/services/example/pkg/time"
)

type Config struct {
	HTTPEndpoint string
}

func New(conf *Config) (time.Client, error) {
	if conf.HTTPEndpoint == "" {
		return nil, errors.New("empty HTTP endpoint")
	}
	if _, err := url.Parse(conf.HTTPEndpoint); err != nil {
		return nil, errors.New("invalid HTTP endpoint")
	}

	return &client{
		httpEndpoint: conf.HTTPEndpoint,
	}, nil
}

type client struct {
	httpEndpoint string
}

func (c *client) GetCurrent() (stdtime.Time, error) {
	res, err := http.Head(c.httpEndpoint)
	if err != nil {
		return stdtime.Time{}, err
	}

	return http.ParseTime(res.Header.Get("Date"))
}
