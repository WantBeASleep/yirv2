// TODO: сделать через zerolog
package log

import (
	"context"
	"io"
	"log/slog"
	"os"

	"yirv2/pkg/ctxlib"
)

type config struct {
	dest io.Writer
}

type LogOption func(*config)

func WithFileOutput(path string) LogOption {
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0o666)
	return LogOption(func(c *config) {
		c.dest = f
	})
}

func defaultConfig() config {
	return config{dest: os.Stdout}
}

type handler struct {
	slog.Handler
}

func (h handler) Handle(ctx context.Context, r slog.Record) error {
	attr := ctxlib.PublicGetAll(ctx)

	contextAttrs := make([]any, 0, len(attr))
	for k, v := range attr {
		contextAttrs = append(contextAttrs, slog.Any(k, v))
	}

	r.AddAttrs(slog.Group("context", contextAttrs...))

	return h.Handler.Handle(ctx, r)
}

func InitLogger(opts ...LogOption) {
	cfg := defaultConfig()
	for _, o := range opts {
		o(&cfg)
	}

	jsonHandler := slog.NewJSONHandler(cfg.dest, nil)

	logger := slog.New(handler{Handler: jsonHandler})
	slog.SetDefault(logger)
}
