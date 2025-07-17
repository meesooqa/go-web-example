package lgr

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
	"testing/slogtest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONHandler(t *testing.T) {
	var buf bytes.Buffer
	cfg := &ConfigMock{
		LevelFunc:        func() slog.Level { return slog.LevelInfo },
		OutputFormatFunc: func() string { return "json" },
	}
	newHandler := func(t *testing.T) slog.Handler {
		buf.Reset()
		return handler(cfg, &buf)
	}

	result := func(t *testing.T) map[string]any {
		var m map[string]any
		err := json.Unmarshal(buf.Bytes(), &m)

		require.NoError(t, err)
		assert.Equal(t, "INFO", m[slog.LevelKey])

		return m
	}

	slogtest.Run(t, newHandler, result)
}

func TestTextHandler_WithSlogtestRun(t *testing.T) {
	var buf bytes.Buffer
	cfg := &ConfigMock{
		LevelFunc:        func() slog.Level { return slog.LevelInfo },
		OutputFormatFunc: func() string { return "text" },
	}
	newHandler := func(t *testing.T) slog.Handler {
		buf.Reset()
		return handler(cfg, &buf)
	}

	result := func(t *testing.T) map[string]any {
		// "time=... level=INFO msg=msg a=b G.c=d G.H.e=f"
		parts := strings.Fields(strings.TrimSpace(buf.String()))

		root := make(map[string]any)
		for _, kv := range parts {
			pair := strings.SplitN(kv, "=", 2)
			if len(pair) != 2 {
				continue
			}
			key, val := pair[0], pair[1]
			segs := strings.Split(key, ".")
			cur := root
			for i := 0; i < len(segs)-1; i++ {
				seg := segs[i]
				next, ok := cur[seg].(map[string]any)
				if !ok {
					next = make(map[string]any)
					cur[seg] = next
				}
				cur = next
			}
			cur[segs[len(segs)-1]] = val
		}
		return root
	}

	slogtest.Run(t, newHandler, result)
}
