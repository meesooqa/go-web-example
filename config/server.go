package config

import "time"

func (cfg *Server) Host() string {
	return cfg.RawHost
}

func (cfg *Server) Port() int {
	return cfg.RawPort
}

func (cfg *Server) ReadHeaderTimeout() time.Duration {
	return cfg.RawReadHeaderTimeout
}

func (cfg *Server) WriteTimeout() time.Duration {
	return cfg.RawWriteTimeout
}

func (cfg *Server) IdleTimeout() time.Duration {
	return cfg.RawIdleTimeout
}
