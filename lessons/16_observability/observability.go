package observability

import (
	"context"
	"log/slog"
	"time"
)

type RequestStats struct {
	Route    string
	Status   int
	Duration time.Duration
}

func LogRequest(ctx context.Context, logger *slog.Logger, stats RequestStats) {
	logger.InfoContext(ctx, "request completed",
		"route", stats.Route,
		"status", stats.Status,
		"duration_ms", stats.Duration.Milliseconds(),
	)
}

func IsSuccessful(status int) bool {
	return status >= 200 && status < 400
}
