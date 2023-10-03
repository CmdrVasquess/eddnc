package eddnc

import (
	"log/slog"
)

var log = slog.Default()

func SetLog(logger *slog.Logger) { log = logger }
