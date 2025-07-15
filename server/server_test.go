package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meesooqa/go-web-example/server/mocks"
)

func TestHandle_NoMiddleware(t *testing.T) {
	cfg := &mocks.ConfigMock{}
	h1 := &mocks.HandlerMock{}
	h2 := &mocks.HandlerMock{}

	// Setup handlers to register endpoints
	h1.HandleFunc = func(mux *http.ServeMux) {
		mux.HandleFunc("/h1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("h1"))
		})
	}
	h2.HandleFunc = func(mux *http.ServeMux) {
		mux.HandleFunc("/h2", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("h2"))
		})
	}

	srv := New(cfg, []Handler{h1, h2}, nil)
	h := srv.handle()

	// Verify handlers attached
	req1 := httptest.NewRequest(http.MethodGet, "/h1", nil)
	res1 := httptest.NewRecorder()
	h.ServeHTTP(res1, req1)
	assert.Equal(t, http.StatusOK, res1.Code)
	assert.Equal(t, "h1", res1.Body.String())

	req2 := httptest.NewRequest(http.MethodGet, "/h2", nil)
	res2 := httptest.NewRecorder()
	h.ServeHTTP(res2, req2)
	assert.Equal(t, http.StatusOK, res2.Code)
	assert.Equal(t, "h2", res2.Body.String())

	// Ensure Handle was called on mocks
	assert.Len(t, h1.HandleCalls(), 1)
	assert.Len(t, h2.HandleCalls(), 1)
}

func TestHandle_WithMiddleware(t *testing.T) {
	cfg := &mocks.ConfigMock{}
	h := &mocks.HandlerMock{}
	m1 := &mocks.MiddlewareMock{}
	m2 := &mocks.MiddlewareMock{}

	// Handler registers a simple endpoint
	h.HandleFunc = func(mux *http.ServeMux) {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
	}

	// Middleware that adds header A then calls next
	m1.HandleFunc = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-A", "1")
			next.ServeHTTP(w, r)
		})
	}
	// Middleware that adds header B then calls next
	m2.HandleFunc = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-B", "2")
			next.ServeHTTP(w, r)
		})
	}

	srv := New(cfg, []Handler{h}, []Middleware{m1, m2})
	hand := srv.handle()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	hand.ServeHTTP(res, req)

	assert.Equal(t, "1", res.Header().Get("X-A"))
	assert.Equal(t, "2", res.Header().Get("X-B"))
	assert.Equal(t, "ok", res.Body.String())

	// Verify mock calls
	assert.Len(t, h.HandleCalls(), 1)
	assert.Len(t, m1.HandleCalls(), 1)
	assert.Len(t, m2.HandleCalls(), 1)
}

func TestRun_Error(t *testing.T) {
	cfg := &mocks.ConfigMock{}
	cfg.HostFunc = func() string { return "invalid_host" }
	cfg.PortFunc = func() int { return 12345 }
	cfg.ReadHeaderTimeoutFunc = func() time.Duration { return time.Second }
	cfg.WriteTimeoutFunc = func() time.Duration { return time.Second }
	cfg.IdleTimeoutFunc = func() time.Duration { return time.Second }

	srv := New(cfg, nil, nil)
	err := srv.Run()
	require.Error(t, err)
}
