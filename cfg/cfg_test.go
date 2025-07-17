package cfg

import (
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	c, err := Load("testdata/config.yml")
	require.NoError(t, err)

	t.Run("Project", func(t *testing.T) {
		t.Parallel()
		assert.IsType(t, &Project{}, c.Project)
		assert.Equal(t, "test", c.Project.Name)
		assert.Equal(t, "0.0.0", c.Project.Version)
		assert.Equal(t, true, c.Project.Debug)
		assert.Equal(t, "development", c.Project.Environment)
	})

	t.Run("Log", func(t *testing.T) {
		t.Parallel()
		assert.IsType(t, &Log{}, c.Log)
		assert.Equal(t, "debug", c.Log.RawLevelCode)
		assert.Equal(t, slog.LevelDebug, c.Log.Level())
		assert.Equal(t, slog.LevelDebug, c.Log.RawLevel)
		assert.Equal(t, "text", c.Log.OutputFormat())
		assert.Equal(t, "text", c.Log.RawOutputFormat)
		assert.True(t, c.Log.IsWriteToFile())
		assert.True(t, c.Log.RawWriteToFile)
		assert.Equal(t, "var/log/app.log", c.Log.Path())
		assert.Equal(t, "var/log/app.log", c.Log.RawPath)
	})

	t.Run("Server", func(t *testing.T) {
		t.Parallel()
		cfg := c.Server
		assert.IsType(t, &Server{}, cfg)
		assert.Equal(t, "1.22.333.004", cfg.Host())
		assert.Equal(t, "1.22.333.004", cfg.RawHost)
		assert.Equal(t, 1111, cfg.Port())
		assert.Equal(t, 1111, cfg.RawPort)
		assert.Equal(t, 100*time.Second, cfg.ReadHeaderTimeout())
		assert.Equal(t, 100*time.Second, cfg.RawReadHeaderTimeout)
		assert.Equal(t, 202*time.Second, cfg.WriteTimeout())
		assert.Equal(t, 202*time.Second, cfg.RawWriteTimeout)
		assert.Equal(t, 333*time.Second, cfg.IdleTimeout())
		assert.Equal(t, 333*time.Second, cfg.RawIdleTimeout)
	})

	t.Run("Handler", func(t *testing.T) {
		t.Parallel()
		cfg := c.Handler
		assert.IsType(t, &Handler{}, cfg)
		assert.Equal(t, "test", cfg.ThemesDir())
		assert.Equal(t, "test", cfg.RawThemesDir)
		assert.Equal(t, "theme_test", cfg.Theme())
		assert.Equal(t, "theme_test", cfg.RawTheme)
	})
}

func TestLoadConfigNotFoundFile(t *testing.T) {
	r, err := Load("/tmp/ea63af9a-f3c2-4f43-b254-61808754a169.txt")
	assert.Nil(t, r)
	assert.EqualError(t, err, "open /tmp/ea63af9a-f3c2-4f43-b254-61808754a169.txt: no such file or directory")
}

func TestLoadConfigInvalidYaml(t *testing.T) {
	r, err := Load("testdata/file.txt")

	assert.Nil(t, r)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `Not Yaml` into config.AppConfig")
}
