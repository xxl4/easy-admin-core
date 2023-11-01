package server

import "time"

type Option func(*options)

type options struct {
	gracefulShutdownTimeout time.Duration
}

func setDefaultOptions() options {
	return options{
		gracefulShutdownTimeout: 5 * time.Second,
	}
}
