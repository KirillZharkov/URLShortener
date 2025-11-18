package slogdiscard

import (
	"context"

	"golang.org/x/exp/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

// instead of honestly handling messages, he will ignore them
type DiscardHandler struct{}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// implementation of interface methods
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// Just ignore the log entry
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// Returns the same handler, since there are no attributes to save
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	// Returns the same handler, since there is no group to save
	return h
}

func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// Always returns false, as the log entry is ignored
	return false
}
