package configuration

import "log/slog"

type LogConfiguration struct {
	LogLevel string
}

func ParseLogLevel(s string) (slog.Level, error) {
	var l slog.Level
	var err = l.UnmarshalText([]byte(s))
	return l, err
}
