package logger

type Option func(*options)

type options struct {
	driver string
	path   string
	level  string
	stdout string
	cap    uint
}

func setDefault() options {
	return options{
		driver: "default",
		path:   "temp/logs",
		level:  "warn",
		stdout: "default",
	}
}

func WithType(s string) Option {
	return func(o *options) {
		o.driver = s
	}
}

func WithPath(s string) Option {
	return func(o *options) {
		o.path = s
	}
}

func WithLevel(s string) Option {
	return func(o *options) {
		o.level = s
	}
}

func WithStdout(s string) Option {
	return func(o *options) {
		o.stdout = s
	}
}

func WithCap(n uint) Option {
	return func(o *options) {
		o.cap = n
	}
}
