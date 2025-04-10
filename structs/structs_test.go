package structs

import (
	"github.com/Bbanks14/ticketing-system/log/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogLevel_String(t *testing.T) {
	testlog.BeginTest()
	defer testlog.EndTest()

	t.Run("infoString", func(t *testing.T) {
		assert.Equal(t, "[INFO]", LevelInfo.String())
	})

	t.Run("warningString", func(t *testing.T) {
		assert.Equal(t, "[WARNING]", LevelWarning.String())
	})

	t.Run("errorString", func(t *testing.T) {
		assert.Equal(t, "[ERROR]", LevelError.String())
	})

	t.Run("fatalErrorString", func(t *testing.T) {
		assert.Equal(t, "[FATAL ERROR]", LevelFatal.String())
	})

	t.Run("testDebugString", func(t *testing.T) {
		assert.Equal(t, "[TEST DEBUG]", LevelTestDebug.String())
	})

	t.Run("undefinedString", func(t *testing.T) {
		assert.Equal(t, "[INFO]", LogLevel(7).String())
	})
}

func TestAsLogLevel_String(t *testing.T) {
	t.Run("infoString", func(t *testing.T) {
		assert.Equal(t, LevelInfo, AsLogLevel("info"))
	})

	t.Run("warningString", func(t *testing.T) {
		assert.Equal(t, LevelWarning, AAsLogLevel("warning"))
	})

	t.Run("fatalString", func(t *testing.T) {
		assert.Equal(t, LevelLevelFatal, AAsLogLevel("fatal"))
	})

	t.Run("undefinedString", func(t *testing.T) {
		assert.Equal(t, LogLevel(-1), AsLogLevel("undefined"))
	})
}

func TestCommand_String(t *testing.T) {
	testlog.BeginTest()
	defer testlog.EndTest()

	t.Run("fetchNumber", func(t *testing.T) {
		assert.Equal(t, "0", CommandFetch.String())
	})

	t.Run("submitNumber", func(t *testing.T) {
		assert.Equal(t, "1", CommandSubmit.String())
	})

	t.Run("exitNumber", func(t *testing.T) {
		assert.Equal(t, "2", CommandExit.String())
	})
}
