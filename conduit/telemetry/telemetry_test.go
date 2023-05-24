package telemetry

import (
	"testing"

	"github.com/algorand/indexer/version"
	"github.com/stretchr/testify/require"
)

func TestTelemetryConfig(t *testing.T) {

	config := MakeTelemetryConfig("test-uri", "test-index", "test-user", "test-password")
	require.Equal(t, true, config.Enable)
	require.Equal(t, "test-uri", config.URI)
	require.Equal(t, "test-index", config.Index)
	require.Equal(t, "test-user", config.UserName)
	require.Equal(t, "test-password", config.Password)
	require.NotEqual(t, "", config.GUID)
}

func TestMakeTelemetryStartupEvent(t *testing.T) {
	config := Config{
		GUID: "test-guid",
	}
	state, err := MakeOpenSearchClient(config)
	require.NoError(t, err)
	event := state.MakeTelemetryStartupEvent()
	require.Equal(t, "starting conduit", event.Message)
	require.Equal(t, "test-guid", event.GUID)
	require.NotEqual(t, version.LongVersion(), event.Version)
}
