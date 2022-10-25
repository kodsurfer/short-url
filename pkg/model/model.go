package model

import "time"

type Request struct {
	URL     string        `json:"url"`
	Short   string        `json:"short"`
	Expires time.Duration `json:"expires"`
}

type Response struct {
	URL            string        `json:"url"`
	Short          string        `json:"short"`
	Expires        time.Duration `json:"expires"`
	XRateLimit     int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}
