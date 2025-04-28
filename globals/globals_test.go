package globals

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Bbanks14/log/testlog"
)

// TestGlobalStructures checks that all variables in the
// global package are correctly initialized
func TestGlobalStructures(t *testing.T) {
	testlog.BeginTest()
	defer testlog.EndTest()

	t.Run("ticketsNotNil", func(t *testing.T) {
		assert.NotNil(t, Tickets, "Tickets should not be nil")
	})

	t.Run("mailsNotNil", func(t *testing.T) {
		assert.NotNil(t, Mails, "Mails should not be nil")
	})

	t.Run("emptyServerConfig", func(t *testing.T) {
		assert.NotNil(t, ServerConfig, "Server Config should not be an empty struct")
	})

	t.Run("emptyLogConfig", func(t *testing.T) {
		assert.NotNil(t, LogConfig, "Log Config should not be an empty struct")
	})

	t.Run("SessionsNotNil", func(t *testing.T) {
		assert.NotNil(t, Sessions, "Sessions should not be nil")
	})
}
