package local

import (
	stdtime "time"

	"github.com/timstudd/go-example/services/example/pkg/time"
)

type Config struct {
}

func New(conf *Config) (time.Client, error) {
	return &client{}, nil
}

type client struct {
}

func (c *client) GetCurrent() (stdtime.Time, error) {
	return stdtime.Now(), nil
}
