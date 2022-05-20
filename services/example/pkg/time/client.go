package time

import "time"

type Client interface {
	// GetCurrent returns the current time
	GetCurrent() (time.Time, error)
}
