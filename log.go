package eddnc

import (
	"golang.org/x/exp/slog"
)

var log = slog.Default()

func SetLog(logger *slog.Logger) { log = logger }
