package server

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/meesooqa/go-web-example/server/mocks"
)

func TestHandle_NoHandlers_ReturnsServeMux(t *testing.T) {
	cfg := &mocks.ConfigMock{
		HostFunc:              func() string { return "localhost" },
		PortFunc:              func() int { return 8080 },
		ReadHeaderTimeoutFunc: func() time.Duration { return time.Second },
		WriteTimeoutFunc:      func() time.Duration { return time.Second },
		IdleTimeoutFunc:       func() time.Duration { return time.Second },
	}

	srv := New(cfg, nil, nil)
	handler := srv.handle()

	_, ok := handler.(*http.ServeMux)
	assert.True(t, ok)
}

func TestHandle_WithHandler_CallsHandle(t *testing.T) {
	handlerMock := &mocks.HandlerMock{
		HandleFunc: func(mux *http.ServeMux) {},
	}

	cfg := &mocks.ConfigMock{
		HostFunc:              func() string { return "localhost" },
		PortFunc:              func() int { return 8080 },
		ReadHeaderTimeoutFunc: func() time.Duration { return time.Second },
		WriteTimeoutFunc:      func() time.Duration { return time.Second },
		IdleTimeoutFunc:       func() time.Duration { return time.Second },
	}

	srv := New(cfg, []Handler{handlerMock}, nil)
	srv.handle()

	assert.Len(t, handlerMock.HandleCalls(), 1)
	assert.NotNil(t, handlerMock.HandleCalls()[0].Mux)
}

func TestHandle_WithMiddleware_CallsMiddleware(t *testing.T) {
	middlewareMock := &mocks.MiddlewareMock{
		HandleFunc: func(next http.Handler) http.Handler {
			return next
		},
	}

	cfg := &mocks.ConfigMock{
		HostFunc:              func() string { return "localhost" },
		PortFunc:              func() int { return 8080 },
		ReadHeaderTimeoutFunc: func() time.Duration { return time.Second },
		WriteTimeoutFunc:      func() time.Duration { return time.Second },
		IdleTimeoutFunc:       func() time.Duration { return time.Second },
	}

	srv := New(cfg, nil, []Middleware{middlewareMock})
	srv.handle()

	calls := middlewareMock.HandleCalls()
	assert.Len(t, calls, 1)

	next, ok := calls[0].Next.(*http.ServeMux)
	assert.True(t, ok)
	assert.NotNil(t, next)
}

func TestHandle_WithMultipleMiddlewares_CallsInOrder(t *testing.T) {
	mw1 := &mocks.MiddlewareMock{
		HandleFunc: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		},
	}

	mw2 := &mocks.MiddlewareMock{
		HandleFunc: func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		},
	}

	cfg := &mocks.ConfigMock{
		HostFunc:              func() string { return "localhost" },
		PortFunc:              func() int { return 8080 },
		ReadHeaderTimeoutFunc: func() time.Duration { return time.Second },
		WriteTimeoutFunc:      func() time.Duration { return time.Second },
		IdleTimeoutFunc:       func() time.Duration { return time.Second },
	}

	srv := New(cfg, nil, []Middleware{mw1, mw2})
	srv.handle()

	assert.Len(t, mw1.HandleCalls(), 1)
	assert.Len(t, mw2.HandleCalls(), 1)
}

func TestRun_CallsConfigMethods(t *testing.T) {
	cfg := &mocks.ConfigMock{
		HostFunc:              func() string { return "localhost" },
		PortFunc:              func() int { return 0 },
		ReadHeaderTimeoutFunc: func() time.Duration { return time.Second },
		WriteTimeoutFunc:      func() time.Duration { return time.Second },
		IdleTimeoutFunc:       func() time.Duration { return time.Second },
	}

	srv := New(cfg, nil, nil)

	errCh := make(chan error)
	go func() {
		errCh <- srv.Run()
	}()

	time.Sleep(100 * time.Millisecond)

	assert.Len(t, cfg.HostCalls(), 1)
	assert.Len(t, cfg.PortCalls(), 1)
	assert.Len(t, cfg.ReadHeaderTimeoutCalls(), 1)
	assert.Len(t, cfg.WriteTimeoutCalls(), 1)
	assert.Len(t, cfg.IdleTimeoutCalls(), 1)
}

func TestRun_ReturnsErrorOnInvalidHost(t *testing.T) {
	cfg := &mocks.ConfigMock{
		HostFunc:              func() string { return "invalid_host" },
		PortFunc:              func() int { return 8080 },
		ReadHeaderTimeoutFunc: func() time.Duration { return time.Second },
		WriteTimeoutFunc:      func() time.Duration { return time.Second },
		IdleTimeoutFunc:       func() time.Duration { return time.Second },
	}

	srv := New(cfg, nil, nil)

	err := srv.Run()
	assert.Error(t, err)
}
