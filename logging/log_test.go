package logging

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConsoleLogger_Text(t *testing.T) {
	buf := &bytes.Buffer{}
	cfg := &ConfigMock{
		IsWriteToFileFunc: func() bool { return false },
		LevelFunc:         func() slog.Level { return slog.LevelInfo },
		OutputFormatFunc:  func() string { return "text" },
	}
	logger := slog.New(handler(cfg, buf))

	logger.Info("hello", "key", "value")

	out := buf.String()
	require.Contains(t, out, "INFO")
	require.Contains(t, out, "hello")
	require.Contains(t, out, "key=value")
}

func TestConsoleLogger_JSON(t *testing.T) {
	buf := &bytes.Buffer{}
	cfg := &ConfigMock{
		IsWriteToFileFunc: func() bool { return false },
		LevelFunc:         func() slog.Level { return slog.LevelWarn },
		OutputFormatFunc:  func() string { return "json" },
	}
	logger := slog.New(handler(cfg, buf))

	logger.Warn("warning", "foo", 123)

	line := buf.String()
	require.NotEmpty(t, line)
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(line), &obj)
	require.NoError(t, err)
	require.Equal(t, "warning", obj["msg"])
	require.Equal(t, float64(123), obj["foo"])
}

func TestFileLogger_WritesToFile(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testlog.*.log")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	cfg := &ConfigMock{
		IsWriteToFileFunc: func() bool { return true },
		PathFunc:          func() string { return tmpfile.Name() },
		LevelFunc:         func() slog.Level { return slog.LevelError },
		OutputFormatFunc:  func() string { return "text" },
	}

	logger, closer := New(cfg)
	defer closer.Close()

	logger.Error("fail", "code", 500)

	data, err := os.ReadFile(tmpfile.Name())
	require.NoError(t, err)
	s := string(data)
	require.Contains(t, s, "ERROR")
	require.Contains(t, s, "fail")
	require.Contains(t, s, "code=500")
}

func TestFileLogger_OpenError(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		cfg := &ConfigMock{
			IsWriteToFileFunc: func() bool { return true },
			PathFunc:          func() string { return "/invalid/path/not/writable.log" },
			LevelFunc:         func() slog.Level { return slog.LevelInfo },
			OutputFormatFunc:  func() string { return "text" },
		}
		New(cfg)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFileLogger_OpenError")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	require.Error(t, err)
	exitErr, ok := err.(*exec.ExitError)
	require.True(t, ok)
	require.Equal(t, 1, exitErr.ExitCode())
}
